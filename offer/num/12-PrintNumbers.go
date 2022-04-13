package num

import "fmt"

var PrintNumbers = printNumbers

func printNumbers(n int) {
	if n <= 0 {
		return
	}
	ans := make([]byte, n)

	printNumbersRecursive(ans, len(ans), 0)
}

func printNumbersRecursive(nums []byte, l, i int) {
	if i == l {
		fmt.Println(string(nums))
		return
	}
	var j byte = 0
	for ; j < 10; j++ {
		nums[i] = '0' + j
		printNumbersRecursive(nums, l, i+1)
	}
}
