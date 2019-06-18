package lesphina

import (
	"strings"
)

// Element stands for the very often used entry declaration.
// You see it every where in go code: "<name> <type>"
type Element struct {
	Name       string `json:"name"`
	RawType    string `json:"raw_type"`
	IsPointer  bool   `json:"is_pointer"`
	IsSlice    bool   `json:"is_slice"`
	IsVariadic bool   `json:"is_variadic"`
	IsMap      bool   `json:"is_map"`
	BaseType   string `json:"base_type"`
	KeyType    string `json:"key_type,omitempty"`
	ValueType  string `json:"value_type,omitempty"`
	RawTag     string `json:"raw_tag,omitempty"` // used in struct definition
}

// -- implement Entry interface --

// GetName ...
func (i *Element) GetName() string { return i.Name }

// GetKind ..
func (i *Element) GetKind() Kind { return KindElement }

// JSON ...
func (i *Element) JSON() string { return marshal(i) }

// JSONFieldName returns the field name of element in json data.
// this comes out from RawTag and only focus on json
func (i *Element) JSONFieldName() string {
	if i.RawTag == "" {
		return ""
	}

	tag := strings.TrimLeft(i.RawTag, "`")
	tag = strings.TrimRight(tag, "`")

	// json:"xxx,bbbbbb" bson:"yyy,cccccc" ...
	m := rJSONFieldName.FindStringSubmatch(tag)
	if len(m) != 2 {
		return ""
	}
	return strings.Split(m[1], ",")[0]
}
