package session

import (
	"fmt"
	"github.com/djordjev/scURL/data"
	"os"
	"log"
)

func Clear() {
	fmt.Println("Clear")
}

func AddSessionOperation(operation string, name string, value string) {
	fmt.Println(fmt.Sprintf("Adding header %s with value %s", name, value))
}

func RemoveSessionOperation(operation string, name string) {
	fmt.Println(fmt.Sprintf("Removing header %s", name))
}

func LoadCurrentSession() []data.SessionOperation {
	//file, err := os.Open("session")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//data := make([]byte, 100)

}
