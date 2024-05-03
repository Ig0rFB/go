package main

import (
	"image"
	_ "image/jpeg"
	"log"
	"os"
)
	
func main(){
	file, err:= os.Open("arriRef.jpg")
	if err != nil{
		log.Fatalln("File not found.")
	}
defer file.Close()

src, _, err := image.Decode(file)
if err != nil {
	log.Fatalln(err)
}

bounds := src.Bounds()
log.Print("Image Bounds: ", bounds)
}