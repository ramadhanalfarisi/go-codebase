package helpers

import "encoding/json"

func CollectGraphqlArguments[T any](input map[string]any, model T) error {
	resJson, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resJson, &model)
	return err
}
