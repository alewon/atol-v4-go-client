# atol-v4-go-client

Go-клиент для ATOL Online v4.

Текущая реализация основана на локальной спецификации из [`doc.md`](./doc.md), которая соответствует документации ATOL версии 5.23.

## Установка

```bash
go get github.com/alewon/atol-v4-go-client
```

## Импорт

```go
import "github.com/alewon/atol-v4-go-client"
```

## Быстрый старт

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

## Доступные методы

- `PostToken`
- `GetToken`
- `Sell`
- `SellRefund`
- `Buy`
- `BuyRefund`
- `SellCorrection`
- `BuyCorrection`
- `Report`

## Вклад в проект

См. [`CONTRIBUTING.md`](./CONTRIBUTING.md).

## Безопасность

См. [`SECURITY.md`](./SECURITY.md).

## Лицензия

MIT. См. [`LICENSE`](./LICENSE).
