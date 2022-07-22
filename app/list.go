package app

import (
	"container/list"
	"fmt"
)

func List() {
	l := list.New()

	z := l.PushBack("张飞")
	l.InsertBefore("关羽", z)
	l.InsertAfter("赵云", z)
	l.PushFront("刘备")

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
