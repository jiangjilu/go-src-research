package app

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func Json() {
	article := Article{}
	err := json.Unmarshal([]byte("{\"title\":\"标题\",\"content\":\"内容\"}"), &article)
	if err != nil {

	}
	fmt.Println(article)
}
