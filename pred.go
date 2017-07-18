package pred

import (
	"bytes"
	"io"
)

// New create new builder
func New(query string, ph string) *Builder {
	return &Builder{
		baseQuery:   query,
		placeHolder: ph,
	}
}

// OperatorType operator type
type OperatorType string

// operators
const (
	InOperatorType       = "in"
	NotInOperatorType    = "not in"
	EqualOperatorType    = "="
	NotEqualOperatorType = "!="
	IsOperatorType       = "is"
	IsNotOperatorType    = "is not"
	LikeOperatorType     = "like"
	NotLikeOperatorType  = "not like"
	AndOperatorType      = "and"
	OrOperatorType       = "or"
)

// Null SQL null type
const Null = "null"

// ValType value type
type ValType interface{}

// Writer defines the interface
type Writer interface {
	io.Writer
	Append(...interface{})
}

var _ Writer = NewWriter()

// BytesWriter implments Writer and save SQL in bytes.Buffer
type BytesWriter struct {
	writer *bytes.Buffer
	buffer []byte
	args   []interface{}
}

// NewWriter creates a new string writer
func NewWriter() *BytesWriter {
	w := &BytesWriter{}
	w.writer = bytes.NewBuffer(w.buffer)
	return w
}

// Write writes data to Writer
func (s *BytesWriter) Write(buf []byte) (int, error) {
	return s.writer.Write(buf)
}

// Append appends args to Writer
func (s *BytesWriter) Append(args ...interface{}) {
	s.args = append(s.args, args...)
}

// Cond defines an interface of predicate
type Cond interface {
	And(...Cond) Cond
	Or(...Cond) Cond
	WriteTo(Writer) error
}

// Pred expression
type Pred struct {
	Col string
	Ope OperatorType
	Val ValType
}

// And and
func (p *Pred) And(conds ...Cond) Cond {
	return And(p, And(conds...))
}

// Or or
func (p *Pred) Or(conds ...Cond) Cond {
	return Or(p, Or(conds...))
}

// WriteTo write sql
func (p *Pred) WriteTo(w Writer) error {
	return nil
}

// Builder predicate builder
type Builder struct {
	baseQuery   string
	placeHolder string
	limit       int
	offset      int
	orderBy     string
	havings     Cond
	wheres      Cond
}

// Where create where clause
func (b *Builder) Where(pred ...Cond) *Builder {
	return b
}

// Eq equal
func Eq(col string, val interface{}) *Pred {
	return &Pred{
		Col: col,
		Ope: EqualOperatorType,
		Val: val,
	}
}

// NotEq not equal
func NotEq(col string, val interface{}) *Pred {
	return &Pred{
		Col: col,
		Ope: NotEqualOperatorType,
		Val: val,
	}
}

// IsNull not equal
func IsNull(col string) *Pred {
	return &Pred{
		Col: col,
		Ope: IsOperatorType,
		Val: Null,
	}
}
