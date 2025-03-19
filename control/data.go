package control

type Data struct {
	ID     int64    `json:"id"`
	Time   int64    `json:"time"`
	Type   string   `json:"type"`
	Values []string `json:"values"`
}
