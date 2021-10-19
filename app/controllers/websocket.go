package controllers

import (
	"github.com/revel/revel"
	"app/app/chatroom"
)

type WebSocket struct {
	*revel.Controller
}

func (c WebSocket) Room(user string) revel.Result {
	user = "hpscript"
	return c.Render(user)
}

func (c WebSocket) RoomSocket(user string, ws revel.ServerWebSocket) revel.Result {

	if ws == nil {
		return nil
	}

	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	chatroom.Join(user)
	defer chatroom.Leave(user)

	for _, event := range subscription.Archive {
		if ws.MessageSendJSON(&event) != nil {
			return nil
		}
	}

	newMessages := make(chan string)
	go func(){
		var msg string
		for {
			err := ws.MessageReceiveJSON(&msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	for {
		select {
		case event := <-subscription.New:
			if ws.MessageSendJSON(&event) != nil {
				return nil
			}
		case msg, ok := <-newMessages:
			if !ok {
				return nil
			}
			chatroom.Say(user, msg)
		}
	}
	return nil
}