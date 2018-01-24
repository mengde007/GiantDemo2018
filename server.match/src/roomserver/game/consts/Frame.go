package consts

import "time"

const (
	FrameTimeMS       time.Duration = 25    //帧数帧时间，毫秒
	FrameTime         float64       = 0.025 //帧数帧时间
	FrameCountBy100MS uint64        = 4     //100毫秒刷新次数
	UpDownSpeed       float64       = 500.0 //方块上升下降的速度
)