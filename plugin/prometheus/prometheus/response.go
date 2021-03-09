package prometheus

import "encoding/json"

type StatusType string

const (
	StatusTypeSuccess StatusType = "success"
	StatusTypeError   StatusType = "error"
)

type ErrorType string

const (
	ErrorNone        ErrorType = ""
	ErrorTimeout     ErrorType = "timeout"
	ErrorCanceled    ErrorType = "canceled"
	ErrorExec        ErrorType = "execution"
	ErrorBadData     ErrorType = "bad_data"
	ErrorInternal    ErrorType = "internal"
	ErrorUnavailable ErrorType = "unavailable"
	ErrorNotFound    ErrorType = "not_found"
)

type Response struct {
	Status    StatusType `json:"status"`
	Data      QueryData  `json:"data,omitempty"`
	ErrorType ErrorType  `json:"errorType,omitempty"`
	Error     string     `json:"error,omitempty"`
	Warnings  []string   `json:"warnings,omitempty"`
}

// ValueType describes a type of a value.
type ValueType string

// The valid value types.
const (
	ValueTypeNone   ValueType = "none"
	ValueTypeVector ValueType = "vector"
	ValueTypeScalar ValueType = "scalar"
	ValueTypeMatrix ValueType = "matrix"
	ValueTypeString ValueType = "string"
)

// QueryData struct
type QueryData struct {
	ResultType ValueType       `json:"resultType"`
	Result     json.RawMessage `json:"result"`
}

type ResultVectorType []ResultVector

type ResultVector struct {
	Metric json.RawMessage `json:"metric"`
	Value  []interface{}   `json:"value"`
}
