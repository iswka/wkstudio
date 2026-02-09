package commonutils

// Package commonutils provides shared utilities for microservices

import (
	"context"
	"encoding/json"
	"errors"
)

// GetUserIDFromContext extracts user ID from context
// This function handles different types that JWT middleware might inject
func GetUserIDFromContext(ctx context.Context) (int64, error) {
	userIDValue := ctx.Value("user_id")
	if userIDValue == nil {
		return 0, errors.New("unauthorized")
	}

	var userID int64
	switch v := userIDValue.(type) {
	case float64:
		userID = int64(v)
	case json.Number:
		var err error
		userID, err = v.Int64()
		if err != nil {
			return 0, errors.New("invalid user ID")
		}
	case int64:
		userID = v
	case int:
		userID = int64(v)
	case int32:
		userID = int64(v)
	default:
		return 0, errors.New("invalid user ID")
	}

	return userID, nil
}
