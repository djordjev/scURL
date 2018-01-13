package main

import (
	"github.com/djordjev/scURL/parser"
	"github.com/djordjev/scURL/session"
	"github.com/djordjev/scURL/request"
)

func main() {
	parsedOperations := parser.ParseArguments()
	endpoint, requestType, body := parser.ParseRequestArguments()

	currentSession := session.UpdateCurrentSession(parsedOperations)

	request.SendRequest(currentSession, endpoint, body, requestType)
}
