package ifunny

import (
	"encoding/json"
	"fmt"

	"github.com/gastrodon/psyduck/sdk"
)

func getContentAuthor(parse func(interface{}) error) (sdk.Transformer, error) {
	return func(data []byte) ([]byte, error) {
		content := new(Content)
		if err := json.Unmarshal(data, content); err != nil {
			panic(fmt.Errorf("can't unmarshal bytes %v as Content: %s", data, err))
		}

		creatorBytes, err := json.Marshal(content.Creator)
		if err != nil {
			panic(err)
		}

		return creatorBytes, nil
	}, nil
}

func getItemID(parse func(interface{}) error) (sdk.Transformer, error) {
	return func(data []byte) ([]byte, error) {
		identity := new(Identity)
		if err := json.Unmarshal(data, identity); err != nil {
			panic(fmt.Errorf("can't unmarshal bytes %v as Identity: %s", data, err))
		}

		return []byte(identity.ID), nil
	}, nil
}
