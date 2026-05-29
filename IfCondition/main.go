package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)
import "bufio"

func main() {
	var outPut string
	fmt.Println("Enter your age:")
	reader := bufio.NewReader(os.Stdin)
	stm, _ := reader.ReadString('\n')
	age, err := strconv.ParseUint(strings.TrimSpace(stm), 10, 64)
	if err != nil {
		fmt.Println(err)
	} else if age <= 15 {
		outPut = "you are not an adult"
	} else {
		outPut = "u r an adult"
	}
	fmt.Println(outPut)
}
