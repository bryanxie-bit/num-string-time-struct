package timeAndunix

import "time"

//时间转时间戳

func TimeToUnix(s string) int64 {
	loc, _ := time.LoadLocation("Local")
	times, _ := time.ParseInLocation("2006-01-02",s, loc)
	timeunix := times.Unix()
	return timeunix
}

//时间转时间戳

func TimeToUnixHourMin(s string) (int64){
	loc, _ := time.LoadLocation("Local")
	times, _ := time.ParseInLocation("2006-01-02 15:04:05",s, loc)
	timeunix := times.Unix()
	return timeunix
}