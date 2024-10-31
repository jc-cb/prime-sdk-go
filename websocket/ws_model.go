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

type WebSocketSubscribeRequest struct {
	L2     bool `json:"l2"`
	Orders bool `json:"orders"`
}

type WebSocketSubscribeResponse struct {
}

type WebSocketUnsubscribeRequest struct {
	L2     bool `json:"l2"`
	Orders bool `json:"orders"`
}

type WebSocketUnsubscribeResponse struct {
}

type WebSocketSubscribeMessage struct {
	Type         string   `json:"type"`
	Channel      string   `json:"channel"`
	AccessKey    string   `json:"access_key"`
	SvcAccountId string   `json:"api_key_id"`
	Timestamp    string   `json:"timestamp"`
	Passphrase   string   `json:"passphrase"`
	Signature    string   `json:"signature"`
	PortfolioId  string   `json:"portfolio_id,omitempty"`
	ProductIds   []string `json:"product_ids,omitempty"`
}

type WebSocketMessage struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
}

type WebSocketErrorMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type WebSocketUnsubscribeMessage struct {
	Type     string   `json:"type"`
	Channels []string `json:"channels"`
}

type WebSocketHeartbeatMessage struct {
	Channel     string                     `json:"channel"`
	Timestamp   string                     `json:"timestamp"`
	SequenceNum int                        `json:"sequence_num"`
	Events      []*WebSocketHeartbeatEvent `json:"events"`
}

type WebSocketHeartbeatEvent struct {
	Subscriptions *WebSocketHeartbeatEventSubscriptions `json:"subscriptions"`
}

type WebSocketHeartbeatEventSubscriptions struct {
	Heartbeats []string `json:"heartbeats"`
}

type WebSocketOrdersMessage struct {
	Channel     string                  `json:"channel"`
	Timestamp   string                  `json:"timestamp"`
	SequenceNum int                     `json:"sequence_num"`
	Events      []*WebSocketOrdersEvent `json:"events"`
}

type WebSocketOrdersSubscriptionMessage struct {
	Channel     string                              `json:"channel"`
	Timestamp   string                              `json:"timestamp"`
	SequenceNum int                                 `json:"sequence_num"`
	Events      []*WebSocketOrdersSubscriptionEvent `json:"events"`
}

type WebSocketOrdersSubscriptionEvent struct {
	Subscriptions WebSocketOrdersSubscriptionEventData `json:"subscriptions"`
}

type WebSocketOrdersSubscriptionEventData struct {
	// These are portfolio ids subscribed to for orders. Json name is orders
	PortfolioIds []string `json:"orders"`
}

type WebSocketOrdersEvent struct {
	Type   string            `json:"type"`
	Orders []*WebSocketOrder `json:"orders"`
}

type WebSocketOrder struct {
	OrderId            string `json:"order_id"`
	ClientOrderId      string `json:"client_order_id"`
	CumulativeQuantity string `json:"cum_qty"`
	LeavesQuantity     string `json:"leaves_qty"`
	AveragePrice       string `json:"avg_px"`
	Fees               string `json:"fees"`
	Status             string `json:"status"`
}

type WebSocketL2Message struct {
	Channel     string              `json:"channel"`
	Timestamp   string              `json:"timestamp"`
	SequenceNum int                 `json:"sequence_num"`
	Events      []*WebSocketL2Event `json:"events"`
}

type WebSocketL2Event struct {
	Type      string               `json:"type"`
	ProductId string               `json:"product_id"`
	Updates   []*WebSocketL2Update `json:"updates"`
}

type WebSocketL2Update struct {
	Side      string `json:"side"`
	EventTime string `json:"event_time"`
	Price     string `json:"px"`
	Quantity  string `json:"qty"`
}

type WebSocketL2SubscriptionMessage struct {
	Channel     string                          `json:"channel"`
	Timestamp   string                          `json:"timestamp"`
	SequenceNum int                             `json:"sequence_num"`
	Events      []*WebSocketL2SubscriptionEvent `json:"events"`
}

type WebSocketL2SubscriptionEvent struct {
	Subscriptions WebSocketL2SubscriptionEventData `json:"subscriptions"`
}

type WebSocketL2SubscriptionEventData struct {
	// These are product ids subscribed to for l2 data. Json name is l2_data
	ProductIds []string `json:"l2_data"`
}
