package session

import (
	"github.com/djordjev/scURL/data"
	"io/ioutil"
	"encoding/json"
	"github.com/djordjev/scURL/parser"
	"github.com/djordjev/scURL/utils"
)

const FILENAME = "session.scurl"

func loadCurrentSession() data.JsonSession {
	raw, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		// this is OK situation, treat as empty session
		return data.JsonSession{}
	}

	var session data.JsonSession
	json.Unmarshal(raw, &session)

	return session
}

func storeCurrentSession(newSession *data.JsonSession) {
	marshaled, err := json.Marshal(&newSession)
	if err != nil {
		panic("Unable to marshal files")
	}

	err = ioutil.WriteFile(FILENAME, marshaled, 0666)
	if err != nil {
		panic("Unable to store in file")
	}
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
