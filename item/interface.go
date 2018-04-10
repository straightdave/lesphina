package item

type Interface struct {
	Name string `json:"name"`
	// Methods []*Function `json:"methods"`
	RawBody string `json:"raw_body"`
}
