package protobuf

import (
	"io"

	"github.com/franky-li/openapi2proto/internal/option"
)

const (
	priorityEnum = iota
	priorityMessage
	priorityExtension
	priorityService
)

// Option is used to pass options to several methods
type Option = option.Option

// Builtin types
var (
	BoolType   = newBuiltin("bool")
	BytesType  = newBuiltin("bytes")
	DoubleType = newBuiltin("double")
	FloatType  = newBuiltin("float")
	Int32Type  = newBuiltin("int32")
	Int64Type  = newBuiltin("int64")
	StringType = newBuiltin("string")
)

// Boxed types
var (
	AnyType         = NewMessage("google.protobuf.Any")
	BoolValueType   = NewMessage("google.protobuf.BoolValue")
	BytesValueType  = NewMessage("google.protobuf.BytesValue")
	DoubleValueType = NewMessage("google.protobuf.DoubleValue")
	FloatValueType  = NewMessage("google.protobuf.FloatValue")
	Int32ValueType  = NewMessage("google.protobuf.Int32Value")
	Int64ValueType  = NewMessage("google.protobuf.Int64Value")
	NullValueType   = NewMessage("google.protobuf.NullValue")
	StringValueType = NewMessage("google.protobuf.StringValue")
)

// value type
var (
	StructType = NewMessage("google.protobuf.Struct")
)

// list type
var (
	ListValueType = NewMessage("google.protobuf.ListValue")
)

var (
	emptyMessage = NewMessage("google.protobuf.Empty")
)

// Encoder is responsible for taking a protobuf.Package object and
// encodes it into textual representation
type Encoder struct {
	dst    io.Writer
	indent string
}

// GlobalOption represents a Protocol Buffers global option
type GlobalOption struct {
	name  string
	value string
}

// Package represnets a Protocol Buffers Package.
type Package struct {
	name     string
	imports  []string
	children []Type
	options  []*GlobalOption
}

// Type is an interface to group different Protocol Buffer types
type Type interface {
	Name() string
	Priority() int
}

// Container is a special type that can have child types
type Container interface {
	Type
	AddType(Type)
	Children() []Type
}

// Enum represents a Protocol Buffers enum type
type Enum struct {
	comment  string
	elements []interface{}
	name     string
}

// Map represents a Protocol Buffers map type
type Map struct {
	key   Type
	value Type
}

// Builtin represents a Protocol Buffers builting type
type Builtin string

// Message is a composite type
type Message struct {
	children []Type
	comment  string
	fields   []*Field
	name     string
}

// Field is a field in a Message
type Field struct {
	comment  string
	index    int
	name     string
	repeated bool
	typ      Type
}

// ExtensionField is a field in an extended field
type ExtensionField struct {
	name   string
	typ    string
	number int
}

// Extension represents an extended message
type Extension struct {
	base   string
	fields []*ExtensionField
}

// RPC represents an RPC call associated with a Service
type RPC struct {
	comment   string
	name      string
	parameter Type
	response  Type

	options []interface{}
}

// Service defines a service with many RPC endpoints
type Service struct {
	name string
	rpcs []*RPC
}

// HTTPAnnotation represents a google.api.http option
type HTTPAnnotation struct {
	method string
	path   string
	body   string
}

// RPCOption represents simple rpc options
type RPCOption struct {
	name  string
	value interface{}
}

// Reference is a special type of Type that can pass the
// protobuf.Type system, but requires that  it be resolved
// at runtime to get the actual Type behind it. This is
// used to resolve circular dependencies that are found
// during compilation phase
type Reference struct {
	name string
}
