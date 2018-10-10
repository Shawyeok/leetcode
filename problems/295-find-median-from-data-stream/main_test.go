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
	if median != 4.0 {
		t.Errorf("expected: %v, actual: %v", 4.0, median)
	}

	mf.AddNum(0)
	mf.AddNum(100)
	median = mf.FindMedian()
	if median != 4.0 {
		t.Errorf("expected: %v, actual: %v", 4.0, median)
	}
}

func Test_FindMedian2(t *testing.T) {
	mf := Constructor()
	mf.AddNum(0)
	mf.AddNum(0)
	median := mf.FindMedian()
	if median != 0.0 {
		t.Errorf("expected: %v, actual: %v", 0.0, median)
	}
}

func Test_FindMedian3(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(0)
	mf.AddNum(0)
	median := mf.FindMedian()
	if median != 0.0 {
		t.Errorf("expected: %v, actual: %v", 0.0, median)
	}
}

func Test_FindMedian4(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(0)
	mf.AddNum(0)
	mf.AddNum(2)
	mf.AddNum(2)
	median := mf.FindMedian()
	if median != 1.0 {
		t.Errorf("expected: %v, actual: %v", 1.0, median)
	}
}
