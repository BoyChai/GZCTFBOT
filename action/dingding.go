package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var dingDingChan chan string

var once sync.Once

type Text struct {
	Content string `json:"content"`
}

type RobotSendRequest struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
}

func NewDingDing() {
	once.Do(func() {
		dingDingChan = newChan()
		go sendDingTalkWebhookMessage()
	})
}
func sendDingTalkWebhookMessage() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for content := range dingDingChan {
		request := RobotSendRequest{
			MsgType: "text",
			Text: Text{
				Content: content,
			},
		}

		jsonData, err := json.Marshal(request)
		if err != nil {
			fmt.Printf("Failed to marshal request: %v\n", err)
			continue
		}

		webhooks := viper.GetStringSlice("DingDingBot.Webhook")
		if len(webhooks) == 0 {
			fmt.Println("No DingDingBot.Webhook configured")
			continue
		}

		var wg sync.WaitGroup
		for _, webhook := range webhooks {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()

				body := bytes.NewBuffer(jsonData)

				req, err := http.NewRequest("POST", url, body)
				if err != nil {
					fmt.Printf("Failed to create request for %s: %v\n", url, err)
					return
				}

				req.Header.Set("Content-Type", "application/json")

				resp, err := client.Do(req)
				if err != nil {
					fmt.Printf("Failed to send to %s: %v\n", url, err)
					return
				}
				defer resp.Body.Close()

				bodyBytes, _ := io.ReadAll(resp.Body)
				if resp.StatusCode != http.StatusOK {
					fmt.Printf("Unexpected status: %s, response: %s\n",
						resp.Status, string(bodyBytes))
					return
				}
				fmt.Printf("Successfully status: %s, response: %s\n",
					resp.Status, string(bodyBytes))
			}(webhook)
		}
		wg.Wait()
	}
}
