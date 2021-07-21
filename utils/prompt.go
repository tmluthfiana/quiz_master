package utils

import (
	"reflect"
)

// Prompt - Default prompt shell
type Prompt struct {
	super string `default:"quizmaster# "`
}

// SetPrompt - Set prompt
func SetPrompt(arg ...string) string {
	var prompt Prompt

	reftype := reflect.TypeOf(prompt)
	if len(arg) > 0 {
		prompt.super = arg[0] + "$ "
	} else {
		f, _ := reftype.FieldByName("super")
		prompt.super = f.Tag.Get("default")
	}

	return prompt.super
}
