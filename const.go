package lesphina

// Const ...
type Const struct {
	Name     string `json:"name"`
	RawType  string `json:"raw_type"`
	RawValue string `json:"raw_value"`
}

// -- implement Entry interface --

// GetName ...
func (i *Const) GetName() string { return i.Name }

// GetKind ...
func (i *Const) GetKind() Kind { return KindConst }

// JSON ...
func (i *Const) JSON() string { return marshal(i) }
