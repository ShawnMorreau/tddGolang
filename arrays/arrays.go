package arrays

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(arrs ...[]int) []int {
	var resArr []int
	for _, arr := range arrs {
		resArr = append(resArr, Sum(arr))
	}
	return resArr
}
