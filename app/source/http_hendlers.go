package main

import (
	"github.com/gomodule/redigo/redis"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func setKey(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Parsing the request body
	reg, err := regexp.Compile("[^a-zA-Z0-9:]+")
	if err != nil {
		panic(err)
	}
	resultString := reg.ReplaceAllString(string(body), "")
	key := strings.Split(resultString, ":")[0]
	value := strings.Split(resultString, ":")[1]
	conn, err := redis.Dial("tcp", "redis:6379")

	_, err = conn.Do("SET", key, value)
	if err != nil {
		panic(err)
	} else {
		w.Write([]byte(key + " : " + value + " successfully set"))
	}

	conn.Close()
}

func getKey(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	conn, err := redis.Dial("tcp", "redis:6379")
	value, err := redis.String(conn.Do("GET", key))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 NOT FOUND"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(value))
	}
	conn.Close()
}

func delKey(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9:]+")
	if err != nil {
		panic(err)
	}
	resultString := reg.ReplaceAllString(string(body), "")
	key := strings.Split(resultString, ":")[1]

	conn, err := redis.Dial("tcp", "redis:6379")
	conn.Do("DEL", key)
	if err != nil {
		panic(err)
	} else {
		w.Write([]byte("successfully deleted"))
	}
	conn.Close()
}

func other(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("403 тута"))
}
