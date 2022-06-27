package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/joud-cxu/hackernews/graph/generated"
	"github.com/joud-cxu/hackernews/graph/model"
	"github.com/joud-cxu/hackernews/internal/links"
	"github.com/joud-cxu/hackernews/internal/users"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	linkID := link.Save()
	return &model.Link{ID: strconv.FormatInt(linkID, 10), Title: link.Title, Address: link.Address}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	userID := user.Save()
	return &model.User{ID: strconv.FormatInt(userID, 10), Name: user.Username}, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks {
		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address})
	}
	return resultLinks, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var resultUsers []*model.User
	var dbUsers []users.User
	dbUsers = users.GetAll()
	for _, user := range dbUsers {
		resultUsers = append(resultUsers, &model.User{ID: user.ID, Name: user.Username})
	}
	return resultUsers, nil
}

func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {
	var resultUser *model.User
	var dbUser *users.User
	// TODO pass id from request context
	dbUser, _ = users.GetUserById("3")
	resultUser = &model.User{ID: dbUser.ID, Name: dbUser.Username}
	return resultUser, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
