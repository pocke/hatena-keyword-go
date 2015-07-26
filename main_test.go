package keyword

import "testing"

func TestSimilarWord(t *testing.T) {
	l, err := SimilarWord([]string{"Hatena", "Perl"})
	if err != nil {
		t.Fatal(err)
	}
	if len(l) == 0 {
		t.Fatal("Response is empty")
	}
}
