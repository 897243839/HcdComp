package HcdComp

import (
	cid "github.com/ipfs/go-cid"
	"time"
)

var (
	Block_hot                   = "blockhot.json"
	Compressflag                = "compressflag.json"
	Maphot                      = New[int]()
	MapLit                      = New[int]()
	Mode         CompressorType = 1
	Tsf                         = make(chan string, 128)
	Num                         = 1000000
	Hotk                        = make(chan string, 128)  //热数块add的key
	Clk                         = make(chan cid.Cid, 128) //冷数据cid
	Hck                         = make(chan cid.Cid, 128) //热数据cid
)

// var maps sync.RWMutex
// var myTimer = time.Now().Unix() // 启动定时器
var Ticker = time.NewTicker(60 * time.Second)     //计时器
var Ticker1 = time.NewTicker(30000 * time.Minute) //计时器
