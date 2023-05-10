package feRobot

import "sync"

var lock sync.Mutex

type Robot struct {
	robotInfo *RobotInfo
}

func startEngine(r *Robot) {

	// needStartQueue := make([]func(), 10)

	// append(needStartQueue, r.Emotion)
	// append(needStartQueue, r.Language)
	// append(needStartQueue, r.Logic)

	var needStartQueue = []func()(){
		r.Emotion,
		r.Language,
		r.Logic,
	}

	for _, execer := range needStartQueue {
		execer()
	}
}

// 创建机器人
func (r *Robot) Create(info *RobotInfo) (creatResult bool) {
	lock.Lock()

	if r.robotInfo == nil {
		r.robotInfo = info
	}

	defer lock.Unlock()

	startEngine(r)

	return r.robotInfo != nil
}

// 创建机器人实例
func Default() (r *Robot) {
	rInstance := &Robot{}

	return rInstance
}
