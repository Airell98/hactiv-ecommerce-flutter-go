package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func CheckAndSaveBodyChan(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		c <- err.Error()
		return
	}

	fmt.Println("Done sending Http Request to:", url)

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	_ = bodyBytes

	if err != nil {
		fmt.Println("error occured", err)
	}

	splitUrl := strings.Split(url, "//")[1]

	file := "coba/" + splitUrl + ".html"

	fmt.Printf("Start writing file: %s\n", splitUrl)

	// if splitUrl == "golang.org" {
	// 	fmt.Println("Sleeping Golang Org")
	// 	time.Sleep(time.Second * 10)
	// }

	if err = ioutil.WriteFile(file, bodyBytes, 0664); err != nil {
		log.Fatalln(err)
	}

	c <- fmt.Sprintf("Done Writing File: %v", splitUrl)

}
