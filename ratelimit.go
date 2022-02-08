package contracts

import "time"

type Limiter interface {
	Take() time.Time
}

type RateLimiter interface {
	Limiter(name string, limiter func() Limiter) Limiter
}
