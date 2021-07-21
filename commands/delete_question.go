package main

import (
	"fmt"
	"os"
	"quiz_master/utils"
	"strconv"
)

func help() {
	fmt.Println("[delete_question] - Delete a question based on its id")
}

func do(cmd []string) {
	if len(cmd) > 0 && cmd[0] == "-h" {
		help()
	}

	if len(cmd) < 1 {
		fmt.Println("Error - You need to specify id")
		help()
	} else {
		no, _ := strconv.Atoi(cmd[0])
		fmt.Println(utils.DeleteQuestion(no))
	}
}

func main() {
	args := os.Args[1:]
	do(args)
}
