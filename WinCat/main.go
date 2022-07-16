package main

/*
*  15.07.2022
*  WinCat - utility for reading files in Windows CMD
*  By https://github.com/Avdushin
*  Copyright (c) 2022
 */

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile(os.Args[1])
	if len(os.Args) > 2 {
		log.Fatal("Please coose one file...")
	}
	if err != nil {
		log.Fatal("Type a file to read!")
	}
	fmt.Println(string(file))
}
