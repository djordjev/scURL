package main

import (
	"github.com/djordjev/scURL/parser"
	"fmt"
	"github.com/djordjev/scURL/session"
)

func main() {
	parsedOperations := parser.ParseArguments()
	currentSession := session.LoadCurrentSession()

	for _, op := range parsedOperations {
		fmt.Println(fmt.Sprintf("Get operation %d", op.Operation))
	}
}
