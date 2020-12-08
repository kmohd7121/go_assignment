package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("sample.text")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hi, my name is array and this file was created using go")
	file.Close()
	stream,err:=ioutil.ReadFile("sample.text")
	if err!=nil{
		log.Fatal(err)
	}
	s1:=string(stream)
	fmt.Println(s1)
}