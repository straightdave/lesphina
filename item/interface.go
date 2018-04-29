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

func (m *InterfaceMethod) FirstInParamLike(pattern string) *Element {
	if strings.HasPrefix(pattern, "~") {
		// end with something

		pattern = strings.TrimPrefix(pattern, "~")
		for _, p := range m.In {
			if strings.HasSuffix(p.BaseType, pattern) {
				return p
			}
		}
	}

	if strings.HasSuffix(pattern, "~") {
		// start with something

		pattern = strings.TrimSuffix(pattern, "~")
		for _, p := range m.In {
			if strings.HasPrefix(p.BaseType, pattern) {
				return p
			}
		}
	}

	// full word match
	for _, p := range m.In {
		if p.BaseType == pattern {
			return p
		}
	}

	// return non-nil empty value to surpress nil pointer panic
	return &Element{}
}

func (m *InterfaceMethod) FirstOutParamLike(pattern string) *Element {
	if strings.HasPrefix(pattern, "~") {
		// end with something

		pattern = strings.TrimPrefix(pattern, "~")
		for _, p := range m.Out {
			if strings.HasSuffix(p.BaseType, pattern) {
				return p
			}
		}
	}

	if strings.HasSuffix(pattern, "~") {
		// start with something

		pattern = strings.TrimSuffix(pattern, "~")
		for _, p := range m.Out {
			if strings.HasPrefix(p.BaseType, pattern) {
				return p
			}
		}
	}

	// full word match
	for _, p := range m.Out {
		if p.BaseType == pattern {
			return p
		}
	}

	// return non-nil empty value to surpress nil pointer panic
	return &Element{}
}
