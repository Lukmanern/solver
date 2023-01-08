package main

func main() {

}

func twoSum(nums []int, target int) []int {
	var i, j, v1, v2 int
	var isBreak bool = false

	if len(nums) == 2 {
		return []int{0, 1}
	}
	for i, v1 = range nums {
		for j, v2 = range nums {
			if j <= i {
				continue
			}
			if v1+v2 == target {
				isBreak = true
				break
			}
		}
		if isBreak {
			break
		}
	}

	return []int{i, j}
}