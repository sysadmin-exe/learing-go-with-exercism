package sorting

import "fmt"
import "strconv"
// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

type Number struct {
    value int
}

func (n Number) number() int {
    return n.value
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float32(nb.Number()))
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
	if fn, ok := fnb.(FancyNumber); ok {
		num, err := strconv.Atoi(fn.Value())
		if err != nil {
			return 0
		}
		return num
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	if fn, ok := fnb.(FancyNumber); ok {
		num, err := strconv.Atoi(fn.Value())
		if err != nil {
			return fmt.Sprintf("This is a fancy box containing the number 0.0")
		}
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float32(num))
	}
    return fmt.Sprintf("This is a fancy box containing the number 0.0")
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
        case int: 
        	return DescribeNumber(float64(v))
        case float64: 
        	return DescribeNumber(v)
        case NumberBox:
        	return DescribeNumberBox(v)
        case FancyNumberBox:
        	return DescribeFancyNumberBox(v)
        default:
        	return fmt.Sprintf("Return to sender")
    }
}
