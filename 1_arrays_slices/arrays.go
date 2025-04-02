package arraysslices

func Sum (nums [5]int) int{
	sum := 0
	for i := 0; i<5; i++{
		sum += nums[i];
	}
	return sum;
}

// range returns 2 values:- index, and value. we can replace _ with index/value if we dont need it.
func Sum2 (nums [5]int) int{
	sum := 0
	for _, value:= range nums{
		sum += value;
	}
	return sum;
}