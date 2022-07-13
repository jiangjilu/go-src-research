package app

import "fmt"

func Delete() {
	myMap := map[string]string{
		"Name":    "Jack",
		"Address": "Shenzhen",
	}
	delete(myMap, "Name")
	fmt.Println(myMap)
}

type Person struct {
	Name string
	Age  int
}

func EchoName(person Person) string {
	return person.Name
}
