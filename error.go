package josh

// An error response. The structure follows the JSON:API spec.
//
// https://jsonapi.org/format/#error-objects
type Error struct {
	// ID is a unique identifier for this particular occurrence of a problem.
	ID string `json:"id,omitempty"`

	// Title is a short, human-readable summary of the problem.
	//
	// It SHOULD NOT change from occurrence to occurrence of the problem,
	// except for purposes of localization.
	Title string `json:"title,omitempty"`

	// Detail is a human-readable explanation specific to this occurrence of the problem.
	//
	// Like title, this field's value can be localized.
	Detail string `json:"detail,omitempty"`

	// Status is the HTTP status code applicable to this problem, expressed as a string value.
	Status string `json:"status,omitempty"`

	// Code is an application-specific error code, expressed as a string value.
	Code string `json:"code,omitempty"`

	// Source is a reference to the primary source of the error.
	//
	// Constructed with either [SourcePointer], [SourceParameter], or [SourceHeader].
	Source source `json:"source,omitempty"`

	// Meta is an object containing non-standard meta-information about the error.
	Meta *map[string]interface{} `json:"meta,omitempty"`
}

type source struct {
	// A JSON Pointer (RFC6901) to the value in the request document that caused the error.
	Pointer string `json:"pointer,omitempty"`

	// Which URI query parameter caused the error.
	Parameter string `json:"parameter,omitempty"`

	// The name of a single request header which caused the error.
	Header string `json:"header,omitempty"`
}

// A JSON Pointer (RFC6901) to the value in the request document that caused the error.
func SourcePointer(v string) source {
	return source{Pointer: v}
}

// Which URI query parameter caused the error.
func SourceParameter(v string) source {
	return source{Parameter: v}
}

// The name of a single request header which caused the error.
func SourceHeader(v string) source {
	return source{Header: v}
}
