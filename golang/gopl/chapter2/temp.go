//temp 
package temp

type Celsius float64 //摄氏温度

type Fahrenheit float64 //华氏温度

const (
	AbsoluteZeroC  	Celsius = -273.15
	FreeXingC 		Celsius = 0
	BolingC 		Celsius = 100
)

//摄氏度转换为华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//华氏温度转为摄氏温度

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}