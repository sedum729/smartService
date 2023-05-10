package feRobot

import "fmt"

type LanguageType int



const (
	Chinese LanguageType = iota + 1
)

func (r *Robot) Language() {
	fmt.Println("初始化语言系统", r.robotInfo.Language)
}

func (r *Robot) Send(words string) {
	fmt.Println("recived>>", words)
}
