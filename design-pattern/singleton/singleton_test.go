package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	// インスタンス取得
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error(
			"expected pointer to Singleton after calling GetInstance(). not nil")
	}

	expectedCounter := counter1
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf(
			"After calling for the first time to count, the count must be 1 but it is %d\n",
			currentCount,
		)
	}

	// インスタンス取得
	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error(
			"Expected same instance in counter2 but it got a different instance")
	}

	// インスタンス取得
	counter3 := GetInstance()
	currentCount = counter3.RemoveOne()
	if currentCount != 0 {
		t.Errorf(
			"After calling for the third time to count, the count must be 0 but it is %d\n",
			currentCount,
		)
	}
}
