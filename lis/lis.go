package lis

//find longest increasing subsequence
//result O(n log n) with binary search

func Lis(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	longest, middle := 0, 0
	prev := make([]int, len(input))
	mem := make([]int, len(input)+1)
	for i := 0; i < len(input); i++ {
		left := 1
		right := longest
		for left <= right {
			middle = (left + right) / 2
			if input[mem[middle]] < input[i] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
		prev[i] = mem[left-1]
		mem[left] = i
		if left > longest {
			longest = left
		}
	}
	ret := make([]int, longest)
	j := mem[longest]
	for i := longest - 1; i >= 0; i-- {
		ret[i] = input[j]
		j = prev[j]
	}
	return ret
}
