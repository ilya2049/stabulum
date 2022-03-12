package product

import (
	"context"
	"net/http"
	"stabulum/internal/common/logger"
	"stabulum/internal/common/testfixture"
	"stabulum/internal/common/testfixture/mocks"
	"stabulum/internal/domain/product"
	mockproduct "stabulum/internal/domain/product/mocks"
	productview "stabulum/internal/infrastructure/api/router/product"
	"stabulum/internal/infrastructure/config"
	"stabulum/internal/infrastructure/di"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductCreate(t *testing.T) {
	diContainer, _, _ := di.NewTestContainer(
		config.Config{
			ProductUsecases: config.ProductUsecasesConfig{
				Retry: config.ProductUsecasesRetryConfig{
					MaxAttempt: 10,
					RetryDelay: time.Millisecond,
				},
			},
		},
		mocks.Config{
			ConfigureProductRepository: func(r *mockproduct.Repository, logger logger.Logger) {
				const maxFailedAttempt = 3
				attempt := 1

				r.On("Add", mock.Anything, mock.Anything).
					Return(func(_ context.Context, p product.Product) error {
						if attempt >= maxFailedAttempt {
							logger.Println("product added in the mock repository:", p.String())
							attempt = 0

							return nil
						}

						attempt++

						return testfixture.ErrTestUnexpected
					})
			},
		},
	)

	testHTTPServer := diContainer.APIHTTPTestServer
	defer testHTTPServer.Close()

	httpClient := httpexpect.New(t, testHTTPServer.URL)

	httpClient.POST("/product").
		WithJSON(productview.ProductView{Name: "Sticker"}).
		Expect().
		Status(http.StatusCreated)

	assert.Equal(t, []string{
		"usecase products create attempt 1 error unexpected test error\n",
		"usecase products create attempt 2 error unexpected test error\n",
		"product added in the mock repository: Product{Name: Sticker}\n",
	}, diContainer.SpyLogger.Logs)
}
