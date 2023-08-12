package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Id   int
	Name string
	age  int
}

// TestType 通过TypeOf()得到Type类型
func TestType(t *testing.T) {
	typeInt := reflect.TypeOf(1)
	typeString := reflect.TypeOf("hello")
	fmt.Println(typeInt)
	fmt.Println(typeString)

	typeUser := reflect.TypeOf(&User{})
	fmt.Println(typeUser)
	fmt.Println(typeUser.Kind())
	fmt.Println(typeUser.Elem().Kind())
}

// TestPointerType 指针Type转为非指针Type
func TestPointerType(t *testing.T) {
	typePointerUser := reflect.TypeOf(&User{})
	// *reflect.User
	fmt.Println(typePointerUser)
	typeStructUser := reflect.TypeOf(User{})
	// reflect.User
	fmt.Println(typeStructUser)
	// true
	fmt.Println(typePointerUser.Elem() == typeStructUser)
}

// TestGetStructFiled 获取struct成员变量的信息
func TestGetStructFiled(t *testing.T) {
	typeUser := reflect.TypeOf(User{})
	// NumField returns a programstructure type's field count.
	// It panics if the type's Kind is not Struct.
	// 返回结构体的字段数量，必须是Struct类型
	fieldNum := typeUser.NumField()
	for i := 0; i < fieldNum; i++ {
		structFiled := typeUser.Field(i)
		fmt.Printf("%d %s offset:%d anonymous:%t type:%s exported:%t json tag:%s\n", i,
			structFiled.Name,            // 变量名称
			structFiled.Offset,          // 相对于结构体首地址的内存偏移量，string类型会占据16个字节
			structFiled.Anonymous,       // 是否为匿名成员
			structFiled.Type,            // 数据类型，reflect.Type类型
			structFiled.IsExported(),    // 包外是否可见（即是否以大写字母开头）
			structFiled.Tag.Get("json")) // 获取成员变量后面``里面定义的tag
	}

	// 通过FieldByName获取Field
	if nameField, ok := typeUser.FieldByName("Name"); ok {
		fmt.Printf("Name is exported %t\n", nameField.IsExported())
	}
	// 根据FieldByIndex获取Field
	thirdField := typeUser.FieldByIndex([]int{2}) //参数是个slice，因为有struct嵌套的情况
	fmt.Printf("third field name %s\n", thirdField.Name)
}
