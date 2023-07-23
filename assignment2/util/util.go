package util

import "strconv"

func SeparateArrays(input []string) ([]float64, []string) {
	var floatArray []float64
	var stringArray []string

	for _, val := range input {
		// Attempt to convert the value to a float64
		if num, err := strconv.ParseFloat(val, 64); err == nil {
			floatArray = append(floatArray, num)
		} else {
			// If conversion fails, assume it's a string
			stringArray = append(stringArray, val)
		}
	}

	return floatArray, stringArray
}

func ConvertToOutputFormat(input []float64, input2 []string) []string {
	var output []string

	for _, val := range input {
		output = append(output, strconv.FormatFloat(val, 'f', -1, 64))
	}

	output = append(output, input2...)

	return output
}
