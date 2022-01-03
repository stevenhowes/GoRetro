package GoRetro

import "time"

const (
	ScreenWidth  = 1024
	ScreenHeight = 768

	TargetTicksPerSecond   = 60
	DebugStatePrintSeconds = 1

	dataDir = "data/"
)

var Delta float64
var LastDebugStatePrint time.Time
var DebugTick bool
