package parser

type BySubject []ParseResult

func (a BySubject) Len() int {
	return len(a)
}

func (a BySubject) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a BySubject) Less(i, j int) bool {
	return a[i].Subject < a[j].Subject
}
