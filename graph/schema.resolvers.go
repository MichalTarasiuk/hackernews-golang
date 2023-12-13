package graph

import (
	"context"
	"fmt"

	"github.com/MichalTarasiuk/hackernews/graph/model"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := model.User{
		Name: "test",
	}
	link := model.Link{
		Address: input.Address,
		Title:   input.Title,
		User:    &user,
	}

	return &link, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var links []*model.Link
	dummyLink := model.Link{
		Title:   "our dummy link",
		Address: "https://address.org",
		User:    &model.User{Name: "admin"},
	}
	links = append(links, &dummyLink)
	return links, nil
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
