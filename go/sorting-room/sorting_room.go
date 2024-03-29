package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)

}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if _, ok := fnb.(FancyNumber); !ok {
		return 0
	}
	num, ok := strconv.Atoi(fnb.Value())
	if ok != nil {
		num = 0
	}
	return num
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	num := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %d.0", num)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	response := "Return to sender"
	switch num := i.(type) {
		case int:
			response = DescribeNumber(float64(num)) 
		case float64:
			response = DescribeNumber(num)
		case NumberBox:
			response = DescribeNumberBox(num)
		case FancyNumberBox:
			response = DescribeFancyNumberBox(num)
		}
	return response
}
