package secret

import (
	"fmt"
	"strconv"
)

func IntegerToBinary(n uint) string {
	binString := strconv.FormatInt(int64(n), 2)
	for len(binString) < 5 {
		binString = fmt.Sprintf("0%s", binString)
	}
	return binString
}

func Reverse(a[]string) []string {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func Handshake(code uint) []string {
	binString := IntegerToBinary(code)
	actions := make([]string, 0)
	reverse := false

	for i:=4; i>=0; i-- {
		if binString[i] == '1' {
			switch i {
			case 4:
				actions = append(actions, "wink")
			case 3:
				actions = append(actions, "double blink")
			case 2:
				actions = append(actions, "close your eyes" )
			case 1:
				actions = append(actions, "jump")
			case 0:
				reverse = true
			}
		}
		if reverse {
			actions = Reverse(actions)
		}
	}

	return actions
}
