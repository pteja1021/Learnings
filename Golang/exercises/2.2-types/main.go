package main
import "fmt"
import "strconv"

func stringToInt(str string) int {
	// convert to int
	result,_ := strconv.Atoi(str)
	return result
}

func stringToFloat(str string) float64 {
	// convert to float with 4 digits of precision
	result, _ := (strconv.ParseFloat(str,64))
	s := fmt.Sprintf("%.4f",result)
	result, _ = strconv.ParseFloat(s,64)
	return result
}

func FloatToString(value float64) string {
	// convert to float with 2 digits of precision
	result := strconv.FormatFloat(value,'g',2,64)
	return result
}

func main() {
	isFailed := false
	if stringToInt("10") != 10 {
		fmt.Println("Failed: stringToInt")
		isFailed = true
	}

	if stringToFloat("123.33333333333") != 123.3333 {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}

	if FloatToString(1.0/3) != "0.33" {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}

	if !isFailed {
		fmt.Println("All tests passed")
	}
}