package lesphina

// Struct ...
type Struct struct {
	Name    string     `json:"name"`
	Fields  []*Element `json:"fields"`
	RawBody string     `json:"raw_body"`
}

// -- implement Entry interface --

// GetName ...
func (i *Struct) GetName() string { return i.Name }

// GetKind ...
func (i *Struct) GetKind() Kind { return KindStruct }

// JSON ...
func (i *Struct) JSON() string { return marshal(i) }
