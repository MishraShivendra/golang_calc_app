/* Program: Sample code for Calculator
 * Usage: ./program number_a number_b operator
 * Supported operators: *, /, +, -,%;
 */

/*Import necessary modules and packages*/
package main
import (
"fmt"
"os"
"strconv"
)

/*Function to read arguments*/
func read_args() []string {
	return os.Args[1:]
}

/*Function to perform arthmetic operation*/
func adder(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) float32 {
	return float32 (a / b)
}

func mod(a, b int) int {
	return a % b
}

/*Function checking operator and showing result*/
func calc(arguments []string) {

	var (
		number_1, number_2 int
		err error
		error_msg="Error in atoi"
	    )

	switch arguments[2] {
	case "+":
		number_1, err = strconv.Atoi(arguments[0])
		if err != nil {
			fmt.Println(error_msg)
		}
		number_2, err = strconv.Atoi(arguments[1])
		if err != nil {
			fmt.Println(error_msg)
		}

		fmt.Println("Ans:", adder(number_1, number_2))
	case "-":
		number_1, err = strconv.Atoi(arguments[0])
		if err != nil {
			fmt.Println(error_msg)
		}
		number_2, err = strconv.Atoi(arguments[1])
		if err != nil {
			fmt.Println(error_msg)
		}

		fmt.Println("Ans:", sub(number_1, number_2))
	case "*":
		number_1, err = strconv.Atoi(arguments[0])
		if err != nil {
			fmt.Println(error_msg)
		}
		number_2, err = strconv.Atoi(arguments[1])
		if err != nil {
			fmt.Println(error_msg)
		}

		fmt.Println("Ans:", mul(number_1, number_2))
	case "/":
		number_1, err = strconv.Atoi(arguments[0])
		if err != nil {
			fmt.Println(error_msg)
		}
		number_2, err = strconv.Atoi(arguments[1])
		if err != nil {
			fmt.Println(error_msg)
		}

		fmt.Println("Ans:", div(number_1, number_2))
	case "%":
		number_1, err = strconv.Atoi(arguments[0])
		if err != nil {
			fmt.Println(error_msg)
		}
		number_2, err = strconv.Atoi(arguments[1])
		if err != nil {
			fmt.Println(error_msg)
		}

		fmt.Println("Ans:", mod(number_1, number_2))
	
	default:
		fmt.Println("Wrong option")
	}
}


/*Main Function*/
func main() {
	
	var main_args []string
    	main_args = read_args()
	calc(main_args)
}
