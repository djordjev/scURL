package session

import (
	"fmt"
	"github.com/djordjev/scURL/data"
	"io/ioutil"
	"encoding/json"
	"os"
	"github.com/djordjev/scURL/parser"
	"github.com/djordjev/scURL/utils"
)

func loadCurrentSession() data.JsonSession {
	raw, err := ioutil.ReadFile("./session")
	if err != nil {
		// this is OK situation, treat as empty session
		return data.JsonSession{}
	}

	var session data.JsonSession
	json.Unmarshal(raw, &session)

	return session
}

func storeCurrentSession(newSession *data.JsonSession) bool {
	err := os.Truncate("./session", 0)
	if err != nil {
		fmt.Println("Unable to truncate file")
		return false
	}

	marshaled, err := json.Marshal(&newSession)
	if err != nil {
		fmt.Println("Unable to marshal files")
		return false
	}

	err = ioutil.WriteFile("output.json", marshaled, 0644)
	if err != nil {
		fmt.Println("Unable to store in file")
		return false
	}
	return true
}

func UpdateCurrentSession(parsedOperations []parser.ParseResult) data.JsonSession {
	currentSession := loadCurrentSession()

	for _, op := range parsedOperations {
		switch op.Subject {
		case parser.SubjectSession: {
			if op.Operation == parser.OperationAdd {
				currentSession = data.JsonSession{}
			}
		}
		case parser.SubjectBaseApi: {
			if op.Operation == parser.OperationAdd {
				currentSession.BaseApi = op.Key
			} else if op.Operation == parser.OperationRemove {
				currentSession.BaseApi = ""
			}
		}
		case parser.SubjectCookie: {
			if op.Operation == parser.OperationAdd {
				currentSession.Cookies = append(currentSession.Cookies, data.Pair{ Name: op.Key, Value: op.Value })
			} else if op.Operation == parser.OperationRemove {
				currentSession.Cookies = utils.RemoveByKey(currentSession.Cookies, op.Key)
			}
		}
		case parser.SubjectHeader: {
			if op.Operation == parser.OperationAdd {
				currentSession.Headers = append(currentSession.Headers, data.Pair{ Name: op.Key, Value: op.Value })
			} else if op.Operation == parser.OperationRemove {
				currentSession.Headers = utils.RemoveByKey(currentSession.Headers, op.Key)
			}
		}
		}
	}

	storeCurrentSession(&currentSession)

	return currentSession
}
