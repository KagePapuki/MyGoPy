package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	//"sync"
	"github.com/KagePapuki/MyGoPy/core"
)

var inputs []string
//Global_map := make(map[string]string)
//var wg sync.WaitGroup
var in_block bool = false
var block_code []string

func process_single(command string) {
	fmt.Printf("Processed command: '%s'\n",print_token_slice(core.LexAndYacc(command)))
}

func process_block() {
	fmt.Printf("Processed block: '%s'\n",print_token_slice(core.LexAndYacc(strings.Join(block_code, "\n"))))
}

func print_token_slice(slice []*core.Token) string {
	var result string = "["
	for _, token := range slice {
		result += (token.Ttype() + ": \"" + token.Tcontent() + "\", ")
	}
	result += "]"
	return result
}

func preprocess(command string) {
	inputs = append(inputs, command)
	if in_block {
		if command == "" {
			in_block = false
			process_block()
			block_code = make([]string,10)
		} else {
			block_code = append(block_code, command)
		}
	} else {
		if strings.HasSuffix(command, ":") {
			in_block = true
			block_code = append(block_code, command)
		} else {
			process_single(command)
		}
	}
}

func handle_input() {
	reader := bufio.NewReader(os.Stdin)

	for true {
		if in_block {
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