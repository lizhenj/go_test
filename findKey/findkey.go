package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type person struct {
	quId   int
	roleId string
	score  int
}

const (
	cqus     = "log/server_"
	cque     = "_2023"
	ccount   = "\"count\""
	ccurrent = ",\"current\""
	croleId  = "\"#role_id\""
	ctimeD   = ",\"#time\""
)

var (
	pers = make(map[string]*person)
)

func main() {
	read, err := os.Open("./findKey/p37kadingche (1).txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(read)
	for {
		a, _, c := reader.ReadLine()
		if c == io.EOF {
			break
		}
		test := string(a)
		//fmt.Println(test)

		counts := strings.Index(test, ccount) + 8
		counte := strings.Index(test, ccurrent) + 0
		count, _ := strconv.Atoi(test[counts:counte])

		roleIds := strings.Index(test, croleId) + 11
		roleIde := strings.Index(test, ctimeD) + 0
		roleId := test[roleIds:roleIde]

		cpus := strings.Index(test, cqus) + 11
		cpue := strings.Index(test, cque) + 0
		quId, _ := strconv.Atoi(test[cpus:cpue])

		//fmt.Printf("quId: %d, count: %s, roleId: %s", quId, count, roleId)
		if per, ok := pers[roleId]; ok {
			if per.score < count {
				per.score = count
			}
		} else {
			per := &person{
				quId,
				roleId,
				count,
			}
			pers[roleId] = per
		}
	}

	for _, per := range pers {
		fmt.Printf("区服: %d, 角色Id: %s ,超星杯数量: %d\n", per.quId, per.roleId, per.score)
	}

}
