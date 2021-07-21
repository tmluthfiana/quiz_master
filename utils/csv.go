package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/divan/num2words"
)

const (
	csvQuestionPath = "/tmp/question.csv"
)

var (
	ErrQAE = errors.New("Error - Question already exist")
	ErrQLE = errors.New("Error - Question list is empty")
)

// QuestionListing - Structure used to organize the item
type QuestionListing struct {
	No       int
	Question string
	Answer   string
}

func trimQuotes(word string) string {
	word = strings.Replace(word, "'", "", -1)
	word = strings.Replace(word, `"`, "", -1)

	return word
}

func WriteQuestion(Question QuestionListing) error {
	// Check if question already exist
	if DoesQuestionExist(Question) {
		return ErrQAE
	}

	file, err := os.OpenFile(csvQuestionPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	item := []string{fmt.Sprintf("%d|%s|%s",
		Question.No,
		trimQuotes(Question.Question),
		trimQuotes(Question.Answer))}

	res := wr.Write(item)
	if res == nil {
		fmt.Println("Question ", Question.No, "created:")
		fmt.Println("Q: ", Question.Question)
		fmt.Println("A: ", Question.Answer)
	}

	return res
}

func GetQuestions() ([]QuestionListing, error) {
	entries := ReadQuestions()
	if len(entries) == 0 {
		return nil, ErrQLE
	}
	questionList := []QuestionListing{}
	fmt.Fprintln(os.Stdout, "No | Question | Answer")
	for index, _ := range entries {
		a := strings.Split(entries[index][0], "|")
		fmt.Println(a[0], " ", a[1], " ", a[2])
		no, _ := strconv.Atoi(a[0])
		question := QuestionListing{
			No:       no,
			Question: a[1],
			Answer:   a[2],
		}
		questionList = append(questionList, question)
	}

	return questionList, nil
}

func ReadQuestions() [][]string {
	file, err := os.OpenFile(csvQuestionPath, os.O_RDONLY, 0644)
	if err != nil {
		return nil
	}
	r := csv.NewReader(file)
	r.LazyQuotes = true
	lines, err := r.ReadAll()
	if err != nil {
		return nil
	}

	if len(lines) == 0 {
		return nil
	}

	return lines
}

func DoesQuestionExist(Question QuestionListing) bool {
	entries := ReadQuestions()
	if len(entries) == 0 {
		return false
	}

	for _, entry := range entries {
		splEntry := strings.Split(entry[0], "|")
		no, _ := strconv.Atoi(splEntry[0])
		question := splEntry[1]
		answer := splEntry[2]
		if Question.No == no &&
			trimQuotes(Question.Question) == trimQuotes(question) &&
			trimQuotes(Question.Answer) == trimQuotes(answer) {

			return true
		}
	}

	return false
}

// DeleteQuestion - Remove an question from csv question file
func DeleteQuestion(no int) (err error) {
	entries := ReadQuestions()
	if len(entries) == 0 {
		return ErrQLE
	}

	for index, entry := range entries {
		splEntry := strings.Split(entry[0], "|")
		_no, _ := strconv.Atoi(splEntry[0])
		if no == _no {
			entries[index][0] = ""
			err = nil
			break
		} else {
			err = errors.New("Error - Question does not exist")
		}
	}

	if err != nil {
		return err
	}

	err = ReGenerateCSVQuestion(entries)
	if err != nil {
		return err
	}

	return err
}

// ReGenerateQuestion - Regenerates the csv question file
//                        for delete/update operations
func ReGenerateCSVQuestion(lines [][]string) error {
	file, err := os.Create(csvQuestionPath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	for _, line := range lines {
		if len(line[0]) > 0 {
			err = w.Write(line)
			if err != nil {
				return err
			}
		}
	}
	w.Flush()
	return nil
}

func GetQuestionByNo(no int) error {
	entries := ReadQuestions()
	if len(entries) == 0 {
		return ErrQLE
	}

	for index, entry := range entries {
		splEntry := strings.Split(entry[0], "|")
		_no, _ := strconv.Atoi(splEntry[0])
		if no == _no {
			a := strings.Split(entries[index][0], "|")
			fmt.Println("Q: ", a[1])
			fmt.Println("A: ", a[2])
			break
		} else {
			return errors.New("Error - Question not found")
		}
	}

	return nil
}

func GetAnswer(no int, answer string) (err error) {
	entries := ReadQuestions()
	if len(entries) == 0 {
		return ErrQLE
	}

	for index, entry := range entries {
		splEntry := strings.Split(entry[0], "|")
		_no, _ := strconv.Atoi(splEntry[0])
		if no == _no {
			a := strings.Split(entries[index][0], "|")
			correctAnswer := a[2]
			if answerNum, err := strconv.Atoi(a[2]); err == nil {
				correctAnswer = num2words.Convert(answerNum)
			}
			if a[2] == answer || correctAnswer == answer {
				fmt.Println("Correct!")
			} else {
				fmt.Println("Not Correct!")
			}
			err = nil
			break
		} else {
			err = errors.New("Error - Question not found")
		}
	}

	if err != nil {
		return err
	}

	return err
}
