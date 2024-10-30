/**
 * Copyright 2023-present Coinbase Global, Inc.
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

package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/credentials"
)

const defaultWebSocketUrl = "wss://ws-feed.prime.coinbase.com"

type WebSocketClient interface {
	Credentials() *credentials.Credentials

	// SetWebSocketUrl sets the URL for the WebSocket connection. It overrides
	// anything passed in the DialerConfig if that is set.
	SetWebSocketUrl(u string) WebSocketClient
	WebSocketUrl() string

	SetWebSocketDialTimeout(seconds time.Duration) WebSocketClient
	WebSocketDialTimeout() time.Duration

	// SetWebSocketDialerConfig sets the WebSocket dialer config. You must call
	// SetWebSocketUrl if you to override the default URL.
	SetWebSocketDialerConfig(c core.DialerConfig) WebSocketClient
}

type webSocketClientImpl struct {
	credentials  *credentials.Credentials
	dialerConfig core.DialerConfig
	url          string
	dialTimeout  time.Duration
}

func NewWebSocketClient(credentials *credentials.Credentials) WebSocketClient {
	return &webSocketClientImpl{
		url:          defaultWebSocketUrl,
		credentials:  credentials,
		dialerConfig: core.DefaultDialerConfig(defaultWebSocketUrl),
		dialTimeout:  5 * time.Second,
	}
}

func (c *webSocketClientImpl) SetWebSocketDialerConfig(conf core.DialerConfig) WebSocketClient {
	if len(conf.Url) == 0 || conf.Url != c.url {
		conf.Url = c.url
	}
	c.dialerConfig = conf
	return c
}

func (c *webSocketClientImpl) WebSocketUrl() string {
	return c.url
}

func (c *webSocketClientImpl) SetWebSocketUrl(url string) WebSocketClient {
	c.url = url
	c.dialerConfig.Url = url
	return c
}

func (c *webSocketClientImpl) SetWebSocketDialTimeout(t time.Duration) WebSocketClient {
	c.dialTimeout = t
	return c
}

func (c *webSocketClientImpl) WebSocketDialTimeout() time.Duration {
	return c.dialTimeout
}

func (c *webSocketClientImpl) Credentials() *credentials.Credentials {
	return c.credentials
}

func signWebSocket(channel, portfolioId, svcAccountId, timestamp, key, signingKey string, productIds []string) string {
	h := hmac.New(sha256.New, []byte(signingKey))
	message := fmt.Sprintf("%s%s%s%s%s%s", channel, key, svcAccountId, timestamp, portfolioId, strings.Join(productIds, ""))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
