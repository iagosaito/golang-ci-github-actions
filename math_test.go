package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Invalid result. Result %d, Expected %d", total, 30)
	}
}
