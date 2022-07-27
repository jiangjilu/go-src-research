package app

import "fmt"

func Iota() {
	const (
		x = iota // 0
		y        // 1
		z        // 2
	)
	const (
		m = iota // 0
		p = 100  // 插队
		_        // 跳过计数
		n = iota // 3
	)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	fmt.Println("m:", m)
	fmt.Println("p:", p)
	fmt.Println("n:", n)

}
