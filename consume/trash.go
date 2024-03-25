package consume

import (
	"github.com/psyduck-std/sdk"
)

func Trash(sdk.Parser, sdk.SpecParser) (sdk.Consumer, error) {
	return func(recv <-chan []byte, errs chan<- error, done chan<- struct{}) {
		for range recv {
		}
		close(done)
	}, nil
}
