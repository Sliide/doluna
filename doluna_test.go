package doluna_test

import "testing"

func AssertEqual(t *testing.T, x interface{}, z interface{}) {
	if x != z {
		t.Errorf("%+v is not equal to %+v", x, z)
	}
}

func AssertNotEqual(t *testing.T, x interface{}, z interface{}) {
	if x == z {
		t.Errorf("%+v is equal to %+v", x, z)
	}
}
