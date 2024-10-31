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

/*
	// SetErrorCallback registers a callback when WebSocket errors or server errors are received.
	SetErrorCallback(callback WebSockeErrorCallback) Client

	SetWebSockeL2Callback(callback WebSockeL2Callback) Client
	SetWebSockeOrdersCallback(callback WebSockeOrdersCallback) Client

	// SetWebSocketProductIds defines the supported product IDs for WebSockets. This must
	// be set before subscribing to any WebSocket channels. This cannot be changed once the
	// a subscription has been created.
	SetWebSocketProductIds(productIds []string) Client

	WebSocketSubscribe(req *WebSocketSubscribeRequest) (*WebSocketSubscribeResponse, error)
	WebSocketUnsubscribe(req *WebSocketUnsubscribeRequest) (error, *WebSocketUnsubscribeResponse)
*/

type WebSockeErrorCallback func(event *WebSocketErrorMessage)

// The WebSocket l2 data callback. This function must return ASAP because
// it will block receiving additional messages. The best practice is to place the message
// on an ordered queue and process asynchronously.
type WebSockeL2Callback func(event *WebSocketL2Message)

// The WebSocket orders listener callback. This function must return ASAP because
// it will block receiving additional messages. The best practice is to place the message
// on an ordered queue and process asynchronously.
type WebSockeOrdersCallback func(event *WebSocketOrdersMessage)

// The WebSocket heartbeat callback. This function must return ASAP because
// it will block receiving additional messages. The best practice is to place the message
// on an ordered queue and process asynchronously.
type webSockeHeartbeatCallback func(event *WebSocketHeartbeatMessage)

func (c *clientImpl) WebSocketSubscribe(req *WebSocketSubscribeRequest) (*WebSocketSubscribeResponse, error) {
	return nil, nil
}

func (c *clientImpl) WebSocketUnsubscribe(req *WebSocketUnsubscribeRequest) (error, *WebSocketUnsubscribeResponse) {

	return nil, nil
}

// Heartbeat channels are opened for product IDs subscribed via the Orders or L2 channels.
func (c *clientImpl) webSocketHeartbeatsSubscribe(callback webSockeHeartbeatCallback) error {

	// TODO
	return nil
}

// This is called internally when there are Orders or L2 subscriptions for
// particular product IDs.
func (c *clientImpl) webSocketHeartbeatsUnsubscribe() error {
	// TODO
	return nil
}

func (c *clientImpl) SetErrorCallback(callback WebSockeErrorCallback) Client {
	c.webSocket.errorCallback = callback
	return c
}

func (c *clientImpl) SetWebSockeL2Callback(callback WebSockeL2Callback) Client {
	c.webSocket.l2Callback = callback
	return c

}

func (c *clientImpl) SetWebSockeOrdersCallback(callback WebSockeOrdersCallback) Client {
	c.webSocket.ordersCallback = callback
	return c

}

func (c *clientImpl) SetWebSocketProductIds(productIds []string) Client {
	c.webSocket.productIds = productIds
	return c
}
