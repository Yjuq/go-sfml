package system

// Time -> Represents a time value.
type Time struct {
	Microseconds int64
}

// TimeZero -> Predefined "zero" time value
var TimeZero Time = Time{}

// AsSeconds -> Return a time value as a number of seconds
func (time Time) AsSeconds() float32 {
	return float32(time.Microseconds) / 1000000
}

// AsMilliseconds -> Return a time value as a number of milliseconds
func (time Time) AsMilliseconds() int32 {
	return int32(time.Microseconds) / 1000
}

// AsMicroseconds -> Return a time value as a number of microseconds
func (time Time) AsMicroseconds() int64 {
	return time.Microseconds
}

// Seconds -> Construct a time value from a number of seconds
func Seconds(amount float32) Time {
	time := Time{}
	time.Microseconds = int64(amount) * 1000000
	return time
}

// Milliseconds -> Construct a time value from a number of milliseconds
func Milliseconds(amount int32) Time {
	time := Time{}
	time.Microseconds = int64(amount) * 1000
	return time
}

// Microseconds -> Construct a time value from a number of microseconds
func Microseconds(amount int64) Time {
	time := Time{}
	time.Microseconds = amount
	return time
}