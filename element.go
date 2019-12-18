package lesphina

import (
	"regexp"
	"strings"
)

var regexJSONFieldName = regexp.MustCompile(`json:"(.+?)"`)

// Element is the very often used entry declaration.
// You see it everywhere in go code like "<name> <type>".
type Element struct {
	Name       string `json:"name"`
	RawType    string `json:"rawType"`
	IsPointer  bool   `json:"isPointer"`
	IsSlice    bool   `json:"isSlice"`
	IsVariadic bool   `json:"isVariadic"`
	IsMap      bool   `json:"isMap"`
	BaseType   string `json:"baseType"`
	KeyType    string `json:"keyType,omitempty"`
	ValueType  string `json:"valueType,omitempty"`
	RawTag     string `json:"rawTag,omitempty"` // used in struct definition
}

// GetName ...
func (i *Element) GetName() string { return i.Name }

// GetKind ..
func (i *Element) GetKind() Kind { return KindElement }

// JSON ...
func (i *Element) JSON() string { return marshal(i) }

// JSONFieldName returns the field name of such element when used in a JSON.
// e.g `json:"xxx,bbbbbb" bson:"yyy,cccccc"`
func (i *Element) JSONFieldName() string {
	if i.RawTag == "" {
		return ""
	}

	tag := strings.TrimLeft(i.RawTag, "`")
	tag = strings.TrimRight(tag, "`")

	m := regexJSONFieldName.FindStringSubmatch(tag)
	if len(m) != 2 {
		return ""
	}
	return strings.Split(m[1], ",")[0]
}
