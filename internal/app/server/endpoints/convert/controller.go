package convert

import "context"

func NewController(ctx context.Context) *Controller {
	return &Controller{
		ctx: ctx,
	}
}

type Controller struct {
	ctx context.Context // root ctx
}
