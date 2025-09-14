package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	result := 0
	for _, value := range nums {
		result = result ^ value
	}
	return result
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversHalf := 0
	for x > reversHalf {
		reversHalf = reversHalf*10 + x%10
		x /= 10
	}
	return x == reversHalf || x == reversHalf/10
}

func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if mapping[char] != top {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := ""
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 0; j < len(strs); j++ {
			if i > len(strs[j]) || strs[j][i] != char {
				return prefix
			}
		}
		prefix += string(char)
	}
	return prefix
}

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, ok := numMap[complement]; ok {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	fmt.Println("singleNumber()= ", singleNumber([]int{1, 1, 2, 2, 4}))
	fmt.Println("isPalindrome()= ", isPalindrome(1221))
	fmt.Println("isValid()= ", isValid("()[]{}"))
	fmt.Println("longestCommonPrefix()= ", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("plusOne()= ", plusOne([]int{1, 9, 9, 9}))
	fmt.Println("removeDuplicates()= ", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println("merge()= ", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println("twoSum()= ", twoSum([]int{2, 7, 11, 15}, 9))
}
