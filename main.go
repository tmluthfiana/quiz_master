package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"quiz_master/utils"
	"regexp"
	"strings"
)

const (
	quizmaster = "quizmaster# "
	banner     = "\t\t Welcome to Quiz Master!\n"
)

var (
	prompt = utils.SetPrompt()
)

func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(strings.ToLower(commandStr), "\n")
	commandStr = strings.TrimSuffix(commandStr, "\n")
	rgx := regexp.MustCompile("'.*?'|\".*?\"|\\S+")
	argCommandStr := rgx.FindAllString(commandStr, -1)

	if len(argCommandStr) <= 0 {
		return nil
	}

	switch argCommandStr[0] {
	case "help":
		fmt.Fprintln(os.Stdout, "Command | Description")
		fmt.Fprintln(os.Stdout, "create_question <no> <question> <answer> | Creates a question")
		fmt.Fprintln(os.Stdout, "update_question <no> <question> <answer> | Updates a question")
		fmt.Fprintln(os.Stdout, "delete_question <no> | Deletes a question")
		fmt.Fprintln(os.Stdout, "question <no> | Shows a question")
		fmt.Fprintln(os.Stdout, "questions | Shows question liss")
	case "exit":
		if prompt != quizmaster {
			prompt = quizmaster
			return nil
		}
		os.Exit(0)
	default:
		if len(argCommandStr) > 0 {
			cmd := []string{argCommandStr[0]}
			fcmd, err := utils.FindCmd(cmd)
			if err != nil {
				return err
			}

			runCmd := exec.Command(fcmd, argCommandStr[1:]...)
			runCmd.Stderr = os.Stderr
			runCmd.Stdout = os.Stdout

			err = runCmd.Run()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// print banner
	bannerl := "\t\t " + strings.Repeat("-", 20)
	fmt.Fprintln(os.Stdout, bannerl)
	fmt.Fprintln(os.Stdout, banner)

	for {
		fmt.Print(prompt)

		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stdout, err)
		}
	}
}
