package appcontext

import (
	"context"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

const (
	//Add here for another context fields
	AccountID = "account_id"
)

var (
	CtxContents = []string{middleware.RequestIDHeader, AccountID}
)

func SetCtxVal(ctx context.Context, key string, val interface{}) context.Context {
	return context.WithValue(ctx, key, val)
}

func GetCtxContent(ctx context.Context) map[string]interface{} {
	if ctx == nil {
		return nil
	}
	ctxVal := make(map[string]interface{})
	for _, val := range CtxContents {
		if v := ctx.Value(val); v != nil {
			ctxVal[val] = v
		}
	}
	return ctxVal
}

func GetCtxAccountID(ctx context.Context) uuid.UUID {
	if ctx == nil {
		return uuid.Nil
	}
	if v := ctx.Value(AccountID); v != nil {
		id, _ := uuid.Parse(v.(string))
		return id
	}
	return uuid.Nil
}
