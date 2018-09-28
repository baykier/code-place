// package temp

package temp

//摄氏度转换为华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//华氏温度转为摄氏温度

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}