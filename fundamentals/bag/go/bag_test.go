package bag_test

var bagIsEmptyTests = []struct {
	bag      Bag  // input
	expected bool // expected result
}{
	{Bag{}, true},
}
