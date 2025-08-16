package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var app = fiber.New()

func TestRoutingHelloWorld(t *testing.T) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	request := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "Hello World", string(bytes))
}

func TestCtx(t *testing.T) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")
		return c.SendString("Hello " + name)
	})

	request := httptest.NewRequest("GET", "/hello?name=Rangga", nil)
	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "Hello Rangga", string(bytes))

	request = httptest.NewRequest("GET", "/hello", nil)
	resp, err = app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	bytes, err = io.ReadAll(resp.Body)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "Hello Guest", string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(c *fiber.Ctx) error {
		first := c.Get("firstname")   // header
		last := c.Cookies("lastname") // cookies
		return c.SendString("Hello " + first + " " + last)
	})

	request := httptest.NewRequest("GET", "/request", nil)
	request.Header.Set("firstname", "Rangga")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Mahendra"})
	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "Hello Rangga Mahendra", string(bytes))
}

func TestRouteParameter(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		orderId := c.Params("orderId")
		return c.SendString("Get Order " + orderId + " From User " + userId)
	})

	request := httptest.NewRequest("GET", "/users/rangga/orders/2", nil)
	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "Get Order 2 From User rangga", string(bytes))
}

func TestFormRequest(t *testing.T) {
	app.Post("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return c.SendString("Hello " + name)
	})

	body := strings.NewReader("name=Rangga")
	request := httptest.NewRequest("POST", "/hello", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "Hello Rangga", string(bytes))
}
