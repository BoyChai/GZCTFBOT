package control

import (
	"fmt"
	"time"
)

// formatTime 将时间戳格式化为可读字符串
func formatTime(timestamp int64) string {
	return time.UnixMilli(timestamp).Format("2006-01-02 15:04:05")
}

// formatFirstBlood 格式化 FirstBlood 事件（队伍在一道题上拿一血）
func formatFirstBlood(e Data) string {
	if len(e.Values) != 2 {
		fmt.Println("无效的 FirstBlood 数据: %v", e.Values)
		return ""
	}
	return fmt.Sprintf("[解题播报] 恭喜队伍 %s 拿到 %s 一血\n时间: %s",
		e.Values[0], e.Values[1], formatTime(e.Time))
}

// formatSecondBlood 格式化 SecondBlood 事件（队伍在一道题上拿二血）
func formatSecondBlood(e Data) string {
	if len(e.Values) != 2 {
		fmt.Println("无效的 SecondBlood 数据: %v", e.Values)
		return ""
	}
	return fmt.Sprintf("[解题播报] 恭喜队伍 %s 拿到 %s 二血\n时间: %s",
		e.Values[0], e.Values[1], formatTime(e.Time))
}

// formatThirdBlood 格式化 ThirdBlood 事件（队伍在一道题上拿三血）
func formatThirdBlood(e Data) string {
	if len(e.Values) != 2 {
		fmt.Println("无效的 ThirdBlood 数据: %v", e.Values)
		return ""
	}
	return fmt.Sprintf("[解题播报] 恭喜队伍 %s 拿到 %s 三血\n时间: %s",
		e.Values[0], e.Values[1], formatTime(e.Time))
}

// formatNewHint 格式化 NewHint 事件（新提示）
func formatNewHint(e Data) string {
	if len(e.Values) != 1 {
		fmt.Println("无效的 NewHint 数据: %v", e.Values)
		return ""
	}
	return fmt.Sprintf("[提示播报] 题目 %s 发布新提示\n时间: %s",
		e.Values[0], formatTime(e.Time))
}

// formatNewChallenge 格式化 NewChallenge 事件（新挑战）
func formatNewChallenge(e Data) string {
	if len(e.Values) != 1 {
		fmt.Println("无效的 NewChallenge 数据: %v", e.Values)
		return ""
	}
	return fmt.Sprintf("[题目播报] 题目 %s 已上线\n时间: %s",
		e.Values[0], formatTime(e.Time))
}

// formatNormal 格式化 Normal 事件（通知事件）
func formatNormal(e Data) string {
	if len(e.Values) != 1 {
		fmt.Println("无效的 Normal 数据: %v", e.Values)
		return ""
	}
	return fmt.Sprintf("[赛事通知] \n%s\n时间: %s ",
		e.Values[0], formatTime(e.Time))
}

// formatData 通用的格式化函数，根据类型调用对应函数
func formatData(e Data) string {
	switch e.Type {
	case "FirstBlood":
		return formatFirstBlood(e)
	case "SecondBlood":
		return formatSecondBlood(e)
	case "ThirdBlood":
		return formatThirdBlood(e)
	case "NewHint":
		return formatNewHint(e)
	case "NewChallenge":
		return formatNewChallenge(e)
	case "Normal":
		return formatNormal(e)
	default:
		fmt.Println("未知事件类型: %s, 数据: %v", e.Type, e)
		return ""
	}
}
