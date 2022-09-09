package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gookit/color"
)

func main() {
	ls(nil)
}

func ls(err error) error {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			color.Green.Printf(color.Bold.Sprintf("%s\t", file.Name()))
		} else if !file.IsDir() {
			fmt.Printf("%s\t", file.Name())
		}
		if err != nil {
			log.Print(err)
		}
	}
	if err != nil {
		log.Print(err)
	}
	return nil
}
