package main

import "testing"

func TestTravisHello(t *testing.T) {
	sayTravisHello()
	var hk NoteHook
	var pr PRHook
	err := parsePayload([]byte(jsonNote), &hk)
	if err != nil {
		t.Error(err)
	}
	err = parsePayload([]byte(jsonPr), &pr)
	if err != nil {
		t.Error(err)
	}
}
