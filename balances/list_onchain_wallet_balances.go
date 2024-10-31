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

package balances

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
)

type ListOnchainWalletBalancesRequest struct {
	PortfolioId         string            `json:"portfolio_id"`
	WalletId            string            `json:"wallet_id"`
	VisiblilityStatuses []string          `json:"visibility_statuses"`
	Pagination          *PaginationParams `json:"pagination_params"`
}

type ListOnchainWalletBalancesResponse struct {
	Balances              []*Balance                        `json:"balances"`
	Type                  string                            `json:"type"`
	TradingWalletBalances *BalanceWithHolds                 `json:"trading_balances"`
	VaultWalletBalances   *BalanceWithHolds                 `json:"vault_balances"`
	Request               *ListOnchainWalletBalancesRequest `json:"request"`
}

func (s *balancesServiceImpl) ListOnchainWalletBalances(
	ctx context.Context,
	request *ListOnchainWalletBalancesRequest,
) (*ListOnchainWalletBalancesResponse, error) {

	path := fmt.Sprintf(
		"/portfolios/$s/wallets/%s/web3_balances",
		request.PortfolioId,
	)

	var queryParams string

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	for _, v := range request.VisiblilityStatuses {
		queryParams = core.AppendHttpQueryParam(queryParams, "visibility_statuses", v)
	}

	response := &ListOnchainWalletBalancesResponse{Request: request}

	if err := core.HttpGet(ctx, s.client, path, queryParams, successStatusCodes, request, response, s.client.HeadersFunc()); err != nil {
		return nil, err
	}

	return response, nil
}
