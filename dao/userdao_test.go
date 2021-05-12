package dao

import (
	"fmt"
	"testing"
)

func TestGetUserByPhone(t *testing.T) {

	phone, err := GetUserByPhone("18632111763")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(phone)
	}

}
