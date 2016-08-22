package types

import "testing"

func TestMessageFromBytesLF(t *testing.T) {
	msg, err := MessageFromBytes([]byte("From: test@example.com\n\nbody"))
	if err != nil {
		t.Fatalf("parse message failed: %v", err)
	}
	if msg.Headers.Get("From") != "test@example.com" {
		t.Fatalf("header invalid")
	}
	if msg.Body != "body" {
		t.Fatalf("body invalid")
	}
}

func TestMessageFromBytesCRLF(t *testing.T) {
	msg, err := MessageFromBytes([]byte("From: test@example.com\r\n\r\nbody"))
	if err != nil {
		t.Fatalf("parse message failed: %v", err)
	}
	if msg.Headers.Get("From") != "test@example.com" {
		t.Fatalf("header invalid")
	}
	if msg.Body != "body" {
		t.Fatalf("body invalid")
	}
}
