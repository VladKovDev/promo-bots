package services

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	ChatID string `json:"chatID"`
}

func SetSchedules(sendMessage func(string)) error {
	data, err := os.ReadFile("data/schedule_backup.json")
	if err != nil {
		return err
	}

	var tasks map[string]interface{}
	if err := json.Unmarshal(data, &tasks); err != nil {
		return err
	}
	for k, v := range tasks {
		sendTimeStr, ok := v.(string)
		if !ok {
			continue
		}
		setSchedule(sendTimeStr, k, sendMessage)
	}
	return nil
}

func setSchedule(sendTimeStr string, chatID string, sendMessage func(string)) {
	sendTime, err := time.Parse(time.RFC3339, sendTimeStr)
	if err != nil {
		return
	}

	delay := time.Until(sendTime)
	if delay <= 0 {
		sendMessage(chatID)
		return
	}

	time.AfterFunc(delay, func() {
		sendMessage(chatID)
	})
}
