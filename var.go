package lesphina

type Var struct {
	Name     string `json:"name"`
	RawType  string `json:"raw_type"`
	RawValue string `json:"raw_value"`
	IsFunc   bool   `json:"is_func,omitempty"`
}

func (i *Var) GetName() string { return i.Name }
func (i *Var) GetKind() Kind   { return KindVar }
func (i *Var) Json() string    { return marshal(i) }
