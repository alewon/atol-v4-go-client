package atol

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestNewClient(t *testing.T) {
	client := NewClient("", nil)

	if client.BaseURL != DefaultBaseURL {
		t.Fatalf("unexpected base url: %q", client.BaseURL)
	}

	if client.HTTPClient == nil {
		t.Fatal("http client is nil")
	}

	client = NewClient("https://example.com/", http.DefaultClient)

	if client.BaseURL != "https://example.com" {
		t.Fatalf("unexpected trimmed base url: %q", client.BaseURL)
	}
}

func TestPostToken(t *testing.T) {
	var gotMethod string
	var gotPath string
	var gotBody PostTokenRequest

	client := NewClient("https://example.com", &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			gotMethod = r.Method
			gotPath = r.URL.Path

			if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
				t.Fatalf("decode request body: %v", err)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body: io.NopCloser(strings.NewReader(`{
				"token":"test-token",
				"error":null,
				"timestamp":"2026-03-13T10:20:30+03:00"
			}`)),
			}, nil
		}),
	})

	response, err := client.PostToken(context.Background(), PostTokenRequest{
		Login:  "login",
		Pass:   "password",
		Source: "source",
	})
	if err != nil {
		t.Fatalf("PostToken returned error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Fatalf("unexpected method: %q", gotMethod)
	}

	if gotPath != "/getToken" {
		t.Fatalf("unexpected path: %q", gotPath)
	}

	if gotBody.Login != "login" || gotBody.Pass != "password" || gotBody.Source != "source" {
		t.Fatalf("unexpected request body: %+v", gotBody)
	}

	if response.Token != "test-token" {
		t.Fatalf("unexpected token: %q", response.Token)
	}
}

func TestGetToken(t *testing.T) {
	var gotMethod string
	var gotPath string
	var gotLogin string
	var gotPass string
	var gotSource string

	client := NewClient("https://example.com", &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			gotMethod = r.Method
			gotPath = r.URL.Path
			gotLogin = r.URL.Query().Get("login")
			gotPass = r.URL.Query().Get("pass")
			gotSource = r.URL.Query().Get("source")

			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body: io.NopCloser(strings.NewReader(`{
				"token":"test-token-get",
				"error":null,
				"timestamp":"2026-03-13T10:20:30+03:00"
			}`)),
			}, nil
		}),
	})

	response, err := client.GetToken(context.Background(), GetTokenRequest{
		Login:  "login",
		Pass:   "password",
		Source: "source",
	})
	if err != nil {
		t.Fatalf("GetToken returned error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Fatalf("unexpected method: %q", gotMethod)
	}

	if gotPath != "/getToken" {
		t.Fatalf("unexpected path: %q", gotPath)
	}

	if gotLogin != "login" || gotPass != "password" || gotSource != "source" {
		t.Fatalf("unexpected query: login=%q pass=%q source=%q", gotLogin, gotPass, gotSource)
	}

	if response.Token != "test-token-get" {
		t.Fatalf("unexpected token: %q", response.Token)
	}
}

func TestSell(t *testing.T) {
	var gotMethod string
	var gotPath string
	var gotToken string
	var gotBody SellRequest

	client := NewClient("https://example.com", &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			gotMethod = r.Method
			gotPath = r.URL.Path
			gotToken = r.URL.Query().Get("token")

			if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
				t.Fatalf("decode request body: %v", err)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body: io.NopCloser(strings.NewReader(`{
				"uuid":"uuid-1",
				"status":"wait",
				"timestamp":"2026-03-13T10:20:31+03:00",
				"error":null
			}`)),
			}, nil
		}),
	})
	client.Token = "client-token"

	sum := 100.0
	vatSum := 0.0

	response, err := client.Sell(context.Background(), "group-code", SellRequest{
		Timestamp:  "2026-03-13T10:20:30+03:00",
		ExternalID: "order-1",
		Receipt: SellRequestReceipt{
			Client: SellRequestReceiptClient{
				Email: "user@example.com",
			},
			Company: SellRequestReceiptCompany{
				Email:          "shop@example.com",
				SNO:            "osn",
				INN:            "5544332219",
				PaymentAddress: "https://example.com",
			},
			Items: []SellRequestReceiptItem{
				{
					Name:     "Item",
					Price:    100,
					Quantity: 1,
					Sum:      &sum,
					VAT: &SellRequestReceiptItemVAT{
						Type: "none",
						Sum:  &vatSum,
					},
				},
			},
			Payments: []SellRequestReceiptPayment{
				{
					Type: 1,
					Sum:  &sum,
				},
			},
			Total: 100,
		},
	})
	if err != nil {
		t.Fatalf("Sell returned error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Fatalf("unexpected method: %q", gotMethod)
	}

	if gotPath != "/group-code/sell" {
		t.Fatalf("unexpected path: %q", gotPath)
	}

	if gotToken != "client-token" {
		t.Fatalf("unexpected token query: %q", gotToken)
	}

	if gotBody.ExternalID != "order-1" {
		t.Fatalf("unexpected external id: %q", gotBody.ExternalID)
	}

	if len(gotBody.Receipt.Items) != 1 {
		t.Fatalf("unexpected items count: %d", len(gotBody.Receipt.Items))
	}

	if response.UUID != "uuid-1" || response.Status != "wait" {
		t.Fatalf("unexpected response: %+v", response)
	}
}

func TestSellCorrection(t *testing.T) {
	var gotPath string
	var gotToken string
	var gotBody SellCorrectionRequest

	client := NewClient("https://example.com", &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			gotPath = r.URL.Path
			gotToken = r.URL.Query().Get("token")

			if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
				t.Fatalf("decode request body: %v", err)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body: io.NopCloser(strings.NewReader(`{
				"uuid":"uuid-correction-1",
				"status":"wait",
				"timestamp":"2026-03-13T10:21:10+03:00",
				"error":null
			}`)),
			}, nil
		}),
	})
	client.Token = "client-token"

	sum := 100.0
	response, err := client.SellCorrection(context.Background(), "group-code", SellCorrectionRequest{
		Timestamp:  "2026-03-13T10:21:10+03:00",
		ExternalID: "correction-1",
		Correction: SellCorrectionRequestCorrection{
			Company: SellCorrectionRequestCorrectionCompany{
				SNO:            "osn",
				INN:            "5544332219",
				PaymentAddress: "https://example.com",
			},
			CorrectionInfo: SellCorrectionRequestCorrectionCorrectionInfo{
				Type:       "self",
				BaseDate:   "13.03.2026",
				BaseNumber: "1",
			},
			Payments: []SellCorrectionRequestCorrectionPayment{
				{Type: 1, Sum: &sum},
			},
			Vats: []SellCorrectionRequestCorrectionVat{
				{Type: "none", Sum: func() *float64 { zero := 0.0; return &zero }()},
			},
		},
	})
	if err != nil {
		t.Fatalf("SellCorrection returned error: %v", err)
	}

	if gotPath != "/group-code/sell_correction" {
		t.Fatalf("unexpected path: %q", gotPath)
	}

	if gotToken != "client-token" {
		t.Fatalf("unexpected token: %q", gotToken)
	}

	if gotBody.ExternalID != "correction-1" {
		t.Fatalf("unexpected external id: %q", gotBody.ExternalID)
	}

	if response.UUID != "uuid-correction-1" {
		t.Fatalf("unexpected uuid: %q", response.UUID)
	}
}

func TestReport(t *testing.T) {
	client := NewClient("https://example.com", &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			if r.Method != http.MethodGet {
				t.Fatalf("unexpected method: %q", r.Method)
			}

			if r.URL.Path != "/group-code/report/uuid-1" {
				t.Fatalf("unexpected path: %q", r.URL.Path)
			}

			if r.URL.Query().Get("token") != "client-token" {
				t.Fatalf("unexpected token query: %q", r.URL.Query().Get("token"))
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body: io.NopCloser(strings.NewReader(`{
				"uuid":"uuid-1",
				"status":"done",
				"timestamp":"2026-03-13T10:20:32+03:00",
				"error":null,
				"payload":{
					"total":100,
					"fns_site":"www.nalog.gov.ru",
					"fn_number":"1234567890123456",
					"shift_number":1,
					"receipt_datetime":"2026-03-13T10:20:31+03:00",
					"fiscal_receipt_number":10,
					"fiscal_document_number":20,
					"ecr_registration_number":"0000000000000001",
					"fiscal_document_attribute":30,
					"ofd_inn":"1234567890",
					"ofd_receipt_url":"https://example.com/receipt"
				}
			}`)),
			}, nil
		}),
	})
	client.Token = "client-token"

	response, err := client.Report(context.Background(), "group-code", "uuid-1")
	if err != nil {
		t.Fatalf("Report returned error: %v", err)
	}

	if response.Status != "done" {
		t.Fatalf("unexpected status: %q", response.Status)
	}

	if response.Payload == nil {
		t.Fatal("payload is nil")
	}

	if response.Payload.Total != 100 {
		t.Fatalf("unexpected total: %v", response.Payload.Total)
	}
}

func TestDoInvalidJSON(t *testing.T) {
	client := NewClient("https://example.com", &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body:       io.NopCloser(strings.NewReader(`{invalid json`)),
			}, nil
		}),
	})

	_, err := client.PostToken(context.Background(), PostTokenRequest{
		Login: "login",
		Pass:  "password",
	})
	if err == nil {
		t.Fatal("expected decode error")
	}
}

func TestDoTransportError(t *testing.T) {
	client := NewClient("https://example.com", &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return nil, errors.New("transport error")
		}),
	})

	_, err := client.PostToken(context.Background(), PostTokenRequest{
		Login: "login",
		Pass:  "password",
	})
	if err == nil {
		t.Fatal("expected transport error")
	}
}

func TestDoHTTPError(t *testing.T) {
	client := NewClient("https://example.com", &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusUnauthorized,
				Body:       io.NopCloser(strings.NewReader(`{"error":"unauthorized"}`)),
			}, nil
		}),
	})

	_, err := client.PostToken(context.Background(), PostTokenRequest{
		Login: "login",
		Pass:  "password",
	})
	if err == nil {
		t.Fatal("expected http error")
	}

	var responseErr *ErrorResponse
	if !errors.As(err, &responseErr) {
		t.Fatalf("expected ErrorResponse, got %T", err)
	}

	if responseErr.StatusCode != http.StatusUnauthorized {
		t.Fatalf("unexpected status code: %d", responseErr.StatusCode)
	}
}
