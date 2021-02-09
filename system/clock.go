package system

import "golang.org/x/sys/windows"

var clockCreate *windows.LazyProc
var clockCopy *windows.LazyProc
var clockDestroy *windows.LazyProc
var clockGetElapsedTime *windows.LazyProc
var clockRestart *windows.LazyProc

func clock() {
	clockCreate = system.NewProc("sfClock_create")
	clockCopy = system.NewProc("sfClock_copy")
	clockDestroy = system.NewProc("sfClock_destroy")
	clockGetElapsedTime = system.NewProc("sfClock_getElapsedTime")
	clockRestart = system.NewProc("sfClock_restart")
}

// Clock -> Utility class that measures the elapsed time.
type Clock struct {
	ptr uintptr
}

// ClockCreate -> Create a new clock and start it
func ClockCreate() Clock {
	ret, _, _ := clockCreate.Call()
	return Clock{ptr: ret}
}

// Copy -> Create a new clock by copying an existing one
func (clock Clock) Copy() Clock {
	ret, _, _ := clockCopy.Call(clock.ptr)
	return Clock{ptr: ret}
}

// Destroy -> Destroy a clock
func (clock Clock) Destroy() {
	clockDestroy.Call(clock.ptr)
}

// GetElapsedTime -> Get the time elapsed in a clock
func (clock Clock) GetElapsedTime() Time {
	ret, _, _ := clockGetElapsedTime.Call(clock.ptr)
	return Microseconds(int64(ret))
}

// Restart -> Restart a clock
func (clock Clock) Restart() Time {
	ret, _, _ := clockRestart.Call(clock.ptr)
	return Microseconds(int64(ret))
}