package numTOstringORstringTonum

import "strconv"

//字符串转int64

func StrToInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

//字符串转int32

func StrToInt32(s string) int32 {
	n, _ := strconv.ParseInt(s, 10, 32)
	return int32(n)
}

//字符串转int

func StrToInt(s string)int{
	n,_ := strconv.Atoi(s)
	return n
}

//字符串转float64

func StrToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return f
}