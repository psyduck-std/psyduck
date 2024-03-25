package transform

import (
	"time"

	"github.com/psyduck-etl/sdk"
)

func Wait(parse sdk.Parser, _ sdk.SpecParser) (sdk.Transformer, error) {
	config := new(struct {
		Ms int64 `psy:"milliseconds"`
	})

	if err := parse(config); err != nil {
		return nil, err
	}

	t := time.Duration(config.Ms) * time.Millisecond
	return func(d []byte) ([]byte, error) {
		time.Sleep(t)
		return d, nil
	}, nil
}
