package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mixarchitecture/microp/server/http/result"
	"github.com/ssibrahimbas/cillop/app/command"
	"github.com/ssibrahimbas/cillop/app/query"
)

func (s *srv) ProductCreate(ctx *fiber.Ctx) error {
	cmd := &command.ProductCreateCmd{}
	s.parseBody(ctx, cmd)
	res, err := s.app.Commands.ProductCreate.Handle(ctx.UserContext(), *cmd)
	return result.IfSuccessDetail(err, ctx, *s.i18n, Messages.Ok, func() interface{} {
		return res
	})
}

func (s *srv) ProductGet(ctx *fiber.Ctx) error {
	q := &query.ProductGetQuery{}
	s.parseParams(ctx, q)
	res, err := s.app.Queries.ProductGet.Handle(ctx.UserContext(), *q)
	return result.IfSuccessDetail(err, ctx, *s.i18n, Messages.Ok, func() interface{} {
		return res
	})
}
