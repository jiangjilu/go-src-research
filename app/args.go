package app

import "fmt"

func VarArgs(s string, a ...interface{}) {
	fmt.Println(s)
	fmt.Println(a)
}
