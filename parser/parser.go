package parser

import (
	"flag"
	"strings"
	"sort"
	"fmt"
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

type args struct {
	newSession bool
	addCookie string
	removeCookie string
	addHeader string
	removeHeader string
	addBaseApi string
	removeBaseApi bool
	endpoint string
	requestType string
	body string
}

var programArgs args

func init() {
	// flag definitions
	newSession := flag.Bool("n", false, "requires -n flag")
	addCookie := flag.String("ac", "", "requires -ac flag")
	removeCookie := flag.String("rc", "", "requires -rc flag")
	addHeader := flag.String("ah", "", "requires -ah flag")
	removeHeader := flag.String("rh", "", "requires -ar flag")
	addBaseApi := flag.String("aa", "", "requires -aa flag")
	removeBaseApi := flag.Bool("ra", false, "requires -ra flag")
	endpointFlag := flag.String("ep", "", "endpoint for request")
	requestTypeFlag := flag.String("X", "GET", "request type")
	bodyData := flag.String("b", "", "request body")

	flag.Parse()

	programArgs = args{
		newSession: *newSession,
		addCookie: *addCookie,
		removeCookie: *removeCookie,
		addHeader: *addHeader,
		removeHeader: *removeHeader,
		addBaseApi: *addBaseApi,
		removeBaseApi: *removeBaseApi,
		endpoint: *endpointFlag,
		requestType: *requestTypeFlag,
		body: *bodyData,
	}
}

func makeParseResult (operation int, subject int, key string, value string) ParseResult {
	return ParseResult{ operation, subject, key, value }
}

func ParseArguments() []ParseResult {
	result := make([]ParseResult, 0)

	if programArgs.newSession == true {
		result = append(result, makeParseResult(OperationAdd, SubjectSession, "", ""))
	}

	if programArgs.addCookie != "" {
		split := strings.Split(programArgs.addCookie, "=")
		result = append(result, makeParseResult(OperationAdd, SubjectCookie, split[0], split[1]))
	}

	if programArgs.removeCookie != "" {
		result = append(result, makeParseResult(OperationRemove, SubjectCookie, programArgs.removeCookie, ""))
	}

	if programArgs.addHeader != "" {
		split := strings.Split(programArgs.addHeader, "=")
		result = append(result, makeParseResult(OperationAdd, SubjectHeader, split[0], split[1]))
	}

	if programArgs.removeHeader != "" {
		result = append(result, makeParseResult(OperationRemove, SubjectHeader, programArgs.removeHeader, ""))
	}

	if programArgs.addBaseApi != "" {
		result = append(result, makeParseResult(OperationAdd, SubjectBaseApi, programArgs.addBaseApi, ""))
	}

	if programArgs.removeBaseApi == true {
		result = append(result, makeParseResult(OperationRemove, SubjectBaseApi, "", ""))
	}

	sort.Sort(BySubject(result))
	// TODO failsafe check if some of operations are conflicting
	return result
}

func ParseRequestArguments () (endpoint, requestType, body string) {


	flag.Parse()

	endpoint = programArgs.endpoint
	requestType = strings.ToUpper(programArgs.requestType)
	body = programArgs.body

	if !checkIsRequestTypeValid(requestType) {
		panic(fmt.Sprintf("Request type %s is not valid request type", requestType))
	}

	return
}

func checkIsRequestTypeValid(reqType string) bool {
	var validTypes = [...]string{"POST", "GET", "UPDATE", "DELETE", "PUT"}

	for _, elem := range validTypes {
		if elem == reqType {
			return true
		}
	}

	return false
}

