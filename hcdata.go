package HcdComp

import (
	"time"
)

var (
	block_hot                   = "blockhot.json"
	compressflag                = "compressflag.json"
	maphot                      = New[int]()
	mapLit                      = New[int]()
	Mode         CompressorType = 1
)

// var maps sync.RWMutex
// var myTimer = time.Now().Unix() // 启动定时器
var ticker = time.NewTicker(60 * time.Second)     //计时器
var ticker1 = time.NewTicker(30000 * time.Minute) //计时器

var cb = func(exists bool, valueInMap int, newValue int) int {
	if !exists {
		return newValue
	}
	if valueInMap > 999 {
		return 1000
	}
	valueInMap += newValue
	return valueInMap

}
