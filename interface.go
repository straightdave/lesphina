package lesphina

// Interface ...
type Interface struct {
	Name    string             `json:"name"`
	RawBody string             `json:"rawBody"`
	Methods []*InterfaceMethod `json:"methods"`
}

// GetName ...
func (i *Interface) GetName() string { return i.Name }

// GetKind ...
func (i *Interface) GetKind() Kind { return KindInterface }

// JSON ...
func (i *Interface) JSON() string { return marshal(i) }
