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
var in_separator_block bool = false
var in_add_separator_block bool = false
var block_code []string

func process_single(command string) string {
	tokens, err := core.LexAndYacc(command)
	if err == "" {
		fmt.Printf("Processed command: '%s'\n",print_token_slice(tokens))
	}
	return err
}

func process_block() string {
	tokens, err := core.LexAndYacc(strings.Join(block_code, "\n"))
	if err == "" {
		fmt.Printf("Processed block: '%s'\n",print_token_slice(tokens))
	}
	return err
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
		if command == "" && in_separator_block == false {
			err := process_block()
			switch err {
			case "":
				in_block = false
				block_code = make([]string,10)
			case "INCOMPLETE BLOCK":
				in_separator_block = true
				in_add_separator_block = true
				block_code = append(block_code, command)
			default:
				in_block = false
				block_code = make([]string,10)
				fmt.Println("[Error]" + err)
			}
		} else if in_separator_block {
			block_code = append(block_code, command)
			err := process_block()
			switch err {
			case "":
				if in_add_separator_block {
					in_add_separator_block = false
					in_separator_block = false
				} else {
					block_code = make([]string,10)
					in_separator_block = false
					in_block = false
				}
			case "INCOMPLETE BLOCK":
				//block_code = append(block_code, command)
			default:
				in_add_separator_block = false
				in_separator_block = false
				block_code = make([]string,10)
				fmt.Println("[Error]" + err)
			}
		} else {
			block_code = append(block_code, command)
		}
	} else {
		if strings.HasSuffix(command, ":") || strings.HasSuffix(command, "\\") {
			in_block = true
			block_code = []string{}
			block_code = append(block_code, command)
		} else {
			err := process_single(command)
			switch err {
			case "":
				//no error
			case "INCOMPLETE BLOCK":
				in_block = true
				in_separator_block = true
				block_code = []string{}
				block_code = append(block_code, command)
			default:
				fmt.Println("[Error]" + err)
			}
		}
	}
}

func handle_input() {
	reader := bufio.NewReader(os.Stdin)

	for true {
		if in_block {
			fmt.Printf("... ")
		} else {
			fmt.Printf(">>> ")
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