package util

import (
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}
// 封装了waitGroup库
func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)   // 计数器加一
	go func() {
		cb()
		w.Done()     // 计数器减一
	}()
}
