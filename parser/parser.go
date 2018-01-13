package parser

import (
	"flag"
	"strings"
	"sort"
)

const (
	OperationAdd = iota
	OperationRemove
)

const (
	SubjectSession = iota
	SubjectBaseApi
	SubjectCookie
	SubjectHeader
)

type ParseResult struct {
	Operation int
	Subject int
	Key string
	Value string
}

func makeParseResult (operation int, subject int, key string, value string) ParseResult {
	return ParseResult{ operation, subject, key, value }
}

func ParseArguments() []ParseResult {
	result := make([]ParseResult, 0)

	// flag definitions
	newSession := flag.Bool("n", false, "requires -n flag")
	addCookie := flag.String("ac", "", "requires -ac flag")
	removeCookie := flag.String("rc", "", "requires -rc flag")
	addHeader := flag.String("ah", "", "requires -ah flag")
	removeHeader := flag.String("rh", "", "requires -ar flag")
	addBaseApi := flag.String("aa", "", "requires -aa flag")
	removeBaseApi := flag.Bool("ra", false, "requires -ra flag")

	flag.Parse()

	if *newSession == true {
		result = append(result, makeParseResult(OperationAdd, SubjectSession, "", ""))
	}

	if *addCookie != "" {
		split := strings.Split(*addCookie, "=")
		result = append(result, makeParseResult(OperationAdd, SubjectCookie, split[0], split[1]))
	}

	if *removeCookie != "" {
		result = append(result, makeParseResult(OperationRemove, SubjectCookie, *removeCookie, ""))
	}

	if *addHeader != "" {
		split := strings.Split(*addHeader, "=")
		result = append(result, makeParseResult(OperationAdd, SubjectHeader, split[0], split[1]))
	}

	if *removeHeader != "" {
		result = append(result, makeParseResult(OperationRemove, SubjectHeader, *removeHeader, ""))
	}

	if *addBaseApi != "" {
		result = append(result, makeParseResult(OperationAdd, SubjectBaseApi, *addBaseApi, ""))
	}

	if *removeBaseApi == true {
		result = append(result, makeParseResult(OperationRemove, SubjectBaseApi, "", ""))
	}

	sort.Sort(BySubject(result))
	// TODO failsafe check if some of operations are conflicting
	return result
}

func ParseRequestArguments () (endpoint, requestType, body string) {
	endpointFlag := flag.String("e", "", "endpoint for request")
	requestTypeFlag := flag.String("X", "GET", "request type")
	bodyData := flag.String("b", "", "request body")

	endpoint = *endpointFlag
	requestType = *requestTypeFlag
	body = *bodyData

	return
}

