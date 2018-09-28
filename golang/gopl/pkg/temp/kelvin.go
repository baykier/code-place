//package temp

package temp

import (
	"fmt"
)

type Kelvin float64

//开尔文转为摄氏温度
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

//开尔文转为华氏温度

func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k * 9 / 5 - 459.67)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}