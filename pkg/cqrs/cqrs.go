package cqrs

import (
	"context"

	"github.com/mixarchitecture/i18np"
)

type Handler[TParams any, TResponse any] interface {
	Handle(context.Context, TParams) (TResponse, *i18np.Error)
}
