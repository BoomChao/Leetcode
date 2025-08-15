package design

import (
	"sync"
	"time"
)

// 设计一个限流器,限制每秒的请求速率
type RateLimiter struct {
	rate   int         // 每秒允许的请求树
	ticker time.Ticker // 定时器
	tokens int         // 当前令牌数
	mu     sync.Mutex  // 同步访问的互斥锁
}

func NewRateLimiter(rate int) *RateLimiter {
	rl := &RateLimiter{
		rate:   rate,
		tokens: rate,
		ticker: *time.NewTicker(time.Second),
	}
	go func() {
		for range rl.ticker.C {
			rl.mu.Lock()
			rl.tokens = rate
			rl.mu.Unlock()
		}
	}()
	return rl
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

func main() {
	limiter := NewRateLimiter(100)
	for i := 0; i < 100; i++ {
		if limiter.Allow() {
			// do something
		}
	}
}
