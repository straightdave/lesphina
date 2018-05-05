package lesphina

type Interface struct {
	Name    string             `json:"name"`
	RawBody string             `json:"raw_body"`
	Methods []*InterfaceMethod `json:"methods"`
}

func (i *Interface) GetName() string { return i.Name }
func (i *Interface) GetKind() Kind   { return KindInterface }
func (i *Interface) Json() string    { return marshal(i) }

type InterfaceMethod struct {
	Name    string     `json:"name"`
	RawType string     `json:"raw_type"`
	In      []*Element `json:"in,omitempty"`
	Out     []*Element `json:"out,omitempty"`
}

func (i *InterfaceMethod) GetName() string                       { return i.Name }
func (i *InterfaceMethod) GetKind() Kind                         { return KindInterfaceMethod }
func (i *InterfaceMethod) Json() string                          { return marshal(i) }
func (i *InterfaceMethod) InParams() []*Element                  { return i.In }
func (i *InterfaceMethod) OutParams() []*Element                 { return i.Out }
func (i *InterfaceMethod) FirstInParam(pattern string) *Element  { return firstInParam(i, pattern) }
func (i *InterfaceMethod) FirstOutParam(pattern string) *Element { return firstOutParam(i, pattern) }
