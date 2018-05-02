package entry

type Struct struct {
	Name    string     `json:"name"`
	Fields  []*Element `json:"fields"`
	RawBody string     `json:"raw_body"`
}

func (i *Struct) GetName() string { return i.Name }
func (i *Struct) GetKind() Kind   { return KindStruct }
func (i *Struct) Json() string    { return marshal(i) }
