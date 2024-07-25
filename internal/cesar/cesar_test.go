package cesar

import "testing"

func TestRotateRight(t *testing.T) {
	cases := []struct {
		input    string
		key      int
		expected string
	}{{"test", 1, "uftu"}, {"aaaaa", 3, "ddddd"}, {"bUrp1", 20, "vOlj1"}}

	for _, test := range cases {
		t.Run(test.input, func(t *testing.T) {
			if got := Rotate(test.input, Right, test.key); got != test.expected {
				t.Errorf("expected %s, got: %s", test.expected, got)
			}
		})
	}

}

func TestRotateLeft(t *testing.T) {
	cases := []struct {
		input    string
		key      int
		expected string
	}{{"uftu", 1, "test"}, {"aaaaa", 3, "xxxxx"}, {"azoij21B", 20, "gfuop21H"}}

	for _, test := range cases {
		t.Run(test.input, func(t *testing.T) {
			if got := Rotate(test.input, Left, test.key); got != test.expected {
				t.Errorf("expected %s, got: %s", test.expected, got)
			}
		})
	}

}
