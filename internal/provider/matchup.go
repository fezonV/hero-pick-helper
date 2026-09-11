package provider

import (
	"context"

	"github.com/fezonV/hero-pick-helper/internal/models"
)

type Provider interface {
	GetHeroMatchups(ctx context.Context, heroID int64) ([]models.Matchup, error)
}
