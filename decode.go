package main

import (
	"fmt"
	"os"
	"bufio"
	"image"

	"gopkg.in/auyer/steganography.v2"

)

func main() {
	inFile, _ := os.Open("./assets/test.jpg") // opening file
	inFile.Close()

	reader := bufio.NewReader(inFile) // buffer reader
	img, _, _ := image.Decode(reader) // decoding to golang's image.Image

	sizeOfMessage := steganography.GetMessageSizeFromImage(img) // retrieving message size to decode in the next line

	msg := steganography.Decode(sizeOfMessage, img) // decoding the message from the file
	fmt.Println(string(msg))
}
