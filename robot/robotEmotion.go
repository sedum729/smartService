package feRobot

import "fmt"

type EmotionType int

const (
	// 低落
	Emotion1 EmotionType = iota + 1
	// 一般
	Emotion2
	// 兴奋
	Emotion3
)

func (r *Robot) Emotion() {

	if r.robotInfo != nil {
		fmt.Println("初始化情绪管理", r.robotInfo.Emotion)
	}

}
