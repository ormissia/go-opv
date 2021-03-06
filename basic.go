// Copyright (c) 2021 安红豆. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package go_opv

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var defaultVerifierOptions = verifierOptions{
	tagPrefix: defaultTagPrefix,
	separator: defaultSeparator,
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
	tagPrefix  string
	separator  string
	conditions map[string]bool
}

// SetTagPrefix Default separator is "go-opv".
func SetTagPrefix(seq string) VerifierOption {
	return func(o *verifierOptions) {
		o.separator = seq
	}
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
	Verify(obj interface{}, rules ...Rules) (err error)

	Ne(limit string) string
	Gt(limit string) string
	Lt(limit string) string
	Ge(limit string) string
	Le(limit string) string
}

type verifier struct {
	tagPrefix  string
	separator  string
	conditions map[string]bool
}

func NewVerifier(opts ...VerifierOption) Verifier {
	options := defaultVerifierOptions
	for _, opt := range opts {
		opt(&options)
	}
	return verifier{
		tagPrefix:  options.tagPrefix,
		separator:  options.separator,
		conditions: options.conditions,
	}
}

func (verifier verifier) Verify(st interface{}, rules ...Rules) (err error) {
	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st)

	fieldName := make(map[string]string)
	var conditions Rules
	if len(rules) > 0 {
		conditions = rules[0]
	} else {
		conditions = Rules{}
	}

	if val.Kind() != reflect.Struct {
		return errors.New("expect struct")
	}
	num := val.NumField()
	// 遍历需要验证对象的所有字段
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i)
		field := val.Field(i)
		fieldName[tagVal.Name] = tagVal.Name

		if len(conditions[tagVal.Name]) == 0 {
			// 当传进来的规则中没有该字段的规则，则去该字段的struct tag中查找
			// `go-opv:"ne:0,eq:10"`
			// conditionsStr = "ne:0,eq:10"
			if conditionsStr, ok := tagVal.Tag.Lookup(verifier.tagPrefix); ok && conditionsStr != "" {
				conditionStrs := strings.Split(conditionsStr, ",")
				// 判断第一个是条件还是自定义的字段名
				if len(conditionStrs) > 0 && len(strings.Split(conditionStrs[0], verifier.separator)) == 1 {
					fieldName[tagVal.Name] = conditionStrs[0]
					// 如果第一个字段是字段名的话，将第一个去掉，剩下的作为条件
					conditionStrs = conditionStrs[1:]
				}
				conditions[tagVal.Name] = conditionStrs
			} else {
				// 如果tag也没有定义则去校验下一个字段
				continue
			}
		} else {
			// 当传进来的自定义规则有该字段的规则，判断第一个是否是字段名
			conditionStrs := conditions[tagVal.Name]
			// 判断第一个是条件还是自定义的字段名
			if len(conditionStrs) > 0 && len(strings.Split(conditionStrs[0], verifier.separator)) == 1 {
				fieldName[tagVal.Name] = conditionStrs[0]
				// 如果第一个字段是字段名的话，将第一个去掉，剩下的作为条件
				conditions[tagVal.Name] = conditionStrs[1:]
				fmt.Println(conditions[tagVal.Name])
			}
		}

		for _, v := range conditions[tagVal.Name] {
			switch {
			case verifier.conditions[strings.Split(v, verifier.separator)[0]]:
				if !compareVerify(field, v, verifier.separator) {
					return errors.New(fmt.Sprintf("%s length or value is illegal: %s", fieldName[tagVal.Name], v))
				}
			default:
				condition := strings.Split(v, verifier.separator)[0]
				return errors.New(fmt.Sprintf("Unsupported condition: %s", condition))
			}
		}
	}
	return nil
}
