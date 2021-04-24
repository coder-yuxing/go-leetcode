package two_sum

// TwoSum given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
//
// 输入一个int类型数组和一个int类型的目标值，返回数组中相加之和等于目标值的两个元素的数组下标
//
// 你可以假定，每个数组中仅存在一种满足条件的结果，并且你不能使用同一元素两次
//
// 你可以以任意顺序返回结果
func TwoSum(nums []int, target int) []int {
	var result []int

	var temp = make(map[int]int)

	for i, num := range nums {
		if j, ok := temp[num]; ok {
			if target == num * 2 {
				result = append(result, j)
				result = append(result, i)
				return result
			}
		} else {
			temp[num] = i
		}
	}

	for n := range temp {
		otherNum := target - n
		if i, ok := temp[otherNum]; ok {
			j := temp[n]
			if i == j {
				continue
			}
			result = append(result, i)
			result = append(result, j)
			break
		}
	}
	return result
}
