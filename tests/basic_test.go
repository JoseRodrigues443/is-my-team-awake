package tests

import "testing"

// "TestAdd is a function that takes a testing.T object and compares the result of 4 + 6 to 10."
// Its here has a test that will never fail normally, good to find config or system fail (for example using then flag.Parse in a wrong way)
//
// The first line of the function is a call to the testing.T object's Errorf method. This method takes
// a format string and a list of arguments. The format string is a string that contains text
// interspersed with placeholders for the arguments. The arguments are substituted into the format
// string in place of the placeholders
func TestAdd(t *testing.T) {

	got := 4 + 6
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
