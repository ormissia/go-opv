package go_opv

import "log"

type User struct {
	Name string `go-opv:"ge:0,le:20"`  //Name >=0 && Name <=20
	Age  int    `go-opv:"ge:0,lt:100"` //Age >= 0 && Age < 100
}

func init() {
	//使用默认配置：struct tag名字为"go-opv"，规则与限定值的分隔符为":"
	myVerifier = NewVerifier()
	//初始化一个验证规则：Age字段大于等于0，小于200
	userRequestRules = Rules{
		"Age": []string{myVerifier.Ge("0"), myVerifier.Lt("200")},
	}

	//初始化自定义属性的验证器，并将struct tag名字设为"go-opv"，规则与限定值的分隔符设为":"
	//customVerifier := NewVerifier(SetTagPrefix("myVerifier"), SetSeparator("#"))
}

var myVerifier Verifier
var userRequestRules Rules

func ExampleVerifier_Verify() {
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
