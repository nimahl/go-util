package hs106

import "testing"

func TestNewHS106(t *testing.T) {
	_, err := NewHS106(record)
	if err != nil {
		t.Fatal(err)
	}
}
