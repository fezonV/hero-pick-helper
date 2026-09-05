package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fezonV/hero-pick-helper/internal/models"
)

func GetHeroes(ctx context.Context) ([]models.HeroModel, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.opendota.com/api/heroes", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	var heroes []models.HeroModel

	err = json.NewDecoder(resp.Body).Decode(&heroes)
	if err != nil {
		return nil, err
	}
	return heroes, nil
}

func GetHeroStats(ctx context.Context) (map[int64]models.Stats, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.opendota.com/api/heroStats", nil)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	stats := make([]models.Stats, 0)
	err = json.NewDecoder(resp.Body).Decode(&stats)

	if err != nil {
		return nil, err
	}

	statsMap := make(map[int64]models.Stats)

	for i := range stats {
		statsMap[stats[i].Id] = stats[i]
	}

	return statsMap, nil
}
