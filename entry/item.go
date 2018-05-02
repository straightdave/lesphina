package entry

import (
	"encoding/json"
	"strings"
)

type Kind int

const (
	KindUnknown Kind = iota
	KindImport
	KindElement
	KindVar
	KindInterface
	KindStruct
	KindFunction
	KindInterfaceMethod
)

type Entry interface {
	GetName() string
	GetKind() Kind
}

func marshal(i Entry) string {
	j, _ := json.MarshalIndent(i, "", "    ")
	return string(j)
}

type FuncLike interface {
	InParams() []*Element
	OutParams() []*Element
}

func firstInParam(f FuncLike, pattern string) *Element {
	if strings.HasPrefix(pattern, "~") && strings.HasSuffix(pattern, "~") {
		// something in the middle

		pattern = strings.TrimLeft(pattern, "~")
		pattern = strings.TrimRight(pattern, "~")
		for _, p := range f.InParams() {
			if strings.Contains(p.BaseType, pattern) {
				return p
			}
		}
	}

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

func firstOutParam(f FuncLike, pattern string) *Element {
	if strings.HasPrefix(pattern, "~") && strings.HasSuffix(pattern, "~") {
		// something in the middle

		pattern = strings.TrimLeft(pattern, "~")
		pattern = strings.TrimRight(pattern, "~")
		for _, p := range f.OutParams() {
			if strings.Contains(p.BaseType, pattern) {
				return p
			}
		}
	}

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
