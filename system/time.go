package system

type Time struct {
	Microseconds int64
}

var TimeZero Time = Time{}

func (time Time) AsSeconds() float32 {
	return float32(time.Microseconds / 1000000)
}

func (time Time) AsMilliseconds() int32 {
	return int32(time.Microseconds / 1000)
}

func (time Time) AsMicroseconds() int64 {
	return time.Microseconds
}

func Seconds(amount float32) Time {
	time := Time{}
	time.Microseconds = int64(amount) * 1000000
	return time
}

func Milliseconds(amount int32) Time {
	time := Time{}
	time.Microseconds = int64(amount) * 1000
	return time
}

func Microseconds(amount int64) Time {
	time := Time{}
	time.Microseconds = amount
	return time
}