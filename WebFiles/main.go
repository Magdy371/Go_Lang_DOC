package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://api.github.com/users/github"

func main() {
	var client http.Client = http.Client{}
	var file *os.File
	var err error
	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	checkError(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	var res *http.Response
	res, err = client.Do(req)
	checkError(err)
	defer res.Body.Close()
	fmt.Printf("response Type: %T \n", res)
	var bytes []byte
	bytes, err = io.ReadAll(res.Body)
	checkError(err)
	var content string = string(bytes)
	fmt.Println(content)
	var fileName string = "./content.txt"
	file, err = os.Create(fileName)
	checkError(err)
	defer file.Close()
	fmt.Printf("file name: %v\n created", fileName)
	_, err = io.WriteString(file, content)
	checkError(err)
	fileContent, err := os.ReadFile(fileName)
	fmt.Printf("file content is \n: %v\n", string(fileContent))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
