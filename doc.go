// Copyright (c) 2021 安红豆. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
	Package go_opv implements verify go object parameter using custom rules.

	Allow rules:
		eq : ==
		ne : !=
		gt : >
		lt : <
		ge : >=
		le : <=

	Example:
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
*/
package go_opv
