package request

import (
	"github.com/djordjev/scURL/data"
	"fmt"
)

func SendRequest(session data.JsonSession, endpoint, body, requestType string) {
	fmt.Println(session)
	fmt.Println("Sending request %s to endpoint %s with body %s ", requestType, endpoint, body)
}
