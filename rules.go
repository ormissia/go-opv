// Copyright (c) 2021 安红豆. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package go_opv

import "fmt"

// Rules define
type Rules map[string][]string

func (verifier verifier) NotEmpty() string {
	return "notEmpty"
}

// Eq ==
func (verifier verifier) Eq(limit string) string {
	return fmt.Sprintf("%s%s%s", eq, verifier.separator, limit)
}

// Ne !=
func (verifier verifier) Ne(limit string) string {
	return fmt.Sprintf("%s%s%s", ne, verifier.separator, limit)
}

// Gt >
func (verifier verifier) Gt(limit string) string {
	return fmt.Sprintf("%s%s%s", gt, verifier.separator, limit)
}

// Lt <
func (verifier verifier) Lt(limit string) string {
	return fmt.Sprintf("%s%s%s", lt, verifier.separator, limit)
}

// Ge >=
func (verifier verifier) Ge(limit string) string {
	return fmt.Sprintf("%s%s%s", ge, verifier.separator, limit)
}

// Le <=
func (verifier verifier) Le(limit string) string {
	return fmt.Sprintf("%s%s%s", le, verifier.separator, limit)
}
