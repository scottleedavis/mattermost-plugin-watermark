package main

import (
	"bufio"
	"fmt"
	"gopkg.in/auyer/steganography.v2"
	"image"
	"os"
)

//Using the "image/jpeg" module to decode img, err := jpeg.Decode(reader)
func main() {
	inFile, err := os.Open(os.Args[1]) // opening file
	if err != nil {
		fmt.Println("unable to open file")
	}
	defer inFile.Close()

	var img image.Image
	reader := bufio.NewReader(inFile)  // buffer reader
	img, _, err = image.Decode(reader) // decoding to golang's image.Image
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err != nil {
		fmt.Println("unable to decode file")
	}
	sizeOfMessage := steganography.GetMessageSizeFromImage(img) // retrieving message size to decode in the next line

	msg := steganography.Decode(sizeOfMessage, img) // decoding the message from the file
	fmt.Println(string(msg))
}
