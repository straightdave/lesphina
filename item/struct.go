package item

type Struct struct {
	Name    string     `json:"name"`
	Fields  []*Element `json:"fields"`
	RawBody string     `json:"raw_body"`
}
