package service

import (
	"errors"
	"fmt"
	"gql-hero/database"
	"gql-hero/graph/model"
	"math/rand"
)

type HeroService struct {}

func (h *HeroService) GetAllHero() []*model.Hero {
	var heroes []*model.Hero = []*model.Hero{}

	database.DB.Find(&heroes)

	return heroes
}

func (h *HeroService) GetHeroByID(id string) (*model.Hero, error) {
	var hero model.Hero

	result := database.DB.First(&hero, "id = ?", id)

	if result.RowsAffected == 0 {
		return &model.Hero{}, errors.New("Hero not found")
	}
	return &hero, nil
}

func (s *HeroService) CreateHero(input model.HeroInput) model.Hero {
	var newHero model.Hero = model.Hero{
		ID: fmt.Sprintf("%d", rand.Int()),
		NamaHero: input.NamaHero,
		Role: input.Role,
		Emblem: input.Emblem,
		BattleSpell: input.BattleSpell,
	}

	database.DB.Create(&newHero)

	return newHero
}