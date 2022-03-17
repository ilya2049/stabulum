package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ShouldBindJSON(c *gin.Context, obj interface{ Validate() error }) error {
	if err := c.ShouldBindJSON(&obj); err != nil {
		return fmt.Errorf("failed to bind json: %w", err)
	}

	if err := obj.Validate(); err != nil {
		return fmt.Errorf("failed to validate a view: %w", err)
	}

	return nil
}
