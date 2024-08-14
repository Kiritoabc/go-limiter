package fixed_window

import (
	go_limiter "github.com/Kiritoabc/go-limiter"
	"sync"
	"sync/atomic"
	"time"
)

// 固定窗口计数器限流

var (
	once sync.Once
)
var _ go_limiter.RateLimiter = &FixedWindowLimiter{}

// NewFixedWindowLimiter 创建固定窗口计数器限流
func NewFixedWindowLimiter(snippet time.Duration, allowReq int32) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		snippet:  snippet,
		allowReq: allowReq,
	}
}

// FixedWindowLimiter 固定窗口计数器限流
type FixedWindowLimiter struct {
	snippet  time.Duration // 窗口大小
	curReq   int32         // 当前请求数
	allowReq int32         // 允许请求数
}

// Tack 获取请求（判断是否能获取请求）
func (l *FixedWindowLimiter) Tack() error {
	// 启动定时器
	once.Do(func() {
		go func() {
			for {
				select {
				case <-time.After(l.snippet):
					atomic.StoreInt32(&l.curReq, 0)
				}
			}
		}()
	})

	curReq := atomic.LoadInt32(&l.curReq)
	if curReq >= l.allowReq {
		return go_limiter.ErrExceededLimit
	}
	if !atomic.CompareAndSwapInt32(&l.curReq, curReq, curReq+1) {
		return go_limiter.ErrExceededLimit
	}
	return nil
}
