package main

import (
	"context"
	"fmt"

	"github.com/fezonV/hero-pick-helper/service"
)

func main() {
	ctx := context.Background()
	stats, err := service.GetHeroStats(ctx)

	var ID int64
	fmt.Print("Введите номер героя: ")
	fmt.Scan(&ID)

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
}
