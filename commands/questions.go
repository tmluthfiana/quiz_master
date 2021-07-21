package main

import (
	"fmt"
	"os"
	"quiz_master/utils"
)

func help() {
	fmt.Println("[questions] - Get all questions")
}

func do(cmd []string) {
	if len(cmd) > 0 && cmd[0] == "-h" {
		help()
	}

	_, err := utils.GetQuestions()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	args := os.Args[1:]
	do(args)
}
