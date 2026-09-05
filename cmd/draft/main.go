package main

import (
	"context"
	"fmt"

	"github.com/fezonV/hero-pick-helper/service"
)

func main() {
	ctx := context.Background()
	//stats, err := service.GetHeroStats(ctx)
	var ID int64
	fmt.Print("Введите номер героя: ")
	fmt.Scan(&ID)
	matchups, err := service.GetHeroMatchups(ctx, ID)
	hero, err := service.GetHeroByID(ctx, ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Матчапы против ", hero.Localized_name)

	for _, v := range matchups {
		curhero, err := service.GetHeroByID(ctx, v.HeroID)
		if err != nil {
			panic(err)
		}
		winrate, err := v.Winrate()
		fmt.Println(curhero, ": ", winrate)

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
