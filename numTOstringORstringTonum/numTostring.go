package numTOstringORstringTonum

import "strconv"


//int64转字符串

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

//float64转字符串

func Float64ToString(flo float64) string {
	return strconv.FormatFloat(flo, 'f', -1, 64)
}

//int转字符串

func IntToString(i int) string{
	return strconv.Itoa(i)
}

//unit32转string

func Unit32ToString(i uint32) string {
	s := strconv.FormatUint(uint64(i),10)
	return s
}

