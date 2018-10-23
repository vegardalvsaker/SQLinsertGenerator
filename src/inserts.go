package main

import (
	"os"
	"fmt"
	//"strings"
	//"strconv"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	listenForInput()
	//writeInsert()
}

func listenForInput() {
	printCommands()
	fmt.Println("What do you want to do?")

	reader := bufio.NewReader(os.Stdin)

	line, _, err := reader.ReadLine()
	errHandle(err)

	switch string(line) {
	case "quit":
		os.Exit(1)
	case "start":
		startGenerating(reader)
	case "help":
		printCommands()
	}

}
/*
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
}*/

func errHandle(err error) {
	if err != nil {
		fmt.Errorf("Some error %s", err)
	}
}

func printCommands() {
	fmt.Println("These are the commands: (Type help to show again)")
	fmt.Println("help")
	fmt.Println("start")
	fmt.Println("quit")
}

func startGenerating(reader *bufio.Reader) {
	fmt.Println("What table do you want to insert into?")
	tableName,_, err := reader.ReadLine()
	errHandle(err)



	fmt.Println("Which attributes?")
	fmt.Println("(enter in this format: att1, att2, att3")
	attributes,_, err2 := reader.ReadLine()
	attributes = append(attributes, ')')
	errHandle(err2)

	insert := strings.Join([]string{"insert into ", string(tableName), " (", string(attributes),  " values ("}, "")

	fmt.Println(insert)

	fmt.Println("How many entries?")
	numberOfEntries,_, err3 := reader.ReadLine()
	errHandle(err3)

	numberOfEntriesInt, err4 := strconv.Atoi(string(numberOfEntries))
	fmt.Println(numberOfEntriesInt)
	errHandle(err4)

	fmt.Println("")
	numberOfAttributes := 1
	for i := 0; i < len(attributes); i++ {
		if attributes[i] == ',' {
			numberOfAttributes++
		}
	}

	values := []string {}
	for j := 0; j < numberOfAttributes; j++ {
		fmt.Println("Value of the", j+1, "attribute?")
		value, _, err := reader.ReadLine()
		errHandle(err)
		value = append(value, ',')
		values = append(values, string(value))
	}
	fmt.Println("Before removing comma")
	fmt.Println(values)
	lastValue := values[len(values)-1]
	s := strings.Replace(lastValue, ",", ")", -1)
	values[len(values)-1] = s

	fmt.Println("After removing comma", values)
	//query := strings.Join(values, "")

	insert = strings.Join([]string{insert, values[0], values[1]}, "")

	fmt.Println(insert)

}