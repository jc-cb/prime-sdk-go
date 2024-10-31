/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package prime

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/coinbase-samples/core-go"
)

const (
	webSocketHeartbeats = "heartbeats"
	webSocketOrders     = "orders"
	webSocketL2         = "l2_data"
	webSocketError      = "error"
)

type webSocket struct {
	sync.Mutex
	conn         *core.WebSocketConnection
	connected    bool
	dialerConfig core.DialerConfig
	url          string
	dialTimeout  time.Duration

	productIds []string

	l2Callback        WebSockeL2Callback
	ordersCallback    WebSockeOrdersCallback
	heartbeatCallback webSockeHeartbeatCallback
	errorCallback     WebSockeErrorCallback
}

func (s *webSocket) closeHandler(code int, text string) error {
	s.Lock()
	defer s.Unlock()

	// TODO: Send a close error message to the listener
	s.connected = false
	return nil
}

func (s *webSocket) subscribe() error {
	// Ensure the connection
	if err := s.connect(s.dialTimeout); err != nil {
		return err
	}

	// Ensure the listener is running
	go s.listenForWebSocketMessages()

	// Send the subscribe messages

	return nil
}

func (s *webSocket) messageMultiplexer(m []byte) {

	baseMsg := &WebSocketMessage{}

	if err := json.Unmarshal(m, baseMsg); err != nil {
		// TODO: call the error listener
	}

	switch baseMsg.Channel {
	case webSocketHeartbeats:

	case webSocketOrders:
		if s.ordersCallback != nil {
			msg := &WebSocketOrdersMessage{}
			if err := json.Unmarshal(m, msg); err != nil {
				// TODO: call the error listener
			} else {
				s.ordersCallback(msg)
			}
		}
	case webSocketL2:
		if s.l2Callback != nil {
			msg := &WebSocketL2Message{}
			if err := json.Unmarshal(m, msg); err != nil {
				// TODO: call the error listener
			} else {
				s.l2Callback(msg)
			}
		}

	default:
		switch baseMsg.Type {
		case webSocketError:
			// TODO: Call the error
		default:
			// TODO: Call the error
		}
	}
}

func (s *webSocket) listenForWebSocketMessages() {
	if err := core.ListenForWebSocketMessages(s.conn, s.messageMultiplexer); err != nil {
		// TODO: call the error listener

	}
}

func (s *webSocket) heartbeatListener(event *WebSocketHeartbeatMessage) {

	fmt.Println(event)
}

// connect creates a new WebSocket connection. If the connection is already
// open, it does nothing.
func (s *webSocket) connect(timeout time.Duration) error {
	s.Lock()
	defer s.Unlock()

	if s.connected {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := core.DialWebSocket(ctx, s.dialerConfig)
	if err != nil {
		return err
	}

	s.conn.SetCloseHandler(s.closeHandler)

	s.conn = conn
	s.connected = true

	return nil
}

func newWebSocket(url string) *webSocket {
	ws := &webSocket{
		dialerConfig: core.DefaultDialerConfig(url),
		dialTimeout:  5 * time.Second,
	}

	ws.heartbeatCallback = ws.heartbeatListener

	return ws
}
