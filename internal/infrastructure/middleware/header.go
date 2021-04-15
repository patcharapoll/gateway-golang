package middleware

import (
	"context"
	"gateway-golang/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Header
const (
	userID        = "user-id"
	UUID          = "UUID"
	language      = "lang"
	Authorization = "Authorization"
)

// ContextKey
const (
	userIDContextKey        = "userIDContextKey"
	uuidContextKey          = "uuidContextKey"
	languageContextKey      = "languageContextKey"
	authorizationContextKey = "authorizationContextKey"
)

func UserIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader(userID)
		if userID == "" {
			c.String(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		c.Set(userIDContextKey, userID)
		return
	}
}

func HeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(uuidContextKey, c.GetHeader(UUID))
		c.Set(languageContextKey, c.GetHeader(language))
		c.Set(authorizationContextKey, c.GetHeader(Authorization))
		return
	}
}

func ForUUIDContext(ctx context.Context) string {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return ""
	}
	ctxData, _ := gc.Get(uuidContextKey)
	d, ok := ctxData.(string)
	if !ok {
		return ""
	}
	return d
}

func ForUserIDContext(ctx context.Context) string {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return ""
	}
	ctxData, _ := gc.Get(userIDContextKey)
	d, ok := ctxData.(string)
	if !ok {
		return ""
	}
	return d
}

func ForAuthorizationContext(ctx context.Context) string {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return ""
	}

	ctxData, _ := gc.Get(authorizationContextKey)
	d, ok := ctxData.(string)
	if !ok {
		return ""
	}

	return d
}

// ForLanguageContext ...
func ForLanguageContext(ctx context.Context) string {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return ""
	}
	ctxData, _ := gc.Get(languageContextKey)
	d, ok := ctxData.(string)
	if !ok {
		return ""
	}
	return d
}
