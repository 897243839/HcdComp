package HcdComp

import (
	"time"
)

var (
	Block_hot                   = "blockhot.json"
	Compressflag                = "compressflag.json"
	Maphot                      = New[int]()
	MapLit                      = New[int]()
	Mode         CompressorType = 1
)

// var maps sync.RWMutex
// var myTimer = time.Now().Unix() // 启动定时器
var Ticker = time.NewTicker(60 * time.Second)     //计时器
var Ticker1 = time.NewTicker(30000 * time.Minute) //计时器

var Cb = func(exists bool, valueInMap int, newValue int) int {
	if !exists {
		return newValue
	}
	if valueInMap > 999 {
		return 1000
	}
	valueInMap += newValue
	return valueInMap

}
