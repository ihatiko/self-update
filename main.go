package main

import (
	"fmt"
	"github.com/minio/selfupdate"
	"net/http"
)

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = selfupdate.Apply(resp.Body, selfupdate.Options{})
	if err != nil {
		// error handling
	}
	return err
}

func main() {
	err := doUpdate("https://github.com/ihatiko/self-update")
	if err != nil {
		panic(err)
	}
	fmt.Println("123")
}
