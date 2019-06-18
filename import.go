package lesphina

// Import ...
type Import struct {
	Alias string `json:"alias"`
	Name  string `json:"name"` // full import path
}

func (i *Import) GetName() string { return i.Name }
func (i *Import) GetKind() Kind   { return KindImport }
func (i *Import) JSON() string    { return marshal(i) }
