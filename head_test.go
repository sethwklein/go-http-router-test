package main

import (
	"testing"
)

// HEAD: servers should support HEAD for all GET requests, even if they go to
// the work of generating the content anyway, because they should not attempt
// to send content over the wire if the client has requested otherwise.

func TestHEAD(t *testing.T) {
}
