# Golang object parameter verifier

[![Go Reference](https://pkg.go.dev/badge/github.com/ormissia/go-opv.svg)](https://pkg.go.dev/github.com/ormissia/go-opv)
![Repository Size](https://img.shields.io/github/repo-size/ormissia/go-opv)
![Contributor](https://img.shields.io/github/contributors/ormissia/go-opv)
![Last Commit](https://img.shields.io/github/last-commit/ormissia/go-opv)
![License](https://img.shields.io/github/license/ormissia/go-opv)
![Open Issues](https://img.shields.io/github/issues/ormissia/go-opv?color=important)
![Open Pull Requests](https://img.shields.io/github/issues-pr/ormissia/go-opv?color=yellowgreen)

> Golang 对象参数验证器

引用方式

```bash
go get github.com/ormissia/go-opv
```

```go
import "github.com/ormissia/go-opv"
```

## 使用示例

```go
package main

import (
	"log"

	go_opv "github.com/ormissia/go-opv"
)

type User struct {
	Name string `go-opv:"ge:0,le:20"`  //Name >=0 && Name <=20
	Age  int    `go-opv:"ge:0,lt:100"` //Age >= 0 && Age < 100
}

func init() {
	//使用默认配置：struct tag名字为"go-opv"，规则与限定值的分隔符为":"
	myVerifier = go_opv.NewVerifier()
	//初始化一个验证规则：Age字段大于等于0，小于200
	userRequestRules = go_opv.Rules{
		"Age": []string{myVerifier.Ge("0"), myVerifier.Lt("200")},
	}
}

var myVerifier go_opv.Verifier
var userRequestRules go_opv.Rules

func main() {
	// ShouldBind(&user) in Gin framework or other generated object
	user := User{
		Name: "ormissia",
		Age:  190,
	}

	//两种验证方式混合,函数参数中传入自定义规则时候会覆盖struct tag上定义的规则
	//根据自定义规则Age >= 0 && Age < 200，Age的值为190，符合规则，验证通过
	if err := myVerifier.Verify(user, userRequestRules); err != nil {
		log.Println(err)
	} else {
		log.Println("pass")
	}

	//只用struct的tag验证
	//根据tag上定义的规则Age >= 0 && Age < 100，Age的值为190，不符合规则，验证不通过
	if err := myVerifier.Verify(user); err != nil {
		log.Println(err)
	} else {
		log.Println("pass")
	}
}
```

验证结果：

```bash
2021/08/09 22:14:43 pass
2021/08/09 22:14:43 Age length or value is illegal: lt:100
```

## 支持创建自定义属性的验证器

```go
//初始化自定义属性的验证器，并将struct tag名字设为"myVerifier"，规则与限定值的分隔符设为"#"
customVerifier := go_opv.NewVerifier(go_opv.SetTagPrefix("myVerifier"), go_opv.SetSeparator("#"))
```

这时候`struct`标签应该写成：

```go
type User struct {
    Name string `myVerifier:"ge#0,le#20"`  //Name >=0 && Name <=20
    Age  int    `myVerifier:"ge#0,lt#100"` //Age >= 0 && Age < 100
}
```

---

## TODO
- ~~写example~~
- 增加中文长度判断支持
- 写测试用例
- 增加`map`验证器
- ~~增加`tag`标记模式~~
- 通过`tag`生成规则的方式，关注性能

---

## 灵感来源
- [gin-vue-admin中实体参数校验方式](https://github.com/flipped-aurora/gin-vue-admin/blob/186ecbf6b8bd5d2ce2b4856de2f0265846483a50/server/utils/validator.go#L107)
- [函数选项模式](https://ormissia.github.io/posts/knowledge/2001-go-partten-1/)
