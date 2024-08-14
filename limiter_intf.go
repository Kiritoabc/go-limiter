package go_limiter

import "errors"

// 计数器算法
// 滑动窗口算法
// 漏桶算法
// 令牌桶算法
// 单机限流
// 分布式限流

var (
	// ErrExceededLimit 请求太多
	ErrExceededLimit = errors.New("too many req")
)

// RateLimiter 限流器
type RateLimiter interface {
	Tack() error
}
