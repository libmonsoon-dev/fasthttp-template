// +build e2e

package rest

import (
	"encoding/base64"
	"encoding/json"
	"github.com/libmonsoon-dev/fasthttp-template/app/e2e"
	"github.com/valyala/fasthttp"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	const (
		email    = "user@example.com"
		password = "password"
	)
	base64pass := base64.StdEncoding.EncodeToString([]byte(password))

	client, closeServer := e2e.TestApp()
	defer closeServer()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI(e2e.AppUrl + "/rest/auth/sign-up")
	req.SetBodyString(`{
		"email": "` + email + `",
		"password": "` + base64pass + `"
	}`)

	if err := client.Do(req, resp); err != nil {
		t.Fatal(err)
	}

	statusCode := resp.StatusCode()
	if statusCode != http.StatusOK {
		t.Fatalf("Expect status: %v, actual: %v", http.StatusOK, statusCode)
	}

	responseBody := resp.Body()

	var signUpResult map[string]interface{}
	if err := json.Unmarshal(responseBody, &signUpResult); err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", responseBody)
	t.Logf("%#v", signUpResult)
}
