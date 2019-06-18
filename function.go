package lesphina

import (
	"regexp"
)

var rJSONFieldName = regexp.MustCompile(`json:"(.+?)"`)

// Function ...
type Function struct {
	Name    string     `json:"name"`
	Recv    []*Element `json:"recv,omitempty"`
	In      []*Element `json:"in,omitempty"`
	Out     []*Element `json:"out,omitempty"`
	RawBody string     `json:"raw_body"`
}

// InParams ...
func (i *Function) InParams() []*Element { return i.In }

// OutParams ...
func (i *Function) OutParams() []*Element { return i.Out }

// FirstInParam ...
func (i *Function) FirstInParam(pattern string) *Element { return firstInParam(i, pattern) }

// FirstOutParam ...
func (i *Function) FirstOutParam(pattern string) *Element { return firstOutParam(i, pattern) }

// -- implement Entry interface --

// GetName ...
func (i *Function) GetName() string { return i.Name }

// GetKind ...
func (i *Function) GetKind() Kind { return KindFunction }

// JSON ...
func (i *Function) JSON() string { return marshal(i) }
