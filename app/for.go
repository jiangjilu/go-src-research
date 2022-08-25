package app

import (
	"fmt"
	"time"
)

func For() {
	i := 0
	for {
		if i > 3 {
			break
		}

		fmt.Println(time.Now())
		time.Sleep(time.Second)

		i++
	}
}
