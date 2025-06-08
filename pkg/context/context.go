package appcontext

import (
	"context"

	"github.com/go-chi/chi/v5/middleware"
)

const (
	RequestID = "request_id"
	AccountID = "account_id"
)

var (
	CtxContents = []string{AccountID}
)

func GetCtxContent(ctx context.Context) map[string]interface{} {
	if ctx == nil {
		return nil
	}

	ctxVal := make(map[string]interface{})
	requestID := middleware.GetReqID(ctx)
	if requestID != "" {
		ctxVal[RequestID] = requestID
	}

	for _, val := range CtxContents {
		if v := ctx.Value(val); v != nil {
			ctxVal[val] = v
		}
	}
	return ctxVal
}

func SetAccountID(ctx context.Context, accountID int64) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, AccountID, accountID)
}
