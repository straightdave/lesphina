package item

import (
	"strings"
)

type Interface struct {
	Name    string             `json:"name"`
	RawBody string             `json:"raw_body"`
	Methods []*InterfaceMethod `json:"methods"`
}

type InterfaceMethod struct {
	Name    string     `json:"name"`
	RawType string     `json:"raw_type"`
	In      []*Element `json:"in,omitempty"`
	Out     []*Element `json:"out,omitempty"`
}

type funcLike interface {
	InParams() []*Element
	OutParams() []*Element
}

func (i *InterfaceMethod) InParams() []*Element  { return i.In }
func (i *InterfaceMethod) OutParams() []*Element { return i.Out }

func FirstInParam(f funcLike, pattern string) *Element {
	if strings.HasPrefix(pattern, "~") {
		// end with something

		pattern = strings.TrimPrefix(pattern, "~")
		for _, p := range f.InParams() {
			if strings.HasSuffix(p.BaseType, pattern) {
				return p
			}
		}
	}

	if strings.HasSuffix(pattern, "~") {
		// start with something

		pattern = strings.TrimSuffix(pattern, "~")
		for _, p := range f.InParams() {
			if strings.HasPrefix(p.BaseType, pattern) {
				return p
			}
		}
	}

	// full word match
	for _, p := range f.InParams() {
		if p.BaseType == pattern {
			return p
		}
	}

	// return non-nil empty value to surpress nil pointer panic
	return &Element{}
}

func FirstOutParam(f funcLike, pattern string) *Element {
	if strings.HasPrefix(pattern, "~") {
		// end with something

		pattern = strings.TrimPrefix(pattern, "~")
		for _, p := range f.OutParams() {
			if strings.HasSuffix(p.BaseType, pattern) {
				return p
			}
		}
	}

	if strings.HasSuffix(pattern, "~") {
		// start with something

		pattern = strings.TrimSuffix(pattern, "~")
		for _, p := range f.OutParams() {
			if strings.HasPrefix(p.BaseType, pattern) {
				return p
			}
		}
	}

	// full word match
	for _, p := range f.OutParams() {
		if p.BaseType == pattern {
			return p
		}
	}

	// return non-nil empty value to surpress nil pointer panic
	return &Element{}
}
