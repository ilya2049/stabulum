package product

import (
	"context"
	"stabulum/internal/common/event"
)

type Repository interface {
	Atomic(context.Context,
		func(context.Context, Repository, event.Publisher) error,
	) error

	Add(context.Context, Product) error
}
