// Copyright (c) 2021 安红豆. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
	Package go_opv implements verify go object parameter using customer rules.

	Allow rules:
		eq : ==
		ne : !=
		gt : >
		lt : <
		ge : >=
		le : <=

	Example:
		var verifier = go_opv.NewVerifier()

		var aaa = go_opv.Rules{
			"Name": {go_opv.NotEmpty(), go_opv.Lt("20")},
			"Age":  {go_opv.Lt("100")},
		}

		type User struct {
			Name string
			Age  int
		}

		func Example() {
			// ShouldBind(&user) in Gin framework or other generated object
			user := User{
				Name: "Ormissia",
				Age:  90,
			}
			if err := verifier.Verify(user, aaa); err != nil {
				log.Fatal(err)
			}
		}

*/
package go_opv
