package session

import (
	"testing"

	stripe "github.com/max-cape/test-stripe-go"
	_ "github.com/max-cape/test-stripe-go/testing"
	assert "github.com/stretchr/testify/require"
)

func TestBillingPortalSessionNew(t *testing.T) {
	session, err := New(&stripe.BillingPortalSessionParams{
		Customer:  stripe.String("cus_123"),
		ReturnURL: stripe.String("https://stripe.com/return"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, session)
}
