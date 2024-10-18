package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	got := "yolo"
	want := "yolo"

	assert.Equal(t, want, got, "they should be equal")
}
