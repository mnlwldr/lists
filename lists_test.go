package lists

import (
	"testing"
)

var (
	text    = []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
)

func TestHeadString(t *testing.T) {

	want := []string{"The"}
	got := Head(text)

	if len(want) != len(got) {
		t.Fatalf("Want %v, got %v", want, got)
	}

	for index, value := range got {
		if want[index] != value {
			t.Fatalf("Want %v, got %v", want, got)
		}
	}
}

func TestTailString(t *testing.T) {
	want := []string{"quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	got := Tail(text)

	if len(want) != len(got) {
		t.Fatalf("Want %v, got %v", want, got)
	}

	for index, value := range got {
		if want[index] != value {
			t.Fatalf("Want %v, got %v", want, got)
		}
	}
}

func TestHeadInt(t *testing.T) {

	want := []int{1}
	got := Head(numbers)

	if len(want) != len(got) {
		t.Fatalf("Want %v, got %v", want, got)
	}

	for index, value := range got {
		if want[index] != value {
			t.Fatalf("Want %v, got %v", want, got)
		}
	}
}

func TestTailInt(t *testing.T) {
	want := []int{2, 3, 4, 5, 6, 7, 8, 9, 0}
	got := Tail(numbers)

	if len(want) != len(got) {
		t.Fatalf("Want %v, got %v", want, got)
	}

	for index, value := range got {
		if want[index] != value {
			t.Fatalf("Want %v, got %v", want, got)
		}
	}
}
