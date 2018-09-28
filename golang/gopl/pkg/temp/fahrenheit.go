// package temp

package temp

import (
	"fmt"
)

type Fahrenheit float64 //华氏温度

//华氏温度转为摄氏温度

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//华氏温度转为开尔文

func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f + 459.67) * 5 / 9)
}

func (c Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", c)
}
