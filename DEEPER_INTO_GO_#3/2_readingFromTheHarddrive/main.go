package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("example.txt") // membaca isi dari file example.txt
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
