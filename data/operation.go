package data

type SessionOperation struct {
	OpName string `json: "entry"`
	Key string `json:"key"`
	Value string `json:"value"`
}
