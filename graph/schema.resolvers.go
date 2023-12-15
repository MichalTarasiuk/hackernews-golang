package graph

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MichalTarasiuk/hackernews/graph/model"
	"github.com/MichalTarasiuk/hackernews/internal/auth"
	"github.com/MichalTarasiuk/hackernews/internal/links"
	"github.com/MichalTarasiuk/hackernews/internal/users"
	"github.com/MichalTarasiuk/hackernews/pkg/jwt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := users.User{
		Username: input.Username,
		Password: input.Password,
	}
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user := users.User{
		Username: input.Username,
		Password: input.Password,
	}
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
		graphqlUser := &model.User{
			ID:   link.User.ID,
			Name: link.User.Username,
		}
		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address, User: graphqlUser})
	}
	return resultLinks, nil
}

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}
	link := links.Link{
		Title:   input.Title,
		Address: input.Address,
		User:    user,
	}
	linkId := link.Save()
	grahpqlUser := &model.User{
		ID:   user.ID,
		Name: user.Username,
	}
	return &model.Link{ID: strconv.FormatInt(linkId, 10), Title: link.Title, Address: link.Address, User: grahpqlUser}, nil
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
