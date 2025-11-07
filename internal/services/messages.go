package services

import (
	"fmt"
	"os"
)

type MessagesList []string

type MessageData struct {
	Timing    []int    `json:"timing"`
	URLButton []string `json:"url_button"`
}
type MessageMap map[string]MessageData

type Messages struct {
	MessagesList MessagesList `json:"messages_list"`
	Messages     MessageMap   `json:"messages"`
}

func getMessages() (Messages, error) {
	messages, err := ReadJSONRetry[Messages]("data/messages.json", 3)
	if err != nil {
		return messages, err
	}
	return messages, nil
}

func getMessageMap() (MessageMap, error) {
	messages, err := getMessages()
	if err != nil {
		return messages.Messages, err
	}
	return messages.Messages, nil
}

func getMessageData(messageName string) (MessageData, error) {
	var messageData MessageData
	messageMap, err := getMessageMap()
	if err != nil {
		return messageData, err
	}
	return messageMap[messageName], nil
}

func getMessagesList() (MessagesList, error) {
	var messagesList MessagesList
	messages, err := getMessages()
	if err != nil {
		return messagesList, err
	}
	return messages.MessagesList.reverse(), nil
}

func (messagesList MessagesList) reverse() MessagesList {
	for i := 0; i < len(messagesList)/2; i++ {
		j := len(messagesList) - 1 - i
		messagesList[i], messagesList[j] = messagesList[j], messagesList[i]
	}
	return messagesList
}

func GetMessageText(messageName string) (string, error) {
	path := fmt.Sprintf("data/messages/%s.md", messageName)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func GetTiming(messageName string) ([]int, error) {
	messageData, err := getMessageData(messageName)
	if err != nil {
		return nil, err
	}
	return messageData.Timing, nil
}

func GetURLButton(messageName string) (string, string, error) {
	messageData, err := getMessageData(messageName)
	if err != nil {
		return "", "", err
	}
	URL_button := messageData.URLButton
	return URL_button[0], URL_button[1], nil
}

func LastMessage(messagesList MessagesList) (string, error) {
	n := len(messagesList)
	if n == 0 {
		return "", fmt.Errorf("messagesList is empty")
	}
	last := messagesList[n-1]
	return last, nil
}

func GetPhoto(messageName string) (string, error) {
	path := fmt.Sprintf("assets/images/%v.PNG", messageName)
	_, err := os.Stat(path)
	if err == nil || !os.IsNotExist(err) {
		return path, nil
	}
	return "", fmt.Errorf("failed to get photo: %w", err)
}
