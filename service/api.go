package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fezonV/hero-pick-helper/internal/hero"
)

func GetHeroes() ([]hero.HeroModel, error) {
	resp, err := http.Get("https://api.opendota.com/api/heroes")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	var heroes []hero.HeroModel

	err = json.NewDecoder(resp.Body).Decode(&heroes)
	if err != nil {
		return nil, err
	}
	return heroes, nil
}

func GetHeroStats() ([])
