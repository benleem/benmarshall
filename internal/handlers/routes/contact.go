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
	api_key string
}

func NewContactHandler(api_key string) *ContactHandler {
	return &ContactHandler{
		api_key,
	}
}

func (h *ContactHandler) Get(c echo.Context) error {
	page := pages.Contact()
	hxReq := c.Request().Header.Get("hx-request")
	if hxReq != "" {
		return page.Render(context.Background(), c.Response().Writer)
	}
	return templates.Layout(page, "benmarshall").Render(context.Background(), c.Response().Writer)
	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}

func (h *ContactHandler) Post(c echo.Context) error {
	client := &http.Client{}
	apiUrl := "https://api.web3forms.com/submit"
	data, err := c.FormParams()
	if err != nil {
		return err
	}
	data.Set("access_key", h.api_key)
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

	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}
