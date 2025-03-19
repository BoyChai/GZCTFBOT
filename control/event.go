package control

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// 获取最新的通知 ID
func getLatestID(url string, client *http.Client) (int64, error) {
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var events []Data
	err = json.Unmarshal(body, &events)
	if err != nil {
		return 0, err
	}

	if len(events) == 0 {
		return 0, nil
	}

	// 找到最大 ID
	maxID := events[0].ID
	for _, event := range events {
		if event.ID > maxID {
			maxID = event.ID
		}
	}
	return maxID, nil
}

func StartEvent(baseURL, gameID string, interval int) {
	url := fmt.Sprintf("%s/api/game/%s/notices", baseURL, gameID)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 在启动时获取当前最新的 ID
	lastID, err := getLatestID(url, client)
	if err != nil {
		fmt.Printf("获取初始最新ID失败: %v\n", err)
		lastID = 0 // 如果失败，从0开始
	} else {
		fmt.Printf("初始化完成，当前最新通知ID: %d\n", lastID)
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	fmt.Printf("开始监听 %s 的通知...\n", url)

	for {
		select {
		case <-ticker.C:
			resp, err := client.Get(url)
			if err != nil {
				fmt.Printf("请求失败: %v\n", err)
				continue
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Printf("读取响应失败: %v\n", err)
				continue
			}

			var events []Data
			err = json.Unmarshal(body, &events)
			if err != nil {
				fmt.Printf("解析 JSON 失败: %v\n", err)
				continue
			}

			for _, event := range events {
				if event.ID > lastID {
					if msg := formatData(event); msg != "" {
						fmt.Println(msg)
					}
					lastID = event.ID
				}
			}
		}
	}
}
