package josh

import "strconv"

// Type var constraint for all possible integer types.
type integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

// Get and parse an integer value for the named path wildcard in the URL.
//
// Do not use it with [Must]! Wrapping a nil *Error into [error]
// will make it non-nil.
func GetID[T integer](r Req, name string) (T, *Error) {
	raw := r.PathValue(name)
	var def T
	if raw == "" {
		return def, &Error{Detail: "path parameter " + name + " is required"}
	}
	bitSize := getBitSize[T]()

	if isUnsigned[T]() {
		parsed, err := strconv.ParseUint(raw, 10, bitSize)
		if err != nil {
			return def, &Error{Detail: "invalid path parameter " + name}
		}
		return T(parsed), nil
	}

	parsed, err := strconv.ParseInt(raw, 10, bitSize)
	if err != nil {
		return def, &Error{Detail: "invalid path parameter " + name}
	}
	return T(parsed), nil
}

// Get the size in bits of the value of the given type.
func getBitSize[T integer]() int {
	var def T
	switch any(def).(type) {
	case int, uint, uintptr, int64, uint64:
		return 64
	case int8, uint8:
		return 8
	case int16, uint16:
		return 16
	default:
		return 32
	}
}

// Check if the given type is an unsigned integer.
func isUnsigned[T integer]() bool {
	var def T
	switch any(def).(type) {
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return true
	}
	return false
}
