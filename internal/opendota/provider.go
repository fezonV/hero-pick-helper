package opendota

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fezonV/hero-pick-helper/internal/models"
)

type Provider struct {
	client  *http.Client
	baseURL string
}

func NewProvider() *Provider {
	return &Provider{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://api.opendota.com/api",
	}
}

func (p *Provider) GetHeroMatchups(ctx context.Context, heroID int64) ([]models.OpenDotaMatchup, error) {

	url := fmt.Sprintf("%s/heroes/%d/matchups", p.baseURL, heroID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	matchupArr := make([]models.OpenDotaMatchup, 0)
	err = json.NewDecoder(resp.Body).Decode(&matchupArr)

	if err != nil {
		return nil, err
	}
	return matchupArr, nil
}
