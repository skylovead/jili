// Package clock 可以模拟 time 和 context 标准库的部分行为。
//
// All methods are safe for concurrent use.
package clock

import (
	"context"
	"time"
)

// Clock represents an interface to the functions in the standard time and context packages.
type Clock interface {
	After(d time.Duration) <-chan time.Time
	AfterFunc(d time.Duration, f func()) *Timer
	NewTicker(d time.Duration) *Ticker
	NewTimer(d time.Duration) *Timer
	Now() time.Time
	Since(t time.Time) time.Duration
	Sleep(d time.Duration)
	Tick(d time.Duration) <-chan time.Time
	Until(t time.Time) time.Duration

	// DeadlineContext returns a copy of the parent context with the associated
	// Clock deadline adjusted to be no later than d.
	DeadlineContext(parent context.Context, d time.Time) (context.Context, context.CancelFunc)

	// ContextWithTimeout returns DeadlineContext(parent, Now(parent).Add(timeout)).
	ContextWithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
}

// Updatable 驱动 mock clock 的时间流逝。
// Add 与 Set 在设置新的当前时间 t 之前，
// 会触发所有 deadline <= t 的 tick 和 timer
type Updatable interface {
	// Add 在 mock clock 的当前时间加上 d
	// 成为 mock clock 新的当前时间
	Add(d time.Duration) time.Time
	// Set 把 mock clock 的当前时间设置为 t
	// 并返回与上一个当前时间的差值
	// 注意，如果 t 早于 mock clock 的当前时间，会 Panic
	Set(t time.Time) time.Duration
	// Move 会触发离 mock clock 当前时间最近的那个 timer 或 tick
	// 并把 mock clock 的当前时间设置成触发时间，
	// 返回值是新的当前时间和与前一个当前时间的差值。
	Move() (time.Time, time.Duration)
}
