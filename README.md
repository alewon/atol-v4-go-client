# atol-v4-go-client

Simple Go client for ATOL Online v4.

The package is intentionally straightforward:

- no request validation;
- no attempt to deduplicate request and response models across methods;
- request and response structs for each method are declared explicitly;
- public API is kept flat and predictable.

The current implementation is based on the local specification in [`doc.md`](./doc.md).

## Install

```bash
go get github.com/alewon/atol-v4-go-client
```

## Package

```go
import "github.com/alewon/atol-v4-go-client"
```

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	atol "github.com/alewon/atol-v4-go-client"
)

func main() {
	ctx := context.Background()

	client := atol.NewClient("", nil)

	tokenResponse, err := client.PostToken(ctx, atol.PostTokenRequest{
		Login: "your-login",
		Pass:  "your-password",
	})
	if err != nil {
		log.Fatal(err)
	}

	client.Token = tokenResponse.Token

	sum := 100.0
	vatSum := 0.0

	sellResponse, err := client.Sell(ctx, "group-code", atol.SellRequest{
		Timestamp:  "2026-03-13T10:20:30+03:00",
		ExternalID: "example-order-1",
		Receipt: atol.SellRequestReceipt{
			Client: atol.SellRequestReceiptClient{
				Email: "user@example.com",
			},
			Company: atol.SellRequestReceiptCompany{
				Email:          "shop@example.com",
				SNO:            "osn",
				INN:            "5544332219",
				PaymentAddress: "https://example.com",
			},
			Items: []atol.SellRequestReceiptItem{
				{
					Name:     "Test item",
					Price:    100,
					Quantity: 1,
					Sum:      &sum,
					VAT: &atol.SellRequestReceiptItemVAT{
						Type: "none",
						Sum:  &vatSum,
					},
				},
			},
			Payments: []atol.SellRequestReceiptPayment{
				{
					Type: 1,
					Sum:  &sum,
				},
			},
			Total: 100,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	reportResponse, err := client.Report(ctx, "group-code", sellResponse.UUID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reportResponse.Status)
}
```

## Available methods

- `PostToken`
- `GetToken`
- `Sell`
- `SellRefund`
- `Buy`
- `BuyRefund`
- `SellCorrection`
- `BuyCorrection`
- `Report`

## Project principles

- Keep code obvious.
- Prefer explicit types over shared abstractions.
- Do not add validation layer.
- Do not hide API details behind complex helpers.

## Stability

This package follows the local API description in [`doc.md`](./doc.md). If ATOL changes the API, the Go models should be updated explicitly.

## Contributing

See [`CONTRIBUTING.md`](./CONTRIBUTING.md).

## Security

See [`SECURITY.md`](./SECURITY.md).

## License

MIT. See [`LICENSE`](./LICENSE).
