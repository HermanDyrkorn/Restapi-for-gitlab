package assignment2

import "time"

//variable that keeps track of time
var startTime time.Time

//Uptime function
func Uptime() float64 {
	//return the time since start as seconds
	return time.Since(startTime).Seconds()
}

//InitTime function
func InitTime() {
	//starts the timer
	startTime = time.Now()
}
