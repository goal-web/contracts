package contracts

import "time"

type Limiter interface {
	// Take 获取下一次通行时间
	// Get next pass time.
	Take() time.Time
}

type RateLimiter interface {
	// Limiter 根据给定的名称获取限流器
	// Get the current limiter by the given name.
	Limiter(name string, limiter func() Limiter) Limiter
}
