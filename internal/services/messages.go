package services

import (
	"encoding/json"
	"fmt"
	"os"
)

type MessagesList []string

type MessageMap map[string][]int


func getMessagesList() ([]string, error) {
	var data MessageMap

	raw, err := os.ReadFile("data/messages.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}

	return data.getList(), nil
}

func (data MessageMap) getList() (keys MessagesList) {
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}

func GetMessage(messageName string)(string, error){
	path := fmt.Sprintf("data/messages/%s.md", messageName)
	data, err := os.ReadFile(path)
	if err != nil{
		return "", err
	}
	return  string(data), nil
}
