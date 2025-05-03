package helper

import "strings"

type ShiftDirection int

const (
	Left  ShiftDirection = -1
	Right ShiftDirection = 1
)

func Shifter(c byte, direction ShiftDirection) byte {
	keyboardRows := []string{
		"QWERTYUIOP",
		"ASDFGHJKL",
		"ZXCVBNM",
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	}

	for _, row := range keyboardRows {
		idx := strings.IndexByte(row, c)
		if idx != -1 {
			length := len(row)

			// Circular indexing or circular shifting
			newIdx := (idx + int(direction) + length) % length
			return row[newIdx]
		}
	}

	return c
}
