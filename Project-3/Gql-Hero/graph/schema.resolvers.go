package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gql-hero/graph/generated"
	"gql-hero/graph/model"
)

// CreateHero is the resolver for the createHero field.
func (r *mutationResolver) CreateHero(ctx context.Context, input model.HeroInput) (*model.Hero, error) {
	var hero model.Hero = r.heroService.CreateHero(input)

	return &hero, nil
}

// Heroes is the resolver for the heroes field.
func (r *queryResolver) Heroes(ctx context.Context) ([]*model.Hero, error) {
	var heroes []*model.Hero = r.heroService.GetAllHero()

	return heroes, nil
}

// Hero is the resolver for the hero field.
func (r *queryResolver) Hero(ctx context.Context, id string) (*model.Hero, error) {
	hero, err := r.heroService.GetHeroByID(id)

	if err != nil {
		return nil, err
	}
	return hero, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
