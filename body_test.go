package goerror_test

import (
	"github.com/prongbang/goerror"
	"testing"
)

type customErr struct {
	goerror.Body
}

func (u *customErr) Error() string {
	return u.Message
}

func newCustomErr() error {
	return &customErr{
		goerror.Body{
			Code:    "XXX",
			Message: "YYY",
			Data:    nil,
		},
	}
}

func BenchmarkGetBody(b *testing.B) {
	err := newCustomErr()

	for i := 0; i < b.N; i++ {
		_, _ = goerror.GetBody(err)
	}
}

func TestGetBody(t *testing.T) {
	err := newCustomErr()

	actual, _ := goerror.GetBody(err)

	if actual.Code != "XXX" || actual.Message != "YYY" {
		t.Error("Error:", actual)
	}
}
