package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Serve 开启 http 服务
func Serve() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// Request 请求一个网页
func Request() {
	go Serve()

	client := http.Client{}
	res, _ := client.Get("http://localhost:8080")
	c, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(c))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
