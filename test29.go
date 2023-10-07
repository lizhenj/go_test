package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

//func main() {
//	var network bytes.Buffer
//	enc := gob.NewEncoder(&network)
//	dec := gob.NewDecoder(&network)
//
//	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
//	if err != nil {
//		log.Fatal("encode error:", err)
//	}
//
//	var q Q
//	err = dec.Decode(&q)
//	if err != nil {
//		log.Fatal("decode error:", err)
//	}
//	fmt.Printf("%q: {%d,%d}\n", q.Name, q.X, q.Y)
//}

func main29() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
	decode()
}
func decode() {
	file, _ := os.Open("vcard.gob")
	defer file.Close()
	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)
	var kk VCard
	err = dec.Decode(&kk)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(kk)
}
