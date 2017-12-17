package parser

import (
	"flag"
	"strings"
	"github.com/djordjev/scURL/data"
)

const (
	AddOperation = iota
	RemoveOperation
	ClearOperation
)

type ParseResult struct {
	Op     data.SessionOperation
	Action int
}

func ParseArguments() []ParseResult {
	result := make([]ParseResult, 0)

	// flag definitions
	newSession := flag.Bool("n", false, "requires -n flag")
	addCookie := flag.String("ac", "", "requires -ac flag")
	removeCookie := flag.String("rc", "", "requires -rc flag")
	addHeader := flag.String("ah", "", "requires -ah flag")
	removeHeader := flag.String("rh", "", "requires -ar flag")

	flag.Parse()

	if *newSession == true {
		result = append(result, ParseResult{ nil, ClearOperation})
	}

	if *addCookie != "" {
		split := strings.Split(*addCookie, "=")
		result = append(result, ParseResult{
			Op:   data.SessionOperation{
				OpName:"cookie",
				Key: split[0],
				Value: split[1]},
			Action: AddOperation,
			})
	}

	if *removeCookie != "" {
		result = append(result, ParseResult{
			Op:   data.SessionOperation{
				OpName:"cookie",
				Key: *removeCookie,
				Value: ""},
			Action: RemoveOperation,
		})
	}

	if *addHeader != "" {
		split := strings.Split(*addHeader, "=")
		result = append(result, ParseResult{
			Op:   data.SessionOperation{
				OpName:"header",
				Key: split[0],
				Value: split[1]},
			Action: AddOperation,
		})
	}

	if *removeHeader != "" {
		result = append(result, ParseResult{
			Op:   data.SessionOperation{
				OpName:"header",
				Key: *removeHeader,
				Value: ""},
			Action: RemoveOperation,
		})
	}

	// TODO sort here
	return result
}

