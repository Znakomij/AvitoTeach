package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/set_key", setKey)
	http.HandleFunc("/get_key", getKey)
	http.HandleFunc("/del_key", delKey)
	http.HandleFunc("/", other)

	err := http.ListenAndServe("0.0.0.0:8089", nil)
	if err != nil {
		panic(err)
	}
}
