/*
    场景设计题,比如设计限流器等
*/

// Leetcode第359题:设计日志速率限制器
// 参考:https://blog.csdn.net/weixin_41593360/article/details/135205678

class Logger {
public:
    std::unordered_map<std::string, int> x;
    Logger() {

    }

    bool shouldPrintMessage(int timestamp, std::string message) {
        if(!x.count(message)) x.insert({message, timestamp});
        else {
            if(timestamp < x[message] + 10) return false;
            else {
                x[message] = timestamp;
                return true;
            }
        }
        return true;
    }

    // 优化:优化成下一次可以出现的时间而非上次出现的时间,这样做可以把两者合并
    bool shouldPrintMessage(int timestamp, std::string message) {
        bool result = timestamp >= x[message];
        if(result) x[message] = timestamp + 10;
        return result;
    }
};


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
