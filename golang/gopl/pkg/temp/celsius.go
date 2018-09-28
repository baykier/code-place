//temp 
package temp

import (
	"fmt"
)

type Celsius float64 //摄氏温度

//大写开头的函数可以被外部调用
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
//摄氏度转换为华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
//摄氏度转开尔文
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
//测试能否被外部调用
func testCall() {
	fmt.Println("call success!")
}