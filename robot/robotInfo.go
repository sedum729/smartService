package feRobot

type SexType int

const (
	// 男
	Man SexType = iota + 1
	// 女
	Woman
)

// 机器人属性
type RobotInfo struct {
	Name     string       // 名称
	Sex      SexType      // 性别
	Emotion  EmotionType  // 情绪
	Language LanguageType // 语言
}
