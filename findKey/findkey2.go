package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type person2 struct {
	quId   int
	roleId string
	mapId  int
}

const (
	cmaps2    = "\"id\":"
	cmape2    = ",\"item_id\""
	croleIds2 = "\"#role_id\""
	croleIde2 = ",\"#time\""
)

var (
	per2s = make(map[string]*person2)
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	files, err := ioutil.ReadDir("E:/p37_fjweb_logs")
	handleErr(err)

	for _, file := range files {
		fmt.Println(file.Name())
		read, err := os.Open("E:/p37_fjweb_logs/" + file.Name())
		handleErr(err)

		defer read.Close()

		fileName := file.Name()
		quid, _ := strconv.Atoi(fileName[:strings.Index(fileName, "_")])

		reader := bufio.NewReader(read)
		for {
			a, _, c := reader.ReadLine()
			if c == io.EOF {
				break
			}
			test := string(a)
			//fmt.Println(test)

			maps := strings.Index(test, cmaps2) + 5
			mape := strings.Index(test, cmape2) + 0
			mapId, _ := strconv.Atoi(test[maps:mape])

			roleIds := strings.Index(test, croleIds2) + 11
			roleIde := strings.Index(test, croleIde2) + 0
			roleId := test[roleIds:roleIde]

			//fmt.Printf("quId: %d, count: %s, roleId: %s", quId, count, roleId)
			if per, ok := per2s[roleId]; ok {
				if per.mapId < mapId {
					per.mapId = mapId
				}
			} else {
				per2 := &person2{
					quid,
					roleId,
					mapId,
				}
				per2s[roleId] = per2
			}
		}
	}

	writer, err := os.OpenFile("findKey/xinyun.txt", os.O_WRONLY|os.O_CREATE, 0666)
	handleErr(err)
	defer writer.Close()

	w := bufio.NewWriter(writer)

	for _, per := range per2s {
		w.WriteString(fmt.Sprintf("区服: %d, 角色Id: %s ,地图Id: %d\n", per.quId, per.roleId, per.mapId))
		//fmt.Printf("区服: %d, 角色Id: %s ,超星杯数量: %d\n", per.quId, per.roleId, per.score)
	}
	w.Flush()
}
