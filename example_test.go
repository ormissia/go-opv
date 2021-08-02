package go_opv

import (
	"log"
)

type User struct {
	Name string
	Age  int
}

func init() {
	myVerifier = NewVerifier(SetSeparator("#"))
	UserRequestRules = Rules{
		"Name": {myVerifier.NotEmpty(), myVerifier.Lt("10")},
		"Age":  {myVerifier.Lt("100")},
	}
}

var myVerifier Verifier
var UserRequestRules Rules

func ExampleVerifier_Verify() {
	// ShouldBind(&user) in Gin framework or other generated object
	user := User{
		Name: "Ormissia",
		Age:  90,
	}
	if err := myVerifier.Verify(user, UserRequestRules); err != nil {
		log.Fatal(err)
	}
}
