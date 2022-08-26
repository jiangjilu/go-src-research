package app

import "fmt"

// Block
// golang 中使用大括号{} 包起来的就是一个代码块
// 在这个代码块里面定义的局部变量只在这个代码块中起作用
func Block() {
	i := 0
	{
		i := 1
		fmt.Println("In Code Block: ", i)
	}
	fmt.Println(i)

	// Gin以及Hertz框架中路由组示例中出现的{},也是代码块
	// Simple group: v1
	/*
		v1 := h.Group("/v1")
		{
			// loginEndpoint is a handler func
			v1.POST("/login", loginEndpoint)
		}
	*/

}
