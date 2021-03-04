// Package fauna contains code to request Fauna API.
package fauna

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	// DashboardAuthAPI is the SaaS dashboard authentication API.
	DashboardAuthAPI = "https://auth.console.fauna.com/login"
	// DashboardBillingAPI is the SaaS dashboard billing API.
	DashboardBillingAPI = "https://auth.console.fauna.com/billing"
	// DashboardUsageAPI is the SaaS dashboard usage API.
	DashboardUsageAPI = "https://auth.console.fauna.com/usage"
	// FaunaClientTimeout is the timeout of the http client to connect to the Fauna API.
	FaunaClientTimeout = 10
)

// Client holds information to make requests to Fauna.
type Client struct {
	Email    string
	Password string
	Secret   string
	Client   *http.Client
}

// User holds user information given by the Fauna API.
type User struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	ID         string `json:"id"`
	OTPEnabled bool   `json:"otp_enabled"`
	Role       string `json:"role"`
}

// Account holds account information given by the Fauna API.
type Account struct {
	CompanyName   string `json:"company_name"`
	LegacyAccount bool   `json:"legacy_account"`
}

// Authentication holds authentication information after sign-in.
type Authentication struct {
	ID        string  `json:"id"`
	SessionID string  `json:"session_id"`
	Secret    string  `json:"secret"`
	User      User    `json:"user"`
	Account   Account `json:"account"`
}

// BillingLineItem is a plan.
type BillingLineItem struct {
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
}

// BillingType is a type of cost.
type BillingType struct {
	ByteReadOps  int64 `json:"byte_read_ops"`
	ByteWriteOps int64 `json:"byte_write_ops"`
	ComputeOps   int64 `json:"compute_ops"`
	Storage      int64 `json:"storage"`
}

// Billing contains billing info.
type Billing struct {
	StartPeriod  string            `json:"start_period"`
	EndPeriod    string            `json:"end_period"`
	LineItems    []BillingLineItem `json:"line_items"`
	TotalAmount  int64             `json:"total_amount"`
	MetricAmount BillingType       `json:"metric_amount"`
	MetricUsage  BillingType       `json:"metric_usage"`
}

// UsageType is a type of usage.
type UsageType struct {
	ByteReadOps  int64 `json:"byte_read_ops"`
	ByteWriteOps int64 `json:"byte_write_ops"`
	ComputeOps   int64 `json:"compute_ops"`
	Storage      int64 `json:"storage"`
	Versions     int64 `json:"versions"`
	Indexes      int64 `json:"indexes"`
}

// Usage collects all the usages.
type Usage map[string]UsageType

// NewFaunaClient creates a Fauna client.
func NewFaunaClient(email, password string) *Client {
	return &Client{
		Email:    email,
		Password: password,
		Secret:   "",
		Client: &http.Client{
			Timeout:       FaunaClientTimeout * time.Second,
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
		},
	}
}

// Login permits to get the secret used to exchange with Fauna API.
func (f *Client) Login(url string) error {
	data, err := json.Marshal(map[string]string{
		"email":    f.Email,
		"password": f.Password,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal credentials: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-type", "application/json")

	resp, err := f.Client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to request login API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to get body from login response: %w", err)
	}

	var auth Authentication
	err = json.Unmarshal(body, &auth)

	if err != nil {
		return fmt.Errorf("failed to get authentication information from response: %w", err)
	}

	f.Secret = auth.Secret

	return nil
}

// GetBillingUsage returns the billing usage for the chosen number of days.
func (f *Client) GetBillingUsage(url string, days int) (Billing, error) {
	body, err := f.processGetRequest(url + "?mode=preview&days=" + strconv.FormatInt(int64(days), 10))
	if err != nil {
		return Billing{}, err
	}

	var billing Billing
	err = json.Unmarshal(body, &billing)

	if err != nil {
		return Billing{}, fmt.Errorf("failed to unmarshal billing information: %w", err)
	}

	return billing, nil
}

// GetUsage returns the metrics of all collections for the chosen number of days.
func (f *Client) GetUsage(url string, days int) (Usage, error) {
	body, err := f.processGetRequest(url + "?days=" + strconv.FormatInt(int64(days), 10))
	if err != nil {
		return Usage{}, err
	}

	var usage Usage
	err = json.Unmarshal(body, &usage)

	if err != nil {
		return Usage{}, fmt.Errorf("failed to unmarshal usage information: %w", err)
	}

	return usage, nil
}

func (f *Client) processGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("secret", f.Secret)

	resp, err := f.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request Fauna API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to get body from login response: %w", err)
	}

	return body, nil
}
