package timeAndunix

import "time"


//时间戳转时间

func UnixToTime(m int64) string {
	formatTimeStr := time.Unix(m,0).Format("2006-01-02 15:04:05")
	return formatTimeStr
}
//时间戳转日期

func UnixToTimedate(m int64) string {
	formatTimeStr := time.Unix(m,0).Format("2006-01-02")
	return formatTimeStr
}