package bacon

import (
	"context"
	"io"
	"net/http"
	"pie-fire-dire/internal/core/port"
)

var _ port.BaconIpsumAPI = (*Adapter)(nil)

type Dependencies struct {
	Client *http.Client
	URL    string
}

type Adapter struct {
	client *http.Client
	url    string
}

func New(deps Dependencies) *Adapter {
	return &Adapter{
		client: deps.Client,
		url:    deps.URL,
	}
}

func (a *Adapter) GetText(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.url, nil)
	if err != nil {
		return "", err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}
