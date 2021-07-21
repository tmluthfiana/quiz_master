package main

import (
	"fmt"
	"os"
	"quiz_master/utils"
	"strconv"
)

func help() {
	fmt.Println("[question no] - Get a question based on its question no")
}

func do(cmd []string) {
	if len(cmd) > 0 && cmd[0] == "-h" {
		help()
	}

	if len(cmd) < 1 {
		fmt.Println("Error - You need to specify question number")
		help()
	} else {
		no, _ := strconv.Atoi(cmd[0])
		err := utils.GetQuestionByNo(no)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	args := os.Args[1:]
	do(args)
}
