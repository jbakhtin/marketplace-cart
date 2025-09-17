package custom_context

import (
	"context"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/domain"
)

var UserIDKey = "user_id"

// GetUserID безопасно извлекает userID из контекста
func GetUserID(ctx context.Context) (domain.UserID, bool) {
	userIDValue := ctx.Value(UserIDKey)
	if userIDValue == nil {
		return 0, false
	}
	
	userID, ok := userIDValue.(domain.UserID)
	return userID, ok
}

// SetUserID устанавливает userID в контекст
func SetUserID(ctx context.Context, userID domain.UserID) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}
