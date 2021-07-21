package test

import (
	"quiz_master/utils"
	"testing"
)

// test function that get list of lof file in n minutes
func TestCreateQuestion(t *testing.T) {
	question := utils.QuestionListing{
		No:       1,
		Question: "How many letters are there in the English alphabet?",
		Answer:   "26",
	}

	err := utils.WriteQuestion(question)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Process Files")
	}
}

// test function that get a question
func TestGetQuestion(t *testing.T) {
	question_no := 1

	err := utils.GetQuestionByNo(question_no)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Process Files")
	}
}

// test get all question
func TestGetAllQuestions(t *testing.T) {
	res, _ := utils.GetQuestions()
	if res == nil {
		t.Error("Question list is empty")
	}
}

// test to get a correct answer
func TestAnswerQuestion(t *testing.T) {
	question_no := 1
	answer := "five"

	err := utils.GetAnswer(question_no, answer)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Process Files")
	}
}

// test to delete a question
func TestDeleteQuestion(t *testing.T) {
	question_no := 1

	err := utils.DeleteQuestion(question_no)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Process Files")
	}
}
