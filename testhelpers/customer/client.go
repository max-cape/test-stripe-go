//
//
// File generated from our OpenAPI spec
//
//

// Package customer provides the /v1/customers APIs
package customer

import (
	"net/http"

	stripe "github.com/max-cape/test-stripe-go"
)

// Client is used to invoke /v1/customers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Create an incoming testmode bank transfer
func FundCashBalance(id string, params *stripe.TestHelpersCustomerFundCashBalanceParams) (*stripe.CustomerCashBalanceTransaction, error) {
	return getC().FundCashBalance(id, params)
}

// Create an incoming testmode bank transfer
func (c Client) FundCashBalance(id string, params *stripe.TestHelpersCustomerFundCashBalanceParams) (*stripe.CustomerCashBalanceTransaction, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/customers/%s/fund_cash_balance", id)
	customercashbalancetransaction := &stripe.CustomerCashBalanceTransaction{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, customercashbalancetransaction)
	return customercashbalancetransaction, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
