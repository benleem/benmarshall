package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/benleem/benmarshall/internal/templates"
	"github.com/benleem/benmarshall/internal/templates/pages"
	"github.com/labstack/echo/v4"
)

type ContactHandler struct {
	key string
}

func NewContactHandler(key string) *ContactHandler {
	return &ContactHandler{
		key,
	}
}

func (h *ContactHandler) Get(c echo.Context) error {
	page := pages.Contact()
	hxReq := c.Request().Header.Get("Hx-Request")
	if hxReq != "" {
		c.Response().Header().Set(echo.HeaderVary, "Hx-Request")
		return page.Render(context.Background(), c.Response().Writer)
	}
	return templates.Layout(page, "benmarshall - contact").Render(context.Background(), c.Response().Writer)
}

func (h *ContactHandler) Post(c echo.Context) error {
	client := &http.Client{}
	apiUrl := "https://api.web3forms.com/submit"
	data, err := c.FormParams()
	if err != nil {
		return err
	}
	data.Set("access_key", h.key)
	req, _ := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(data.Encode())) // URL-encoded payload
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("%v", resp.StatusCode)
		return nil
	}

	return nil
}
