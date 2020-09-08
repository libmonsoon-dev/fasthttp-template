package e2e

import "github.com/valyala/fasthttp"

func NewClient(client *fasthttp.Client) Client {
	return Client{
		client,
	}
}

type Client struct {
	*fasthttp.Client
}

func (c Client) POST(path string, reqBody []byte) (statusCode int, respBody []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI(AppUrl + path)
	req.SetBody(reqBody)

	if err := c.Do(req, resp); err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.Body(), nil
}
