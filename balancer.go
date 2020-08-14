package balancer

import (
	"strconv"
	"strings"
)

// Solution is (as the name imply) the solution to Jum'at Hek Challenge
// it takes a string of weights (type int) separated by a space
// outputs the minimum difference of weight if distributed into 2 baskets
func Solution(input string) (int, error) {
	//convert string into int slice
	intStrings := strings.Split(input, " ")
	nums := []int{}
	for _, intString := range intStrings {
		num, err := strconv.Atoi(intString)
		//return -1, and error if any
		if err != nil {
			return -1, err
		}
		nums = append(nums, num)
	}
	//calculate the diff and return it
	return Balancer(nums), nil
}

//Balancer calculates minimum difference of sum if a slice of int is distributed into 2 slices
func Balancer(items []int) int {
	//sort it first
	sort(items)
	//calculate the half of the initial sum to be the pivot point
	sum := 0
	for _, item := range items {
		sum += item
	}
	var half float64 = float64(sum) / 2

	//determine the nearest to half, be it lower than half and higher than half
	var i, lowerThanHalfCandidate, higherThanHalfCandidate int
	for i = len(items) - 1; i > -1; i-- {
		newWeight := float64(lowerThanHalfCandidate + items[i])
		if newWeight < half {
			lowerThanHalfCandidate = int(newWeight)
			continue
		}
		if newWeight == half {
			return 0
		}
		if newWeight > half {
			higherThanHalfCandidate = int(newWeight)
		}
	}

	return higherThanHalfCandidate - lowerThanHalfCandidate
}

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		k := i
		for k != 0 && arr[k] < arr[k-1] {
			temp := arr[k-1]
			arr[k-1] = arr[k]
			arr[k] = temp
			k--
		}
	}
}
