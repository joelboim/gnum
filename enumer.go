package gnum

// Enumer is an interface for using Enum instances with generics,
// e.g, `func foo[T Enumer[T]](enum T)` could do any Enum operations
// while preserving the original Enum type (T)
type Enumer[T ~int] interface {
	~int
	Config() *Config
	Enums() []T
	Name() string
	Names() []string
	Parse(name string) (T, error)
	String() string
	Strings() []string
	Type() string
}

// Enums is a static function to handel all enums that implements Enumer[T] interface.
// It returns a list of all Enum[T] declarations mapped to T.
func Enums[T Enumer[T]]() []T {
	return T.Enums(-1)
}

// Names is a static function to handel all enums that implements Enumer[T] interface.
// It returns a list of all Enum[T] names.
// (the programmatic string representation of the enum value).
func Names[T Enumer[T]]() []string {
	return T.Names(-1)
}

// Parse is a static function to handel all enums that implements Enumer[T] interface.
// It will try to parse the given name with the underline Enum.Parse implementation.
func Parse[T Enumer[T]](name string) (T, error) {
	enum, err := T.Parse(-1, name)
	if err != nil {
		return -1, err
	}

	return enum, nil
}

// Strings is a static function to handel all enums that implements Enumer[T] interface.
// It returns a list of all Enum[T] strings.
func Strings[T Enumer[T]]() []string {
	return T.Strings(-1)
}

// Type is a static function to handel all enums that implements Enumer[T] interface.
// It returns the underline type name.
func Type[T Enumer[T]]() string {
	return T.Type(-1)
}
