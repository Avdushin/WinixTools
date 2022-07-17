package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gookit/color"
)

func main() {
	ll(nil)
}

func ll(err error) error {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("%s\t", file.Mode())
			color.Green.Printf(color.Bold.Sprintf("%s\n", file.Name()))
		} else if !file.IsDir() {
			fmt.Printf("%s\t%s\n", file.Mode(), file.Name())
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
