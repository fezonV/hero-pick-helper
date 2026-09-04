package main

import (
	"fmt"

	"github.com/fezonV/hero-pick-helper/service"
)

func main() {
	heroes, err := service.GetHeroes()
	if err != nil {
		fmt.Println(err.Error())
	}

	for i, hero := range heroes {
		fmt.Println(i, hero.Localized_name)
	}
}
