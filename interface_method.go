package lesphina

// InterfaceMethod ...
type InterfaceMethod struct {
	Name    string     `json:"name"`
	RawType string     `json:"rawType"`
	In      []*Element `json:"in,omitempty"`
	Out     []*Element `json:"out,omitempty"`
}

// GetName ...
func (i *InterfaceMethod) GetName() string { return i.Name }

// GetKind ...
func (i *InterfaceMethod) GetKind() Kind { return KindInterfaceMethod }

// JSON ...
func (i *InterfaceMethod) JSON() string { return marshal(i) }

// InParams ...
func (i *InterfaceMethod) InParams() []*Element { return i.In }

// OutParams ...
func (i *InterfaceMethod) OutParams() []*Element { return i.Out }

// FirstInParam ...
func (i *InterfaceMethod) FirstInParam(pattern string) *Element { return firstInParam(i, pattern) }

// FirstOutParam ...
func (i *InterfaceMethod) FirstOutParam(pattern string) *Element { return firstOutParam(i, pattern) }
