package main

import (
	"bufio"
	"fmt"
	"image"
	"os"

	"gopkg.in/auyer/steganography.v2"
)

func main() {
	inFile, err := os.Open(os.Args[1]) // opening file
	if err != nil {
		fmt.Println("unable to open file")
	}
	defer inFile.Close()

	img, _, err := image.Decode(bufio.NewReader(inFile)) // decoding to golang's image.Image
	if err != nil {
		fmt.Println("Unable to decode file: " + err.Error())
		return
	}

	sizeOfMessage := steganography.GetMessageSizeFromImage(img) // retrieving message size to decode in the next line

	msg := steganography.Decode(sizeOfMessage, img) // decoding the message from the file
	fmt.Println(string(msg))
}
