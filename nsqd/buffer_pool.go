package nsqd

import (
	"bytes"
	"sync"
)

var bp sync.Pool

// 使用sync.Pool建立bytes.buffer的pool，减少gc的压力
func init() {
	bp.New = func() interface{} {
		return &bytes.Buffer{}
	}
}

func bufferPoolGet() *bytes.Buffer {
	return bp.Get().(*bytes.Buffer)
}

func bufferPoolPut(b *bytes.Buffer) {
	bp.Put(b)
}
