package main

import "testing"

func TestSum(t *testing.T) {
	want := 5
	got := sum(-5, 10) 
	if want != got {t.Fatalf("Error: got : %v, want %v", got, want)}
	want = 6
	got = sum(-5, 10) 
	if want != got {t.Fatalf("Error: got : %v, want %v", got, want)}

}
