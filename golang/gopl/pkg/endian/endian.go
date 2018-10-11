//判断 机器的大端序还是小端序
package endian

var MSB , LSB bool

func init() {
	int_16 := int16(0x1234)
	int_8 := int8(int_16)

	if 0x34 == int_8 {
		LSB =  true
	} else {
		MSB = true
	}
}
// 是否大端序
func IsMSB() bool {
	return true == MSB
}
func IsLSB() bool {
	return true == LSB
}