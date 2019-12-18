package lesphina

// Const is the variables marked as `const` in source.
type Const struct {
	Name     string `json:"name"`
	RawType  string `json:"rawType"`
	RawValue string `json:"rawValue"`
}

// GetName ...
func (i *Const) GetName() string { return i.Name }

// GetKind ...
func (i *Const) GetKind() Kind { return KindConst }

// JSON ...
func (i *Const) JSON() string { return marshal(i) }
