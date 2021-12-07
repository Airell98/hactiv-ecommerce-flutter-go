package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func CheckAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Println(url, "is down")
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	_ = bodyBytes

	if err != nil {
		fmt.Println("error occured", err)
	}

	splitUrl := strings.Split(url, "//")[1]

	file := "coba/" + splitUrl + ".html"

	fmt.Printf("Writing response body to %s\n", file)

	if splitUrl == "golang.org" {
		fmt.Println("Sleeping Golang Org")
		time.Sleep(time.Second * 10)
	}

	if err = ioutil.WriteFile(file, bodyBytes, 0664); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Done Writing File:", splitUrl+fmt.Sprintln(strings.Repeat("#", 20)))

	wg.Done()
}
