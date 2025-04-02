package slicessumtest

func Sum(nums []int)int{
	sum := 0
	for _, value := range nums {
		sum += value;
	}
	return sum;
}

func SumAll(nums ...[]int)[]int{
	lengthOfNums := len(nums);

	sum := make([]int, lengthOfNums)
	for i, values := range nums{
		sum[i] = Sum(values)
	}
	return sum
}


func SumAllUsingAppend(nums ...[]int)[]int{

	var sum []int

	for _, values := range nums{
		sum = append(sum, Sum(values))
	}
	return sum
}