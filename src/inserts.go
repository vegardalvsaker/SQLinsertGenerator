package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	writeInsert()
}

func writeInsert() {
	file, err := os.Create("insert.sql")
	if err != nil {
		fmt.Errorf("some error %s", err)
	}
	insert := "insert into Notification values"
	file.Write([]byte(insert))
	for i := 0; i < 100; i++ {
		value := strings.Join([]string{"(default, 101, " , strconv.Itoa(i) ,", default, default),"}, "'")
		fmt.Println(value)
		file.Write([]byte(value))
	}
	file.Write([]byte(";"))
}