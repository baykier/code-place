// package temp

package temp

import (
	"fmt"
)

type Fahrenheit float64 //华氏温度


func (c Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", c)
}
