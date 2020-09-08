// +build e2e

package rest

import (
	"encoding/base64"
	"encoding/json"
	"github.com/libmonsoon-dev/fasthttp-template/app/e2e"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	const (
		email    = "user@example.com"
		password = "password"
	)
	base64pass := base64.StdEncoding.EncodeToString([]byte(password))
	reqBody := []byte(`{
		"email": "` + email + `",
		"password": "` + base64pass + `"
	}`)

	client, stop := e2e.Init()
	defer stop()

	var signUpResult map[string]interface{}
	{
		statusCode, responseBody, err := client.POST(
			"/rest/auth/sign-up",
			reqBody,
		)

		if err != nil {
			t.Fatal(err)
		}

		if statusCode != http.StatusOK {
			t.Fatalf("Expect status: %v, actual: %v", http.StatusOK, statusCode)
		}

		if err := json.Unmarshal(responseBody, &signUpResult); err != nil {
			t.Fatal(err)
		}
	}

	var signInResult map[string]interface{}
	{
		statusCode, responseBody, err := client.POST(
			"/rest/auth/sign-in",
			reqBody,
		)

		if err != nil {
			t.Fatal(err)
		}

		if statusCode != http.StatusOK {
			t.Fatalf("Expect status: %v, actual: %v", http.StatusOK, statusCode)
		}

		if err := json.Unmarshal(responseBody, &signInResult); err != nil {
			t.Fatal(err)
		}
	}

	userIdFromSignUp, userIdFromSignIn := int(signUpResult["userId"].(float64)), int(signInResult["userId"].(float64))
	if userIdFromSignUp != userIdFromSignIn {
		t.Errorf("userIdFromSignUp != userIdFromSignIn (%v != %v)", userIdFromSignUp, userIdFromSignIn)
	}

}
