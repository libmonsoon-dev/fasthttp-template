// +build e2e

package rest

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/libmonsoon-dev/fasthttp-template/app/e2e"
	"github.com/libmonsoon-dev/fasthttp-template/app/interface/http/rest/dto"
	"net/http"
	"testing"
)

func TestSignUp(t *testing.T) {
	const (
		email    = "user@example.com"
		password = "password"
	)
	base64pass := base64.StdEncoding.EncodeToString([]byte(password))
	reqBody := []byte(`{
		"email": "` + email + `",
		"password": "` + base64pass + `"
	}`)

	app, client, stop := e2e.Init()
	defer stop()

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

	var signUpResult dto.AuthToken
	if err := json.Unmarshal(responseBody, &signUpResult); err != nil {
		t.Fatal(err)
	}

	if _, err := app.AuthService.DecodeAuthToken(signUpResult.Token); err != nil {
		t.Fatal(err)
	}

}

func TestSignIn(t *testing.T) {
	const (
		email    = "user@example.com"
		password = "password"
	)
	base64pass := base64.StdEncoding.EncodeToString([]byte(password))
	reqBody := []byte(`{
		"email": "` + email + `",
		"password": "` + base64pass + `"
	}`)

	app, client, stop := e2e.Init()
	defer stop()

	ctx := context.Background()
	userId, err := app.UserService.Create(ctx, email, []byte(password))

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

	var signInResult dto.AuthToken
	if err := json.Unmarshal(responseBody, &signInResult); err != nil {
		t.Fatal(err)
	}

	userIdFromSignIn, err := app.AuthService.DecodeAuthToken(signInResult.Token)
	if err != nil {
		t.Fatal(err)
	}

	if userId != userIdFromSignIn {
		t.Fatalf("userIdFromSignUp != userIdFromSignIn (%v != %v)", userId, userIdFromSignIn)
	}

}
