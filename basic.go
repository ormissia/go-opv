// Copyright (c) 2021 安红豆. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package go_opv

import (
	"errors"
	"reflect"
	"strings"
)

var defaultVerifierOptions = verifierOptions{
	separator: ":",
	conditions: map[string]bool{
		eq: true,
		ne: true,
		gt: true,
		lt: true,
		ge: true,
		le: true,
	},
}

type VerifierOption func(o *verifierOptions)
type verifierOptions struct {
	conditions map[string]bool
	separator  string
}

// SetSeparator Default separator is ":".
func SetSeparator(seq string) VerifierOption {
	return func(o *verifierOptions) {
		o.separator = seq
	}
}

func SwitchEq(sw bool) VerifierOption {
	return func(o *verifierOptions) {
		o.conditions[eq] = sw
	}
}

func SwitchNe(sw bool) VerifierOption {
	return func(o *verifierOptions) {
		o.conditions[ne] = sw
	}
}

func SwitchGt(sw bool) VerifierOption {
	return func(o *verifierOptions) {
		o.conditions[gt] = sw
	}
}

func SwitchLt(sw bool) VerifierOption {
	return func(o *verifierOptions) {
		o.conditions[lt] = sw
	}
}

func SwitchGe(sw bool) VerifierOption {
	return func(o *verifierOptions) {
		o.conditions[ge] = sw
	}
}

func SwitchLe(sw bool) VerifierOption {
	return func(o *verifierOptions) {
		o.conditions[le] = sw
	}
}

type Verifier interface {
	Verify(obj interface{}, rules Rules) (err error)

	NotEmpty() string
	Ne(limit string) string
	Gt(limit string) string
	Lt(limit string) string
	Ge(limit string) string
	Le(limit string) string
}

type verifier struct {
	separator  string
	conditions map[string]bool
}

func NewVerifier(opts ...VerifierOption) Verifier {
	options := defaultVerifierOptions
	for _, opt := range opts {
		opt(&options)
	}
	return verifier{
		separator:  options.separator,
		conditions: options.conditions,
	}
}

func (verifier verifier) Verify(st interface{}, rules Rules) (err error) {
	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st)

	if val.Kind() != reflect.Struct {
		return errors.New("expect struct")
	}
	num := val.NumField()
	//遍历需要验证对象的所有字段
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i)
		val := val.Field(i)
		if len(rules[tagVal.Name]) > 0 {
			for _, v := range rules[tagVal.Name] {
				switch {
				case v == "notEmpty":
					if isEmpty(val) {
						return errors.New(tagVal.Name + " value can not be nil")
					}
				case verifier.conditions[strings.Split(v, verifier.separator)[0]]:
					if !compareVerify(val, v, verifier.separator) {
						return errors.New(tagVal.Name + " length or value is illegal," + v)
					}
				}
			}
		}
	}
	return nil
}
