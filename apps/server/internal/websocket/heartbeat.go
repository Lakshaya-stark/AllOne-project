package websocket

import "time"

const (

	PongWait = 60 * time.Second

	PingPeriod = PongWait * 9 / 10

	WriteWait = 10 * time.Second

	MaxMessageSize = 1024 * 1024
)