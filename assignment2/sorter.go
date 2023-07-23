package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/datphamcode295/go-23/assignment2/sort"
	"github.com/datphamcode295/go-23/assignment2/util"
)

func main() {
	intFlag := flag.Bool("int", false, "Specify this flag for an integer array")
	stringFlag := flag.Bool("string", false, "Specify this flag for a string array")
	floatFlag := flag.Bool("float", false, "Specify this flag for a float64 array")
	mixFlag := flag.Bool("mix", false, "Specify this flag for a mixed array")

	flag.Parse()

	if *intFlag {
		args := flag.Args()
		intArray := make([]int, len(args))

		for i, arg := range args {
			num, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Error: Invalid integer input.")
				os.Exit(1)
			}
			intArray[i] = num
		}

		sort.SortArray(intArray)
		fmt.Println("Output:", intArray)
	} else if *stringFlag {
		args := flag.Args()

		sort.SortArray(args)
		fmt.Println("Output:", args)
	} else if *floatFlag {
		args := flag.Args()
		floatArray := make([]float64, len(args))

		for i, arg := range args {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Println("Error: Invalid float64 input.")
				os.Exit(1)
			}
			floatArray[i] = num
		}

		sort.SortArray(floatArray)
		fmt.Println("Output:", floatArray)
	} else if *mixFlag {
		args := flag.Args()

		numberArray, stringArray := util.SeparateArrays(args)

		sort.SortArray(numberArray)
		sort.SortArray(stringArray)
		fmt.Println("Output:", util.ConvertToOutputFormat(numberArray, stringArray))
	}
}
