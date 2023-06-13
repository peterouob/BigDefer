package utils

import "fmt"

func BlackReport(msg string) {
	switch msg {
	case "辱罵他人":
		checkMsgFromUser(msg)
	}
}

func checkMsgFromUser(msg string) {
	fmt.Println(msg)
}
