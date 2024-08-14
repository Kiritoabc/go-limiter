package fixed_window

import (
	"fmt"
	"testing"
	"time"
)

// TestFixedWindowLimiter 测试固定窗口大小限流
func TestFixedWindowLimiter(t *testing.T) {
	limiter := NewFixedWindowLimiter(1*time.Second, 10)
	takes := 0
	for i := 0; i < 100; i++ {
		if err := limiter.Tack(); err != nil {
			continue
		}
		takes++
		t.Logf("takes: %d", takes)
	}
}

func TestPrint(t *testing.T) {
	fmt.Println("hello world")
}
