package arraysslices

import "testing"



func TestArrays(t *testing.T){
	t.Run("normal level testing", func (t *testing.T){
		nums:= [5]int{1, 2, 3, 4, 5};
		got:= Sum(nums);
		want:= 15;
	
		checkIfCorrect(t, got, want, nums);
	})

	t.Run("better data structures: ", func(t *testing.T) {
		nums:= [5]int{1, 2, 3, 4, 5};
		got:= Sum2(nums);
		want:= 15;	
	
		checkIfCorrect(t, got, want, nums)
	})

}

func checkIfCorrect(t *testing.T, got, want int, nums [5]int){
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d, arrya is %v", got, want, nums);
	}
}