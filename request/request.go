package request

import (
	"github.com/djordjev/scURL/data"
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

func SendRequest(session data.JsonSession, endpoint, body, requestType string) {
	client := &http.Client{}

	request, _ := http.NewRequest(requestType, session.BaseApi + endpoint, strings.NewReader(body))

	// Add headers
	for _, header := range session.Headers {
		request.Header.Add(header.Name, header.Value)
	}

	// Add cookies
	request.Header.Set("Cookie", buildCookies(session))

	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	displayResponse(response)
}

func buildCookies (session data.JsonSession) string {
	var allCookies string

	for _, cookie := range session.Cookies {
		allCookies += cookie.Name + "=" + cookie.Value + ";"
	}
	return allCookies
}

func displayResponse (response *http.Response) {
	bodyBytes, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}

	strBody := string(bodyBytes)

	fmt.Println(strBody)
}


