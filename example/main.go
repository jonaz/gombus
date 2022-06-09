package main

import "fmt"

func main() {
	// gombus.SendUD2(nil)
	fmt.Printf("% x\n", toBCD(12345678))
	fmt.Printf("% x\n", FromUint(12345678, 4))
}

func toBCD(i uint64) []byte {
	var bcd []byte
	for i > 0 {
		low := i % 10
		i /= 10
		hi := i % 10
		i /= 10
		var x []byte
		x = append(x, byte((hi&0xf)<<4)|byte(low&0xf))
		// |= (0x0F & (address[i] - '0')) << (4 * k--);

		bcd = append(x, bcd...)
	}
	return bcd
}
func FromUint(value uint64, size int) []byte {
	buf := make([]byte, size)
	if value > 0 {
		remainder := value
		for pos := size - 1; pos >= 0 && remainder > 0; pos-- {
			tail := byte(remainder % 100)
			hi, lo := tail/10, tail%10
			buf[pos] = byte(hi<<4 + lo)
			remainder = remainder / 100
		}
	}
	return buf
}
