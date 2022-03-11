package product

import (
	"context"
	"log"
	"net/http"
	"stabulum/internal/domain/product"
	mockproduct "stabulum/internal/domain/product/mocks"
	productview "stabulum/internal/infrastructure/api/router/product"
	"stabulum/internal/infrastructure/config"
	"stabulum/internal/infrastructure/di"
	"stabulum/internal/testfixture"
	"stabulum/internal/testfixture/mocks"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/mock"
)

func TestProductCreate(t *testing.T) {
	diContainer := di.NewTestContainer(
		config.ReadFromMemory(),
		mocks.Config{
			ConfigureProductRepository: func(r *mockproduct.Repository) {
				const maxFailedAttempt = 5
				attempt := 1

				r.On("Add", mock.Anything, mock.Anything).
					Return(func(_ context.Context, p product.Product) error {
						if attempt >= maxFailedAttempt {
							log.Println("product added in the mock repository:", p.String())
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
}
