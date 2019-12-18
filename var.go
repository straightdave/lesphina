package lesphina

// Var is package level vars.
type Var struct {
	Name     string `json:"name"`
	RawType  string `json:"rawType"`
	RawValue string `json:"rawValue"`
	IsFunc   bool   `json:"isFunc,omitempty"`
}

// GetName ...
func (i *Var) GetName() string { return i.Name }

// GetKind ...
func (i *Var) GetKind() Kind { return KindVar }

// JSON ...
func (i *Var) JSON() string { return marshal(i) }
