package main

import (
	"context"
	"fmt"

	"github.com/fezonV/hero-pick-helper/internal/provider"
	"github.com/fezonV/hero-pick-helper/service"
)

func PrintMatchups(ctx context.Context, p provider.Provider, heroID int64) error {
	matchups, err := p.GetHeroMatchups(ctx, heroID)
	if err != nil {
		return err
	}
	hero, err := service.GetHeroByID(ctx, heroID)
	if err != nil {
		return err
	}

	fmt.Println("Матчапы против: ", hero.Localized_name)
	mapOfHeroes, err := service.GetHeroesMap(ctx)
	if err != nil {
		return err
	}
	for _, matchup := range matchups {
		wr, err := matchup.Winrate()
		if err != nil {
			return err
		}
		fmt.Println(mapOfHeroes[matchup.HeroID].Localized_name, ": ", wr, "(", matchup.Wins, "/", matchup.GamesPlayed, ")")
	}
	return nil
}
