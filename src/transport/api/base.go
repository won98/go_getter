package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (a *ApiLauncher) SetUp() APIMethod {
	return a
}

func (a *ApiContext) RenderTemplate(template string, data map[string]interface{}) error {
	return a.Fiber.Render(template, data)
}

func (a *ApiContext) Render(data map[string]interface{}) error {
	return a.Fiber.Render(a.Name, data)
}

func (a *ApiContext) Reply(data interface{}) error {
	return a.Fiber.JSON(data)
}

func (a *ApiContext) String(str string) error {
	return a.Fiber.SendString(str)
}

func (a *ApiContext) RedirectWithAlert(url, message string) error {
	return a.Fiber.Render("redirect_message", map[string]interface{}{"redirect_url": url, "message": message})
}

func (a *ApiContext) Redirect(url string) error {
	return a.Fiber.Render("redirect", map[string]interface{}{"redirect_url": url})
}

func (a *ApiContext) SendFile(url string) error {
	return a.Fiber.SendFile(url)
}

func (a *ApiContext) GetTokenString() string {
	return a.Fiber.Get("Authorization")
}
func (a *ApiContext) GetReTokenString() string {
	return a.Fiber.Get("RefreshAuthorization")
}

func (a *ApiContext) bind(data interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("error: %v", r)
		}
	}()
	return a.Fiber.BodyParser(data)
}

func (a *ApiLauncher) baseMethod(path, name string, f func(ApiContext) error) func(ctx *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ac := ApiContext{
			Fiber: c,
			Name:  name,
			Path:  path,
			Env:   a.Launcher.Env,
		}
		return f(ac)
	}
}

func (a *ApiLauncher) handle(method, path string, handler func(ApiContext) error, name ...string) {
	routeName := path
	if len(name) > 0 {
		routeName = name[0]
	}

	wrappedHandler := a.baseMethod(path, routeName, handler)

	switch method {
	case fiber.MethodGet:
		a.Fiber.Get(path, wrappedHandler).Name(routeName)
	case fiber.MethodPost:
		a.Fiber.Post(path, wrappedHandler).Name(routeName)
	case fiber.MethodDelete:
		a.Fiber.Delete(path, wrappedHandler).Name(routeName)
	case fiber.MethodPut:
		a.Fiber.Put(path, wrappedHandler).Name(routeName)
	}
}

func (a *ApiLauncher) GET(path string, f func(ApiContext) error, name ...string) {
	a.handle(fiber.MethodGet, path, f, name...)
}

func (a *ApiLauncher) POST(path string, f func(ApiContext) error, name ...string) {
	a.handle(fiber.MethodPost, path, f, name...)
}

func (a *ApiLauncher) DELETE(path string, f func(ApiContext) error, name ...string) {
	a.handle(fiber.MethodDelete, path, f, name...)
}

func (a *ApiLauncher) PUT(path string, f func(ApiContext) error, name ...string) {
	a.handle(fiber.MethodPut, path, f, name...)
}
