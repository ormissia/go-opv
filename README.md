# Golang object parameter verifier

[![Go Reference](https://pkg.go.dev/badge/github.com/ormissia/go-opv.svg)](https://pkg.go.dev/github.com/ormissia/go-opv)
![Repository Size](https://img.shields.io/github/repo-size/ormissia/go-opv)
![Contributor](https://img.shields.io/github/contributors/ormissia/go-opv)
![Last Commit](https://img.shields.io/github/last-commit/ormissia/go-opv)
![License](https://img.shields.io/github/license/ormissia/go-opv)
![Open Issues](https://img.shields.io/github/issues/ormissia/go-opv?color=important)
![Open Pull Requests](https://img.shields.io/github/issues-pr/ormissia/go-opv?color=yellowgreen)

> Golang 对象参数验证器

引用
```bash
go get github.com/ormissia/go-opv
```

```go
import "go_opv"
```

## TODO
- ~~写example~~
- 增加中文长度判断支持
- 写测试用例
- 增加`map`验证器
- 增加`tag`标记模式
- 增加通过函数参数的方式，实现自定义规则校验

---

使用示例

```go
package main

import (
	"github.com/ormissia/go-opv"
	"log"
)

type User struct {
	Name string
	Age  int
}

func init() {
	myVerifier = go_opv.NewVerifier(go_opv.SetSeparator("#"))
	userRequestRules = go_opv.Rules{
		"Name": {myVerifier.NotEmpty(), myVerifier.Lt("10")},
		"Age":  {myVerifier.Lt("100")},
	}
}

var myVerifier go_opv.Verifier
var userRequestRules go_opv.Rules

func main() {
	// ShouldBind(&user) in Gin framework or other generated object
	user := User{
		Name: "Ormissia",
		Age:  900,
	}
	if err := myVerifier.Verify(user, userRequestRules); err != nil {
		log.Fatal(err)
	}
}
```

```bash
2021/08/09 16:57:36 Age length or value is illegal,lt#100
```

由于当前校验对象`Age`值为900，不符合规则，故`err`值返回错误信息


---

## 灵感来源
- [gin-vue-admin中实体参数校验方式](https://github.com/flipped-aurora/gin-vue-admin/blob/186ecbf6b8bd5d2ce2b4856de2f0265846483a50/server/utils/validator.go#L107)
- [函数选项模式](https://ormissia.github.io/posts/knowledge/2021-07-22/)