package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("1")
	io.WriteString(w, "Hello world again!\n")
	resp, err := http.Get("http://web.cloud66.local/counts")
	if err != nil {
		// handle error
		io.WriteString(w, "Count from web.cloud66.local is: <<unavailable>>")
		return
	}
	log.Println("2")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		io.WriteString(w, "Count from web.cloud66.local is: <<unavailable>>")
		return
	}
	io.WriteString(w, "Count from web.cloud66.local is: "+string(body))
	log.Println("3")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
