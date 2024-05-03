package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("arriRed.jpg")

	if err != nil {
		log.Fatalln(err)
	}

}
