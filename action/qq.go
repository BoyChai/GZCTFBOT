package action

import (
	"fmt"
	"sync"

	actionv1 "github.com/BoyChai/CoralBot/action.v1"
	"github.com/BoyChai/CoralBot/structure"
	"github.com/spf13/viper"
)

var action actionv1.Onebot11Action

var qqChan chan string

var qqOnce sync.Once

func NewQQ() {
	qqOnce.Do(func() {
		qqChan = newChan()
		go sendQqCoralBotMessage()
	})
	action.Agreement = "http"
	action.Host = viper.GetString("QQBot.OneBot")
}

func sendQqCoralBotMessage() {

	for content := range qqChan {

		groups := viper.GetIntSlice("QQBot.Group")
		if len(groups) == 0 {
			fmt.Println("No QQBot.Group configured")
			continue
		}

		var wg sync.WaitGroup
		for _, group := range groups {
			wg.Add(1)
			go func(group int) {
				defer wg.Done()
				data, err := action.SendMsg("group", structure.QQMsg{
					GroupId: int64(group),
					Message: content,
				})
				if err != nil {
					fmt.Println("发送失败："+err.Error(), "返回内容", data)
					return
				}
			}(group)
		}
		wg.Wait()
	}
}
