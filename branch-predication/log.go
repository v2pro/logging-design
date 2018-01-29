package test

import "fmt"

const LevelTrace = 10
const LevelDebug = 20

var MinLevel = LevelDebug

func Trace(event string) {
	if LevelTrace < MinLevel {
		return
	}
	fmt.Println(event)
}