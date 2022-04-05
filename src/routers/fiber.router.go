package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberRouter struct {
	*fiber.App
}

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{})
	Params(string) string
}

func NewFiberRouter() *FiberRouter {
	r := fiber.New()

	r.Use(cors.New())
	r.Use(logger.New())

	return &FiberRouter{r}
}

func (r *FiberRouter) POST(path string, handler func(Context)) {
	r.App.Post(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PUT(path string, handler func(Context)) {
	r.App.Put(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}
func (r *FiberRouter) GET(path string, handler func(Context)) {
	r.App.Get(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) DELETE(path string, handler func(Context)) {
	r.App.Delete(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PATCH(path string, handler func(Context)) {
	r.App.Patch(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

type FiberCtx struct {
	*fiber.Ctx
}

func NewFiberCtx(c *fiber.Ctx) *FiberCtx {
	return &FiberCtx{Ctx: c}
}

func (c *FiberCtx) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}
func (c *FiberCtx) JSON(statuscode int, v interface{}) {
	c.Ctx.Status(statuscode).JSON(v)
}

func (c *FiberCtx) Params(key string) string {
	return c.Ctx.Params(key)
}
