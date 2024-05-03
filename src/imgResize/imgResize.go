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

decodedImg, _, err := image.Decode(file)
if err != nil {
	log.Fatalln("Image decode error, check file format")
}

bounds := decodedImg.Bounds()
log.Print("Image Bounds: ", bounds)
}