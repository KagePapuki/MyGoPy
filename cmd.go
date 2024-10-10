package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	//"sync"
)

var Inputs []string
//Global_map := make(map[string]string)
//var wg sync.WaitGroup
var In_block bool = false
var Block_code []string

func process_single(command string) {
	fmt.Printf("Processed command: '%s'\n",command)
}

func process_block() {
	fmt.Printf("Processed block: '%v'\n",Block_code)
}

func preprocess(command string) {
	Inputs = append(Inputs, command)
	if In_block {
		if command == "" {
			In_block = false
			process_block()
			Block_code = make([]string,10)
		} else {
			Block_code = append(Block_code, command)
		}
	} else {
		if strings.HasSuffix(command, ":") {
			In_block = true
			Block_code = append(Block_code, command)
		} else {
			process_single(command)
		}
	}
}

func handle_input() {
	reader := bufio.NewReader(os.Stdin)

	for true {
		if In_block {
			fmt.Printf("...")
		} else {
			fmt.Printf(">>>")
		}
		command, err := reader.ReadString('\n')
		if err == nil {
			command = strings.TrimSuffix(command, "\r\n")
			command = strings.TrimSuffix(command, "\n")
			command = strings.TrimSuffix(command, "\r")
			if command == "exit()" {
				return
			} else {
				preprocess(command)
			}
		} else {
			switch err.Error() {
				case "EOF": return
				default: log.Fatal(err)
			}
		}
	}
}

func main() {
	fmt.Println("MyGoPy 0.1.0\na Python interpreter based on Golang\nGithub: https://github.com/KagePapuki/MyGoPy")
	handle_input()
}