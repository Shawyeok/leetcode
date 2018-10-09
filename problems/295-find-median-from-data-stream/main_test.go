package main

import "testing"

func Test_FindMedian(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	mf.AddNum(4)
	mf.AddNum(10)
	median := mf.FindMedian()
	if median != 3.0 {
		t.Errorf("expected: %v, actual: %v", 3.0, median)
	}

	mf.AddNum(5)
	median = mf.FindMedian()
	if median != 5.0 {
		t.Errorf("expected: %v, actual: %v", 5.0, median)
	}
}
