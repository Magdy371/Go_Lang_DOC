package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var fileName string = "./newFile.txt"
	file, err := os.Create(fileName)
	checkError(err)
	defer file.Close()
	length, err := io.WriteString(file, "My name is Magdy")
	checkError(err)
	fmt.Printf("# of characters in the file is: %v\n", length)
	readFile(fileName)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println(string(data))
}
