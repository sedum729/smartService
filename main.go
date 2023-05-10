package main

import (
	feRobot "fe-robot/robot"
	"fmt"
)

func main() {

	r := feRobot.Default()

	robotInfo := &feRobot.RobotInfo{
		Name:     "Jin1",
		Sex:      feRobot.Man,
		Emotion:  feRobot.Emotion2,
		Language: feRobot.Chinese,
	}

	ok := r.Create(robotInfo)

	if !ok {
		fmt.Println("create failed")
	}

	fmt.Println("完成机器人初始化>>>", *r)

	r.Send("你好")
}
