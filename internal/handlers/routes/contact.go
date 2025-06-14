package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/benleem/benmarshall/internal/templates/components"
)

type Web3ErrorResponse struct {
	Success bool `json:"success"`
	// Data    struct {
	// 	AccessKey string `json:"access_key"`
	// 	Email     string `json:"email"`
	// 	Message   string `json:"message"`
	// 	Name      string `json:"name"`
	// } `json:"data"`
	Message string `json:"message"`
}

type ContactHandler struct {
	key string
}

func NewContactHandler(key string) *ContactHandler {
	return &ContactHandler{
		key,
	}
}

func (h *ContactHandler) Post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// http.Error(w, fmt.Sprintf("error parsing form: %s", err.Error()), http.StatusInternalServerError)
		fmt.Println(fmt.Sprintf("error parsing form: %s", err.Error()))
		components.ContactStatus("failed").Render(r.Context(), w)
	}

	client := &http.Client{}
	apiUrl := "https://api.web3forms.com/submit"
	r.Form.Set("access_key", h.key)
	req, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(r.Form.Encode())) // URL-encoded payload
	if err != nil {
		// http.Error(w, fmt.Sprintf("error sending email: %s", err.Error()), http.StatusInternalServerError)
		fmt.Println(fmt.Sprintf("error sending email: %s", err.Error()))
		components.ContactStatus("failed").Render(r.Context(), w)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		// http.Error(w, fmt.Sprintf("error sending email: %s", err.Error()), http.StatusInternalServerError)
		fmt.Println(fmt.Sprintf("error sending email: %s", err.Error()))
		components.ContactStatus("failed").Render(r.Context(), w)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	// fmt.Println(string(bodyBytes))
	if err != nil {
		// http.Error(w, fmt.Sprintf("error reading web3 response body: %s", err.Error()), http.StatusInternalServerError)
		fmt.Println(fmt.Sprintf("error reading web3 response body: %s", err.Error()))
		components.ContactStatus("failed").Render(r.Context(), w)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// fmt.Printf("%v", resp.StatusCode)
		components.ContactStatus("success").Render(r.Context(), w)
	} else {
		var web3Response Web3ErrorResponse
		err = json.Unmarshal(bodyBytes, &web3Response)
		if err != nil {
			// http.Error(w, fmt.Sprintf("error unmarshalling json: %s", err.Error()), http.StatusInternalServerError)
			fmt.Println(fmt.Sprintf("error unmarshalling json: %s", err.Error()))
			components.ContactStatus("failed").Render(r.Context(), w)
		}
		// http.Error(w, fmt.Sprintf("error sending email: %s", web3Response.Message), http.StatusInternalServerError)
		fmt.Println(fmt.Sprintf("error sending email: %s", web3Response.Message))
		if strings.Contains(web3Response.Message, "spam") {
			components.ContactStatus("failed: spam detected").Render(r.Context(), w)
		} else if strings.Contains(web3Response.Message, "rate limited") {
			components.ContactStatus("failed: rate limited").Render(r.Context(), w)
		} else {

		}

		// case strings.Contains(web3Response.Message, "spam"):
		// 	components.ContactStatus("failed: spam detected").Render(r.Context(), w)
		// case "Rate limited":
		// 	components.ContactStatus("failed: rate limited").Render(r.Context(), w)
		// default:
		// 	components.ContactStatus("failed").Render(r.Context(), w)
		// }
	}
}
