package main

import (
	"github.com/jiangjilu/go-src-research/app"
)

func main() {
	// 操作数据库
	app.Gorm()

	// delete() 删除 map 中的条目
	app.Delete()

	//strings := []string{"Hello", " ", "World", "!"}
	//
	//for i, i2 := range strings {
	//	fmt.Println(i,i2)
	//}

	/*	for i := 0; i < len(strings); i++ {
			fmt.Println(strings[i])
		}
	*/
	//mySlice := append([]string{"hello "}, []string{"world","!"}...)
	//mySlice := append([]string{"hello "}, "world")
	//fmt.Println(mySlice)
	//
	//mySliceDst := []string{"hello 1","","-","Boy", "Girl"}
	//n := copy(
	//	mySliceDst,
	//	mySlice,
	//)
	//fmt.Println(n)
	//fmt.Println(mySliceDst)
	//fmt.Println(mySlice)

	/*p1 := app.Person{
		Name: "Xi",
		Age:  0,
	}
	p1.Age = 9
	fmt.Println(app.EchoName(p1))
	arr := []app.Person{ p1 }

	peoples := arr
	for i := 1; i <= 10; i++ {
		p1.Age++
		peoples = append(peoples, p1)
	}
	fmt.Println(peoples)*/

	//myComplex := []complex64{complex(1,2),complex(1,3)}
	//fmt.Println(myComplex)

	//fmt.Println(complex(2, 200) + complex(1, 100))

	//collection := []bool{true,false}
	//collection := []bool{true,false,true,false}
	//
	//for i := 0; i < len(collection); i++ {
	//	fmt.Println(i,collection[i])
	//}

	//for i, i2 := range collection {
	//	go println(i)
	//	go println(i2)
	//}

	//time.Sleep(time.Second)

}
