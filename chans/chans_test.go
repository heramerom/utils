package chans

import (
	"testing"
)

func TestSafeClose(t *testing.T) {
	ch := make(chan struct{})

	err := SafeClose(ch)
	if err != nil {
		t.Error(err.Error())
	}

	err = SafeClose(ch)
	if err == nil {
		t.Error(err.Error())
	}
}

func TestSafeSend(t *testing.T) {

	ch := make(chan struct{}, 5)

	err := SafeSend(ch, struct{}{})
	if err != nil {
		t.Error(err.Error())
	}

	err = SafeSend(ch, "string")

	if err == nil {
		t.Error(err.Error())
	}

	close(ch)

	err = SafeSend(ch, struct{}{})
	if err == nil {
		t.Error(err.Error())
	}

}
