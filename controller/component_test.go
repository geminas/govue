package controller

import (
	"testing"
)

func TestCreate(t *testing.T) {
	if err := CreateComponent("hello", "./"); err != nil {
		t.Error(err.Error())
	}
}
