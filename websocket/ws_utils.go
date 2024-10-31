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
	"encoding/json"
	"time"
)

func subscriptionMsg(credentials *Credentials, channel string, productIds []string) ([]byte, error) {
	t := time.Now().UTC().Format(time.RFC3339)

	msg := &WebSocketSubscribeMessage{
		Type:         "subscribe",
		Channel:      channel,
		AccessKey:    credentials.AccessKey,
		SvcAccountId: credentials.SvcAccountId,
		Timestamp:    t,
		Passphrase:   credentials.Passphrase,
		Signature: signWebSocket(
			channel,
			credentials.PortfolioId,
			credentials.SvcAccountId,
			t,
			credentials.AccessKey,
			credentials.SigningKey,
			productIds,
		),
		PortfolioId: credentials.PortfolioId,
	}

	if len(productIds) > 0 {
		msg.ProductIds = productIds
	}

	b, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return b, nil

}

func heartbeatSubscriptionMsg(credentials *Credentials, productIds []string) ([]byte, error) {
	return subscriptionMsg(credentials, "heartbeats", productIds)
}

func orderSubscriptionMsg(credentials *Credentials, productIds []string) ([]byte, error) {
	return subscriptionMsg(credentials, "orders", productIds)
}

func l2OSubscriptionMsg(credentials *Credentials, productIds []string) ([]byte, error) {
	return subscriptionMsg(credentials, "l2_data", productIds)
}
