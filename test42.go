package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	keyValue = make(map[int]map[int]int)
)

func main42() {
	reader, err1 := os.Open("actorMsg.txt")
	if err1 != nil {
		panic(err1)
	}

	read := bufio.NewReader(reader)
	for {
		txt, err2 := read.ReadString('\n')
		if err2 == io.EOF {
			break
		}
		//fmt.Println(txt)
		sysIndex := strings.Index(txt, "sysId:")
		sysEnd := strings.Index(txt, "cmdId:")
		sys := txt[sysIndex+7 : sysEnd-2]
		cmd := txt[sysEnd+7 : len(txt)-1]
		sysId, err3 := strconv.Atoi(sys)
		if err3 != nil {
			panic(err3)
		}
		cmdId, err4 := strconv.Atoi(cmd)
		if err4 != nil {
			panic(err4)
		}
		if keyValue[sysId] == nil {
			keyValue[sysId] = map[int]int{}
		}
		keyValue[sysId][cmdId]++
	}

	final := []struct {
		sysId int
		cmdId int
		num   int
	}{}

	for sysId, next := range keyValue {
		for cmdId, num := range next {
			final = append(final, struct {
				sysId int
				cmdId int
				num   int
			}{sysId: sysId, cmdId: cmdId, num: num})
		}
	}

	sort.Slice(final, func(i, j int) bool {
		return final[i].num > final[j].num
	})

	for _, info := range final {
		fmt.Printf("sysId: %d, cmdId: %d, num: %d\n", info.sysId, info.cmdId, info.num)
	}
}
