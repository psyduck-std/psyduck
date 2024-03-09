package produce

import (
	"github.com/psyduck-std/sdk"
)

type constant struct {
	Value     string `psy:"value"`
	StopAfter int    `psy:"stop-after"`
}

func Constant(parse sdk.Parser, specParse sdk.SpecParser) (sdk.Producer, error) {
	config := new(constant)
	parse(config)

	count := 0
	next := func() ([]byte, bool, error) {
		count++
		return []byte(config.Value), config.StopAfter == 0 || count < config.StopAfter, nil
	}

	return func() (chan []byte, chan error) {
		data := make(chan []byte)
		go func() {
			sdk.ProduceChunk(next, specParse, data, nil)
			close(data)
		}()
		return data, nil
	}, nil
}
