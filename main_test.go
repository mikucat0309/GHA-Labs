package main

import "testing"

func TestHello(t *testing.T) {
	if hello("") != "Hello, !" {
		t.Error("testcase 1 fail")
	} else {
		t.Log("testcase 1 pass")
	}

	if hello("World") != "Hello, World!" {
		t.Error("testcase 2 fail")
	} else {
		t.Log("testcase 2 pass")
	}
}
