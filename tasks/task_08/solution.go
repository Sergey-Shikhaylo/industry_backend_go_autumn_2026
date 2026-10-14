package main

import (
	"sync"
	"time"
)

type Clock interface{ Now() time.Time }
type Limiter struct {
	mu     sync.Mutex
	clock  Clock
	rate   float64
	burst  int
	tokens float64
	last   time.Time
}

func NewLimiter(clock Clock, ratePerSec float64, burst int) *Limiter {
	if burst <= 0 || clock == nil {
		return &Limiter{}
	}

	l := &Limiter{
		clock:  clock,
		burst:  burst,
		tokens: float64(burst),
		rate:   ratePerSec,
		last:   clock.Now(),
	}

	go func() {
		l.mu.Lock()
		defer l.mu.Unlock()
		for {
			if l.rate > 0 {
				if l.clock.Now().Sub(l.last).Seconds() >= 1.0 {
					if float64(burst)-l.tokens >= l.rate {
						l.tokens += l.rate
					}
					l.last = l.clock.Now()
				}
			}
		}
	}()

	return l
}

func (l *Limiter) AllowN(n int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.burst <= 0 || l.clock == nil {
		return false
	}

	if float64(n) >= l.tokens || n <= 0 {
		return false
	}

	if l.rate > 0 {
		if l.clock.Now().Sub(l.last).Seconds() >= 1.0 {
			if float64(l.burst)-l.tokens >= l.rate {
				l.tokens += l.rate
			}
			l.last = l.clock.Now()
		}
	}

	l.tokens -= float64(n)
	return true
}
