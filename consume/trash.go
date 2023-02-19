package consume

import (
	"github.com/gastrodon/psyduck/sdk"
)

func Trash(_ sdk.Parser, _ sdk.SpecParser) (sdk.Consumer, error) {
	return func() (chan []byte, chan error, chan bool) {
		data := make(chan []byte)
		done := make(chan bool)

		go func() {
			for range data {
			}

			done <- true
		}()

		return data, nil, done
	}, nil
}
