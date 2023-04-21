package EvenOdd

import "fmt"

func main() {
	//initialize slice of integers from 0 till 10
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//iterate through int slice and check if it is odd or even
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
