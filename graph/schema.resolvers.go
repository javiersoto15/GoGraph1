package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/temp/GQL-Tut/graph/generated"
	"github.com/temp/GQL-Tut/graph/model"
)

func (r *mutationResolver) RegisterNewNinja(ctx context.Context, input *model.NewNinja) (*model.ResponseMessage, error) {
	//If there is a database, here you would insert the new data
	//as both models are similar, we type cast our input into the Ninja model
	newNinja := model.Ninja(*input)

	//we append into our array
	r.Ninjas = append(r.Ninjas, &newNinja)

	return &model.ResponseMessage{Message: "Ninja added"}, nil
}

func (r *queryResolver) FindGenin(ctx context.Context, name string) (*model.Ninja, error) {
	//This simulates a Find in a database
	for _, tempNinja := range r.Ninjas {

		//we look for ninjas that match the status of genin and the name
		if *tempNinja.Name == name && *tempNinja.Rank == "Genin" {
			return tempNinja, nil
		}
	}
	return nil, nil
}

func (r *queryResolver) ReturnAllHokages(ctx context.Context) ([]*model.Ninja, error) {
	hokages := []*model.Ninja{}
	for _, tempNinja := range r.Ninjas {
		//we look for ninjas that match the status of genin and the name
		if *tempNinja.Rank == "Hokage" {
			hokages = append(hokages, tempNinja)
		}
	}
	return hokages, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
