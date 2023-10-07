package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	//input                  = "56.12 / 5212 / Go"
	format = "%f / %d / %s"

	inputReader *bufio.Reader
	input       string
	err         error
)

//func main() {
//	//fmt.Println("Please enter your full name: ")
//	//fmt.Scanln(&firstName, &lastName)
//	//fmt.Printf("Hi %s %s!\n", firstName, lastName)
//	//fmt.Sscanf(input, format, &f, &i, &s)
//	//fmt.Println("From the string we read: ", f, i, s)
//
//	inputReader = bufio.NewReader(os.Stdin)
//	fmt.Println("Please enter some input: ")
//	input, err = inputReader.ReadString('\t')
//	if err == nil {
//		fmt.Printf("The input was: %s\n", input)
//	}
//}

func main20() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program")
		return
	}

	fmt.Printf("Your name is %s", input)

	switch input {
	case "Philip\r\n":
		fmt.Println("qqq")
	case "Charis\r\n":
		fmt.Println("hhh")
	default:
		fmt.Println("555")
	}
}
