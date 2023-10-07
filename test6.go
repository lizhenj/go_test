package main

import "fmt"

func main6() {
	str := "Google"
	sl := []byte(str)
	var rev [100]byte
	j := 0
	for i := len(sl) - 1; i >= 0; i-- {
		rev[j] = sl[i]
		j++
	}
	str_rev := string(rev[:])
	fmt.Printf("The reversed string is -%s-\n", str_rev)
	str2 := "Google"
	sl2 := []byte(str2)
	for i, j := 0, len(sl2)-1; i < j; i, j = i+1, j-1 {
		sl2[i], sl2[j] = sl2[j], sl2[i]
	}
	fmt.Printf("The reversed string is -%s-\n", string(sl2))
	s := "My Test String"
	fmt.Println(s, " --> ", reverse(s))
}

func reverse(s string) string {
	runes := []rune(s)
	n, h := len(runes), len(runes)/2
	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
