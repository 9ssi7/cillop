package event_stream

import (
	"context"
	"encoding/json"

	"github.com/ssibrahimbas/cillop/app/command"
)

func (s *srv) CreateProduct(data []byte) {
	dto := &command.ProductCreateCmd{}
	err := json.Unmarshal(data, &dto)
	if err != nil {
		return
	}
	_, _ = s.app.Commands.ProductCreate.Handle(context.Background(), *dto)
}
