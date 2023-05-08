package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValue(t *testing.T) {
	valueInt := reflect.ValueOf(1)
	valueString := reflect.ValueOf("hello")
	valueUser := reflect.ValueOf(User{
		Id:   1,
		Name: "Name",
		age:  10,
	})
	fmt.Println(valueInt)
	fmt.Println(valueString)
	fmt.Println(valueUser)
}
