package main

import (
	"image"
	"image/color"
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

img, format, err := image.Decode(file)
if err != nil {
	log.Fatalln("Image decode error, check file format")
}

averageColour := averagePixelValue(img)
log.Printf("Average PV: R: %d, G: %d, B: %d, A: %d\n",
averageColour.R, averageColour.G, averageColour.B, averageColour.A)
 log.Printf("Decoded image format: %s", format)
}

func averagePixelValue(img image.Image) color.RGBA {
	bounds := img.Bounds()
	log.Printf("Image Bounds: %s", bounds)
	var rTotal, gTotal, bTotal, aTotal uint64
	var count uint64

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y). RGBA()

			rTotal += uint64(r >> 8)
			gTotal += uint64(g >> 8)
			bTotal += uint64(b >> 8)
			aTotal += uint64(a >> 8)
			count++
		}
	}
	return color.RGBA{
		R: uint8(rTotal / count),
		G: uint8(gTotal / count),
		B: uint8(bTotal / count),
		A: uint8(aTotal / count),
	}
}
