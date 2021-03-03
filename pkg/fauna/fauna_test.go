package fauna_test

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Trois-Six/fauna-exporter/pkg/fauna"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func randInt64() int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000))

	return n.Int64()
}

func mockDashBoardAuthAPI(data fauna.Authentication, status int) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(status)
			if err := json.NewEncoder(rw).Encode(data); err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		}))
}

func mockDashBoardBillingAPI(data fauna.Billing, status int) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(status)
			if err := json.NewEncoder(rw).Encode(data); err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		}))
}

func mockDashBoardUsageAPI(data fauna.Usage, status int) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(status)
			if err := json.NewEncoder(rw).Encode(data); err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		}))
}

func TestFauna_Login(t *testing.T) {
	dashboardAuthAPIMock := mockDashBoardAuthAPI(fauna.Authentication{
		ID:        "1234",
		SessionID: "1234",
		Secret:    "secret",
		User: fauna.User{
			Name:       "name",
			Email:      "foo.bar@example.org",
			ID:         "1234",
			OTPEnabled: false,
			Role:       "role",
		},
		Account: fauna.Account{
			CompanyName:   "company",
			LegacyAccount: false,
		},
	}, http.StatusOK)
	defer dashboardAuthAPIMock.Close()

	f := fauna.NewFaunaClient("foo.bar@example.org", "password")
	err := f.Login(dashboardAuthAPIMock.URL)
	require.NoError(t, err)
	assert.Equal(t, "secret", f.Secret)
}

func TestFauna_GetBillingUsage(t *testing.T) {
	b := fauna.Billing{
		StartPeriod: "1970-01-01T00:00:00Z",
		EndPeriod:   "1970-01-02T00:00:00Z",
		LineItems: []fauna.BillingLineItem{{
			Description: "Business plan",
			Amount:      12500,
		}},
		TotalAmount: 12500,
		MetricAmount: fauna.BillingType{
			ByteReadOps:  randInt64(),
			ByteWriteOps: randInt64(),
			ComputeOps:   randInt64(),
			Storage:      randInt64(),
		},
		MetricUsage: fauna.BillingType{
			ByteReadOps:  randInt64(),
			ByteWriteOps: randInt64(),
			ComputeOps:   randInt64(),
			Storage:      randInt64(),
		},
	}
	dashboardAuthAPIMock := mockDashBoardBillingAPI(b, http.StatusOK)

	defer dashboardAuthAPIMock.Close()

	f := fauna.NewFaunaClient("foo.bar@example.org", "password")
	err := f.Login(dashboardAuthAPIMock.URL)
	require.NoError(t, err)

	billing, err := f.GetBillingUsage(dashboardAuthAPIMock.URL, 7)
	require.NoError(t, err)
	assert.Equal(t, billing, b)
}

func TestFauna_GetUsage(t *testing.T) {
	u := fauna.Usage{
		"collectionA": {
			ByteReadOps:  randInt64(),
			ByteWriteOps: randInt64(),
			ComputeOps:   randInt64(),
			Storage:      randInt64(),
			Versions:     randInt64(),
			Indexes:      randInt64(),
		},
		"collectionB": {
			ByteReadOps:  randInt64(),
			ByteWriteOps: randInt64(),
			ComputeOps:   randInt64(),
			Storage:      randInt64(),
			Versions:     randInt64(),
			Indexes:      randInt64(),
		},
	}
	dashboardAuthAPIMock := mockDashBoardUsageAPI(u, http.StatusOK)

	defer dashboardAuthAPIMock.Close()

	f := fauna.NewFaunaClient("foo.bar@example.org", "password")
	err := f.Login(dashboardAuthAPIMock.URL)
	require.NoError(t, err)

	usage, err := f.GetUsage(dashboardAuthAPIMock.URL, 7)
	require.NoError(t, err)
	assert.Equal(t, usage, u)
}
