package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行っている一人のユーザを表します
type client struct {
	// socketはこのクライアントのためのwebsocketです
	socket *websocket.Conn
	// send はメッセージが送られるチャネルです
	send chan []byte
	// roomはこのクライアントが参加しているチャットルームです
	room *room
	
}
