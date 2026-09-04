package service

import (
	"net/http"

	"github.com/fezonV/hero-pick-helper/internal/hero"
)

func ParseOpenDota() ([]hero.HeroModel, error) {
	resp, err := http.Get("http://example.com")
}
