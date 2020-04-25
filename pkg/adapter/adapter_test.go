package adapter

import "testing"

func TestFacebookAdapter_Send(t *testing.T) {
	msg := "Hello, world."
	want := FbNote + msg
	if got := (FacebookAdapter{}).Send(msg); got != want {
		t.Errorf("FacebookAdapter.Send() = %q, want %q", got, want)
	}
}

func TestBold_Format(t *testing.T) {
	msg := "Hello, world."
	want := TwNote + msg
	if got := (TwitterAdapter{}).Send(msg); got != want {
		t.Errorf("TwitterAdapter.Send() = %q, want %q", got, want)
	}
}
