package main

import (
	"fmt"
	"os"
	"quiz_master/utils"
	"strconv"
)

func do(cmd []string) {
	var ask utils.QuestionListing

	ask.No, _ = strconv.Atoi(cmd[0])
	ask.Question = cmd[1]
	ask.Answer = cmd[2]

	err := utils.WriteQuestion(ask)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	args := os.Args[1:]
	do(args)
}
