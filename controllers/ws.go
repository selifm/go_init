package controllers

/*
 * @Script: ws.go
 * @Author: pangxiaobo
 * @Email: 10846295@qq.com
 * @Create At: 2018-11-29 10:36:11
 * @Last Modified By: pangxiaobo
 * @Last Modified At: 2018-11-29 11:09:15
 * @Description: This is description.
 */

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type WsController struct{}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	HandshakeTimeout:  5 * time.Second,
	// CheckOrigin: 处理跨域问题，线上环境慎用
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ws *WsController) WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	go echo(conn)
}

func echo(conn *websocket.Conn) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		if string(msg) == "ping" {
			fmt.Println("ping")
			time.Sleep(time.Second * 2)
			err = conn.WriteMessage(msgType, []byte("pong"))
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			conn.Close()
			fmt.Println(string(msg))
			return
		}
	}
}