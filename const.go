package lesphina

type Const struct {
	Name     string `json:"name"`
	RawType  string `json:"raw_type"`
	RawValue string `json:"raw_value"`
}

func (i *Const) GetName() string { return i.Name }
func (i *Const) GetKind() Kind   { return KindConst }
func (i *Const) Json() string    { return marshal(i) }
