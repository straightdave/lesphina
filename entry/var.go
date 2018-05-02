package entry

type Var struct {
	Name    string `json:"name"`
	RawType string `json:"raw_type"`
}

func (i *Var) GetName() string { return i.Name }
func (i *Var) GetKind() Kind   { return KindVar }
func (i *Var) Json() string    { return marshal(i) }
