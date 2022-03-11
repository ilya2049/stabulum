package product

import (
	"context"
	"log"
	"stabulum/internal/domain/product"
	"time"
)

type usecasesRetry struct {
	next Usecases

	config UsecasesRetryConfig
}

func (u *usecasesRetry) Create(ctx context.Context, p product.Product) (err error) {
	attempt := 1

	for {
		if err = u.next.Create(ctx, p); err == nil {
			return nil
		}

		if attempt > u.config.MaxAttempt {
			return err
		}

		log.Println("usecase products create", "attempt", attempt, "error", err.Error())

		attempt++
		time.Sleep(u.config.RetryDelay)
	}
}
