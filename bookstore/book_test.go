package bookstore

import (
	"testing"
)

func TestBuy(t *testing.T) {
	t.Parallel()
	b := Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
	want := 1
	result, err := Buy(b)
	if err != nil {
		t.Fatal(err)
	}

	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}
