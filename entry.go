package lesphina

import (
	"encoding/json"
	"strings"
)

// Kind represents different entry types in lesphina.
type Kind int

const (
	KindImport Kind = iota
	KindElement
	KindConst
	KindVar
	KindInterface
	KindStruct
	KindFunction
	KindInterfaceMethod
)

// Entry is the common stuff of all types in lesphina.
type Entry interface {
	GetName() string
	GetKind() Kind
}

func marshal(i Entry) string {
	j, _ := json.MarshalIndent(i, "", "    ")
	return string(j)
}

// FuncLike stands for both functions and interface methods.
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
