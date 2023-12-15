package graph

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MichalTarasiuk/hackernews/graph/model"
	"github.com/MichalTarasiuk/hackernews/internal/links"
	"github.com/MichalTarasiuk/hackernews/internal/users"
	"github.com/MichalTarasiuk/hackernews/pkg/jwt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	linkID := link.Save()
	return &model.Link{ID: strconv.FormatInt(linkID, 10), Title: link.Title, Address: link.Address}, nil

}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	for _, link := range links.GetAll() {
		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address})
	}
	return resultLinks, nil
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
