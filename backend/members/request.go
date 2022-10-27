package members

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ansel1/merry/v2"
)

func sendRequest[t any](ctx context.Context, client *Client, req *http.Request) (*result[t], error) {
	token, err := client.tokenProvider.GetToken(ctx, client.domain)
	if err != nil {
		return nil, merry.Wrap(err, merry.WithUserMessage("unable to retrieve members token"))
	}
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if 200 > res.StatusCode || res.StatusCode > 299 {
		return nil, merry.New("error occurred when trying to fetch data from members", merry.WithHTTPCode(res.StatusCode), merry.WithMessage(string(body)))
	}

	var data result[t]
	err = json.Unmarshal(body, &data)
	return &data, err
}

func get[t any](ctx context.Context, client *Client, endpoint string) (*t, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/%s", client.domain, endpoint), nil)
	if err != nil {
		return nil, err
	}

	res, err := sendRequest[t](ctx, client, req)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
