package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main22() {
	//fileinput()
	//read_write_file1()
	//gzipped()
	fileoutput()
}

func fileoutput() {
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("open error")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}

func gzipped() {
	fName := "MyFile.tar.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Println("open err")
		os.Exit(1)
	}
	defer fi.Close()
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}

func read_write_file1() {
	inputFile := "test.txt"
	outputFile := "test1.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("read err")
		return
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func fileinput() {
	inputFile, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open txt err")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		text, err := inputReader.ReadString('\n')
		fmt.Printf("the info was: %v", text)
		if err == io.EOF {
			return
		}
	}
}
