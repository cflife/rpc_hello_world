package main

import (
	"fmt"
	"testing"
)

func TestInitApplication(t *testing.T) {

	defer func() {
		if data := recover(); data != nil {
			fmt.Printf("hello, panic: %v \n", data)
		}

		fmt.Print("恢复是在这里继续执行")
	}()

	panic("boom !")

	fmt.Print("即便恢复，也不会继续执行")
}