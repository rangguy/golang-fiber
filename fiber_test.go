package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
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
