GO ?= go
SRC := delete_question.go answer_question.go \
	questions.go question.go create_question.go

all: build

build:
	@for file in $(SRC); do \
		go build $$file ;\
	done \

.PHONY: all build
