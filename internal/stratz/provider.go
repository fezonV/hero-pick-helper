package stratz

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/fezonV/hero-pick-helper/internal/models"
)

type Provider struct {
	httpClient *http.Client
	baseURL    string
}

func newProvider() *Provider {
	return &Provider{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://api.stratz.com/graphql",
	}
}

func (p *Provider) GetHeroMatchups(ctx context.Context, heroID int64) ([]models.OpenDotaMatchup, error) {
	var request graphQLRequest
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", p.baseURL, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	
}
