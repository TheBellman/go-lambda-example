package main

import "testing"

func TestHandleRequest(t *testing.T) {
	want := "Hello Fred, from code build!"
	event := MyEvent{Name: "Fred"}

	// our request handler ignores the context and is hardwired to not return an error
	if got, _ := HandleRequest(nil, event); got != want {
		t.Errorf("HandleRequest() = %q, want %q", got, want)
	}
}
