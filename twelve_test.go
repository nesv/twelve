package twelve

import "testing"

func TestInitialize(t *testing.T) {
	SetPrefix("twelve")
}

func TestString(t *testing.T) {
	tests := map[string]string{"string1": "test value one"}
	for envar, exp := range tests {
		if result, err := String(envar); result != exp {
			t.Errorf("wanted %q got %q", exp, result)
		} else if err != nil && err == ErrNotSet {
			t.Error(err)
		}
	}
}

func TestBool(t *testing.T) {
	tests := map[string]bool{
		"bool_true":  true,
		"bool_false": false,
		"bool_fail":  false}

	for envar, exp := range tests {
		if result, err := Bool(envar); result != exp {
			t.Errorf("wanted %v got %v", exp, result)
		} else if err != nil && err == ErrNotSet {
			t.Error(err)
		}

	}
}
