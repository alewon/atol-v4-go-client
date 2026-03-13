package atol

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const DefaultBaseURL = "https://online.atol.ru/possystem/v4"

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

type ErrorResponse struct {
	StatusCode int
	Body       string
}

func (e *ErrorResponse) Error() string {
	if e == nil {
		return ""
	}

	if strings.TrimSpace(e.Body) == "" {
		return fmt.Sprintf("unexpected http status: %d", e.StatusCode)
	}

	return fmt.Sprintf("unexpected http status: %d: %s", e.StatusCode, e.Body)
}

func NewClient(baseURL string, httpClient *http.Client) *Client {
	if strings.TrimSpace(baseURL) == "" {
		baseURL = DefaultBaseURL
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		BaseURL:    strings.TrimRight(baseURL, "/"),
		HTTPClient: httpClient,
	}
}

func (c *Client) newRequest(ctx context.Context, method string, path string, query url.Values, body any) (*http.Request, error) {
	baseURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	baseURL.Path = strings.TrimRight(baseURL.Path, "/") + "/" + strings.TrimLeft(path, "/")

	if query != nil {
		baseURL.RawQuery = query.Encode()
	}

	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		payload, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		requestBody = bytes.NewReader(payload)
	}

	req, err := http.NewRequestWithContext(ctx, method, baseURL.String(), requestBody)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) do(req *http.Request, out any) error {
	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return errors.Join(&ErrorResponse{StatusCode: resp.StatusCode}, readErr)
		}

		return &ErrorResponse{
			StatusCode: resp.StatusCode,
			Body:       string(body),
		}
	}

	if out == nil {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return nil
}

func (c *Client) tokenQuery() url.Values {
	if strings.TrimSpace(c.Token) == "" {
		return nil
	}

	query := url.Values{}
	query.Set("token", c.Token)

	return query
}

func (c *Client) PostToken(ctx context.Context, request PostTokenRequest) (PostTokenResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, "getToken", nil, request)
	if err != nil {
		return PostTokenResponse{}, err
	}

	var response PostTokenResponse
	err = c.do(req, &response)

	return response, err
}

func (c *Client) GetToken(ctx context.Context, request GetTokenRequest) (GetTokenResponse, error) {
	query := url.Values{}
	query.Set("login", request.Login)
	query.Set("pass", request.Pass)

	if request.Source != "" {
		query.Set("source", request.Source)
	}

	req, err := c.newRequest(ctx, http.MethodGet, "getToken", query, nil)
	if err != nil {
		return GetTokenResponse{}, err
	}

	var response GetTokenResponse
	err = c.do(req, &response)

	return response, err
}

func (c *Client) Sell(ctx context.Context, groupCode string, request SellRequest) (SellResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s", groupCode, "sell"), c.tokenQuery(), request)
	if err != nil {
		return SellResponse{}, err
	}

	var response SellResponse
	err = c.do(req, &response)

	return response, err
}

func (c *Client) SellRefund(ctx context.Context, groupCode string, request SellRefundRequest) (SellRefundResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s", groupCode, "sell_refund"), c.tokenQuery(), request)
	if err != nil {
		return SellRefundResponse{}, err
	}

	var response SellRefundResponse
	err = c.do(req, &response)

	return response, err
}

func (c *Client) Buy(ctx context.Context, groupCode string, request BuyRequest) (BuyResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s", groupCode, "buy"), c.tokenQuery(), request)
	if err != nil {
		return BuyResponse{}, err
	}

	var response BuyResponse
	err = c.do(req, &response)

	return response, err
}

func (c *Client) BuyRefund(ctx context.Context, groupCode string, request BuyRefundRequest) (BuyRefundResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s", groupCode, "buy_refund"), c.tokenQuery(), request)
	if err != nil {
		return BuyRefundResponse{}, err
	}

	var response BuyRefundResponse
	err = c.do(req, &response)

	return response, err
}

func (c *Client) SellCorrection(ctx context.Context, groupCode string, request SellCorrectionRequest) (SellCorrectionResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s", groupCode, "sell_correction"), c.tokenQuery(), request)
	if err != nil {
		return SellCorrectionResponse{}, err
	}

	var response SellCorrectionResponse
	err = c.do(req, &response)

	return response, err
}

func (c *Client) BuyCorrection(ctx context.Context, groupCode string, request BuyCorrectionRequest) (BuyCorrectionResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s", groupCode, "buy_correction"), c.tokenQuery(), request)
	if err != nil {
		return BuyCorrectionResponse{}, err
	}

	var response BuyCorrectionResponse
	err = c.do(req, &response)

	return response, err
}

func (c *Client) Report(ctx context.Context, groupCode string, uuid string) (ReportResponse, error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("%s/report/%s", groupCode, uuid), c.tokenQuery(), nil)
	if err != nil {
		return ReportResponse{}, err
	}

	var response ReportResponse
	err = c.do(req, &response)

	return response, err
}
