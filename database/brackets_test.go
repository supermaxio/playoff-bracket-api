package database

// import (
// 	"fmt"
// 	"log"
// 	"reflect"
// 	"testing"

// 	"github.com/supermaxio/nflplayoffbracket/config"
// )

// func TestGetBrackets(t *testing.T) {
// 	config.Setup()
// 	MongoConnect()
// 	log.Println("done")
// 	brackets, err := GetBrackets()
// 	if err != nil {
// 		t.Fail()
// 	}

// 	for _, bracket := range brackets {
// 		usernameTesting := bracket.Username
// 		bracket.Username = "test"
// 		bracket.FinalScoreSum = 0
// 		for _, bracketCompare := range brackets {
// 			if bracketCompare.Username != usernameTesting {
// 				usernameCompareTesting := bracketCompare.Username
// 				bracketCompare.Username = "test"
// 				bracketCompare.FinalScoreSum = 0

// 				if reflect.DeepEqual(bracket, bracketCompare) {
// 					log.Println(fmt.Sprintf("%s and %s share the same bracket", usernameTesting, usernameCompareTesting))
// 					t.Fail()
// 				}
// 			}
// 		}
// 	}
// }
