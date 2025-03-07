package utils

import (
	"encoding/json"
	"fmt"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func PrettyPrint(body []byte) []byte {
	var prettyJSON map[string]interface{}
	err := json.Unmarshal(body, &prettyJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	prettyBody, err := json.MarshalIndent(prettyJSON, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil
	}
	return prettyBody
}
