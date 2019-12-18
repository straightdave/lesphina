package lesphina

// Import ...
type Import struct {
	Name  string `json:"name"` // full import path
	Alias string `json:"alias"`
}

// GetName ...
func (i *Import) GetName() string { return i.Name }

// GetKind ...
func (i *Import) GetKind() Kind { return KindImport }

// JSON ...
func (i *Import) JSON() string { return marshal(i) }
