package services

import (
	"encoding/json"
	"log"
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

	date := time.Until(sendTime)
	if date <= 0 {
		sendMessage(chatID)
		return
	}

	time.AfterFunc(date, func() {
		sendMessage(chatID)
	})
}

func SetNextSchedule(chatID string, sendMessage func(string)) {
	user, err := GetPerson(chatID)
	if err != nil {
		return
	}

	messagesList := user.MessagesList
	n := len(messagesList)
	if n == 0{
		return
	}
	last := messagesList[n-1]

	timing, err := GetTiming(last)
	if err != nil {
		log.Printf("timing fetching error: %s", err)
		return
	}

	sendTimeStr, err := setSendTime(user.RegTime, timing)
	if err != nil{
		log.Printf("sendTime error: %s", err)
		return
	}

	setSchedule(sendTimeStr, chatID, sendMessage)
}

func setSendTime(regTimeStr string, timing []int) (string, error) {
	regTime, err := time.Parse(time.RFC3339, regTimeStr)
	if err != nil {
		return "", err
	}

	sendTime := regTime.Add(time.Duration(timing[0])*time.Hour +
		time.Duration(timing[1])*time.Minute)

	sendTimeStr := sendTime.Format(time.RFC3339)
	return sendTimeStr, nil
}
