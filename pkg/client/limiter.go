package client

import (
	"context"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type RateLimiter struct {
	*rate.Limiter
	StatsLimiter *rate.Limiter
	*ThrottleSeconds
}

func newRateLimiter() *RateLimiter {
	return &RateLimiter{
		Limiter:         rate.NewLimiter(20, 20),
		StatsLimiter:    rate.NewLimiter(60, 60),
		ThrottleSeconds: &ThrottleSeconds{},
	}
}

func (l *RateLimiter) Wait(statRequest IsStatRequest, ctx context.Context) error {
	time.Sleep(l.ThrottleSeconds.Duration())

	if statRequest {
		err := l.StatsLimiter.Wait(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	err := l.Limiter.Wait(ctx)
	if err != nil {
		return err
	}
	return nil
}

type ThrottleSeconds struct {
	int32
	sync.RWMutex
}

func (t *ThrottleSeconds) Set(i int32) {
	defer t.Unlock()
	t.Lock()
	t.int32 = i
}

func (t *ThrottleSeconds) SetInt(i int) {
	defer t.Unlock()
	t.Lock()
	t.int32 = int32(i)
}

func (t *ThrottleSeconds) Duration() time.Duration {
	return time.Duration(t.int32) * time.Second
}

func (t *ThrottleSeconds) Int32() int32 {
	return t.int32
}

func (t *ThrottleSeconds) Int() int {
	return int(t.int32)
}
