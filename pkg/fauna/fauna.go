package fauna

import (
	"bytes"
	"encoding/json"
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
)

// Client holds information to make requests to Fauna.
type Client struct {
	Email    string
	Password string
	Secret   string
	Client   *http.Client
}

type user struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	ID         string `json:"id"`
	OTPEnabled bool   `json:"otp_enabled"`
	Role       string `json:"role"`
}

type account struct {
	CompanyName   string `json:"company_name"`
	LegacyAccount bool   `json:"legacy_account"`
}

type authentication struct {
	ID        string  `json:"id"`
	SessionID string  `json:"session_id"`
	Secret    string  `json:"secret"`
	User      user    `json:"user"`
	Account   account `json:"account"`
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
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Login permits to get the secret used to exchange with Fauna.
func (f *Client) Login(url string) error {
	data, err := json.Marshal(map[string]string{
		"email":    f.Email,
		"password": f.Password,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", "application/json")
	resp, err := f.Client.Do(req)
	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var auth authentication
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return err
	}

	f.Secret = auth.Secret

	return nil
}

// GetBillingUsage returns the billing usage for the chosen number of days.
func (f *Client) GetBillingUsage(url string, days int) (Billing, error) {
	req, err := http.NewRequest("GET", url+"?mode=preview&days="+strconv.FormatInt(int64(days), 10), nil)
	if err != nil {
		return Billing{}, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("secret", f.Secret)
	resp, err := f.Client.Do(req)
	if err != nil {
		return Billing{}, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Billing{}, err
	}

	var billing Billing
	err = json.Unmarshal(body, &billing)
	if err != nil {
		return Billing{}, err
	}

	return billing, nil
}

// GetUsage returns the metrics of all the collections for the chosen number of days.
func (f *Client) GetUsage(url string, days int) (Usage, error) {
	req, err := http.NewRequest("GET", url+"?days="+strconv.FormatInt(int64(days), 10), nil)
	if err != nil {
		return Usage{}, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("secret", f.Secret)
	resp, err := f.Client.Do(req)
	if err != nil {
		return Usage{}, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Usage{}, err
	}

	var usage Usage
	err = json.Unmarshal(body, &usage)
	if err != nil {
		return Usage{}, err
	}

	return usage, nil
}
