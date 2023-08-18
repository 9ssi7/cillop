package rpc

import (
	"context"
	"errors"

	"github.com/ssibrahimbas/cillop/app/command"
	"github.com/ssibrahimbas/cillop/app/query"
	"github.com/ssibrahimbas/cillop/server/rpc/routes"
)

func (s *srv) Create(ctx context.Context, req *routes.CreateProductRequest) (*routes.CreateProductResponse, error) {
	res, err := s.app.Commands.ProductCreate.Handle(ctx, command.ProductCreateCmd{
		Name: req.Name,
	})
	if err != nil {
		msg := s.i18n.TranslateFromError(*err, "en")
		return nil, errors.New(msg)
	}
	return &routes.CreateProductResponse{
		Id: res.ID,
	}, nil
}

func (s *srv) Get(ctx context.Context, req *routes.GetProductRequest) (*routes.GetProductResponse, error) {
	res, err := s.app.Queries.ProductGet.Handle(ctx, query.ProductGetQuery{
		ID: req.Id,
	})
	if err != nil {
		msg := s.i18n.TranslateFromError(*err, "en")
		return nil, errors.New(msg)
	}
	return &routes.GetProductResponse{
		Name: res.Name,
	}, nil
}
