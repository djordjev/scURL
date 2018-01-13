package data

type Pair struct {
	Name string `json: "key"`
	Value string `json: "value"`
}

type JsonSession struct {
	BaseApi string `json: "base_api"`
	Headers []Pair `json: "headers"`
	Cookies []Pair `json: "cookies"`
}