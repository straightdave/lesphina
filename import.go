package lesphina

type Import struct {
	Alias string `json:"alias"`
	Name  string `json:"name"`
}

func (i *Import) GetName() string { return i.Name }
func (i *Import) GetKind() Kind   { return KindImport }
func (i *Import) Json() string    { return marshal(i) }
