package main

import (
	"fmt"
	"os"
	"quiz_master/utils"
	"strconv"
)

func help() {
	fmt.Println("[answer <no> <answer>] - check a answer based on its question no")
}

func do(cmd []string) {
	if len(cmd) > 0 && cmd[0] == "-h" {
		help()
	}

	if len(cmd) < 2 {
		fmt.Println("Error - You need to specify question no and answer")
		help()
	} else {
		no, _ := strconv.Atoi(cmd[0])
		answer := cmd[1]
		fmt.Println(utils.GetAnswer(no, answer))
	}
}

func main() {
	args := os.Args[1:]
	do(args)
}
