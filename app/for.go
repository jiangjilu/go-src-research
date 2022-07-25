package app

import (
	"fmt"
	"time"
)

func For() {
	for {
		fmt.Println(time.Now())
		time.Sleep(time.Second)
	}
}
