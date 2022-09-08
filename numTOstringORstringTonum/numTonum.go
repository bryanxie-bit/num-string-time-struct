package numTOstringORstringTonum

import (
	"fmt"
	"math"
	"strconv"
)

//保留两位有效数字

func FloatTwoFloat(f float64) float64 {
	f1 := math.Trunc(f*1e2+0.5) * 1e-2
	f1Str := strconv.FormatFloat(f1, 'f', 2, 64)
	value, _ := strconv.ParseFloat(f1Str, 64)
	return value
}
//保留五位有效数字

func FloatFiveFloat(f float64) float64 {
	f1 := math.Trunc(f*1e2+0.5) * 1e-2
	f1Str := strconv.FormatFloat(f1, 'f', 5, 64)
	value, _ := strconv.ParseFloat(f1Str, 64)
	return value
}

//取小数点后两位

func FloatTwoPointFloat(f float64) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64)
	return value
}

//小数点后保留五位

func FloatFivePointFloat(f float64)float64{
	value := StrToFloat64(fmt.Sprintf("%.5f", f))
	return value
}

//四舍五入

func FloatToFloat(f float64)float64{
	return math.Floor(f+0.5)
}

//Byte转int64

func ByteToInt64(b []byte) int64 {
	var i int64
	for _, v := range b {
		if v > 47 && v < 58 {
			i = int64(v) - 48 + i*10
		} else {
			return 0
		}
	}
	return i
}

//Byte转int32

func ByteToInt32(b []byte) int32 {
	var i int32
	for _, v := range b {
		if v > 47 && v < 58 {
			i = int32(v) - 48 + i*10
		} else {
			return 0
		}
	}
	return i

}
