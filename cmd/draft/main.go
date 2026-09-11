package main

import (
	"context"
	"fmt"

	"github.com/fezonV/hero-pick-helper/internal/models"
	"github.com/fezonV/hero-pick-helper/internal/opendota"
	"github.com/fezonV/hero-pick-helper/service"
)

func main() {
	ctx := context.Background()
	//stats, err := service.GetHeroStats(ctx)
	var ID int64
	fmt.Print("Введите номер героя: ")
	_, err := fmt.Scan(&ID)
	if err != nil {
		panic(err)
	}
	openDotaProvider := opendota.NewProvider()
	matchups, err := openDotaProvider.GetHeroMatchups(ctx, ID)
	hero, err := service.GetHeroByID(ctx, ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Матчапы против ", hero.Localized_name)
	mapOfHeroes, err := service.GetHeroesMap(ctx)
	if err != nil {
		panic(err)
	}
	for _, v := range matchups {
		curhero, ok := mapOfHeroes[v.HeroID]
		if ok == false {
			panic(models.ErrHeroNotFound)
		}
		winrate, err := v.Winrate()
		if err != nil {
			fmt.Println("Нет игр")
		}
		fmt.Println(curhero.Localized_name, ": ", winrate, "(", v.Wins, "/", v.GamesPlayed, ")")

	}
	/*
		stat, ok := stats[ID]
		if !ok {
			fmt.Println("Герой с таким ID не найден")
			return
		}
		if err != nil {
			fmt.Println(err.Error())
		}

		wr, err := stat.Winrate(4)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Println("Имя:", stat.Localized_name, "Винрейт на титанах:", wr*100, "%")
	*/

}
