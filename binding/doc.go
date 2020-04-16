package binding

/*
Supported reflection types for built-in handlers.

Note that users can roll their own Handlers for app-specific use cases
if the type is not supported by the built-in provided handlers.

Where the type is marked as Supported, the binding package provides
a Handler with the name {{Type}}Handler.

Each of these provided handlers will accommodate inputs from any
of the supported types, and outputs to any of the supported types.

The code for each case of Handler is in the file named {{Type}}.go
in lowercase.

In most cases, type conversion between dissimilar types uses the
following method :

- If the types are the same, the value is passed through without modification.
- Conversion between numeric types is via type casting.
- Otherwise, render the value into a string using the %v printf directive,
  and then parsed from that string using standard Go tooling.
- It underlying types have a String() method, then this is applied during
  fmt.Sprintf("%v")

For exact details on how any type may convert, refer to the behaviour
of the %v directive in the fmt.Sprintf() command.

const (
	Invalid Kind = iota
	Bool        // Supported

	Int
	Int8
	Int16
	Int32
	Int64     // Supported
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64   // Supported

	// No support for complex types
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface

	Ptr
	UnsafePointer

	String      // Supported
	Struct      // Supported

	Slice       // Supported, but no conversions offered
	Map         // Supported, but no conversions offered
)
*/
