package twelve

import "testing"

func TestInitialize(t *testing.T) {
        SetPrefix("twelve")
}

func TestString(t *testing.T) {
        exp := "test value one"
        v := String("string1")
        if *v != exp {
                t.Errorf("wanted %q got %q", exp, *v)
        }
}

func TestStringVar(t *testing.T) {
        t.Log("no test yet")
}

func TestBool(t *testing.T) {
        t.Log("no test yet")
}

func TestBoolVar(t *testing.T) {
        t.Log("no test yet")
}