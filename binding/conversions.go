package binding

// TODO - build up a couple of use cases to get the code right (string <--> float for example)
// then convert this to generated code

import (
	"fmt"
	"strconv"
)

type DataType int

const defaultFloatFormat = "%.2f"

type GetterSetter struct {
	Wrapped Binding
	Getter  func() interface{}
	Setter  func(interface{})
}

type StringWrapper struct {
	GetterSetter
	Fmt string
}

func (s *StringWrapper) Get() string {
	if _, ok := s.Wrapped.(IString); ok {
		return s.Getter().(string)
	}
	if _, ok := s.Wrapped.(IFloat64); ok {
		if v, ok := s.Getter().(float64); ok {
			return fmt.Sprintf(s.Fmt, v)
		}
	}
	return ""
}

func (s *StringWrapper) GetString() string { return s.Get() }

func (s *StringWrapper) Set(str string) {
	if _, ok := s.Wrapped.(IString); ok {
		s.Setter(str)
	}
	if _, ok := s.Wrapped.(IFloat64); ok {
		if v, err := strconv.ParseFloat(str, 64); err == nil {
			s.Setter(v)
		}
	}
}

func (s *StringWrapper) SetString(str string) { s.Set(str) }

func (s *StringWrapper) AddListener(n Notifiable) {
	s.Wrapped.AddListener(n)
}

func (s *StringWrapper) DeleteListener(n Notifiable) {
	s.Wrapped.DeleteListener(n)
}

func (s *StringWrapper) AddStringListener(listener func(string)) *NotifyFunction {
	if b, ok := s.Wrapped.(IString); ok {
		return b.AddStringListener(listener)
	}
	if b, ok := s.Wrapped.(IFloat64); ok {
		println("adding float listener and wrapping it as a string")
		return b.AddFloat64Listener(func(v float64) {
			listener(fmt.Sprintf(s.Fmt, v))
		})
	}
	fmt.Printf("ERROR: unknown wrapped binding type %T %+v\n", s.Wrapped, s.Wrapped)
	return nil
}

// AsString returns a new binding that wraps an existing binding as a stringBinding
func AsString(gs GetterSetter, floatFormat string) *StringWrapper {
	return &StringWrapper{gs, floatFormat}
}
