package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"ex1/db"
	"ex1/entity"
	"ex1/graph/generated"
	"ex1/graph/model"
	"fmt"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := new(model.Todo)
	todo.ID = "1"
	todo.Text = input.Text
	todo.User = new(model.User)
	todo.User.ID = input.UserID
	return todo, nil
}

// CreateCompany is the resolver for the createCompany field.
func (r *mutationResolver) CreateCompany(ctx context.Context, input model.NewCompany) (*entity.Company, error) {
	bytes, _ := json.Marshal(input)
	company := &entity.Company{}
	json.Unmarshal(bytes, &company)
	db.GetDB().Save(company)
	return company, nil
}

// CreateEmployee is the resolver for the createEmployee field.
func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.NewEmployee) (*entity.Employee, error) {
	bytes, _ := json.Marshal(input)
	employee := &entity.Employee{}
	json.Unmarshal(bytes, &employee)
	db.GetDB().Save(employee)
	return employee, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos := make([]*model.Todo, 2)
	todo := new(model.Todo)
	todo.ID = "1"
	todo.Text = "test"
	todos[0] = todo
	todos[1] = todo
	return todos, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users := make([]*model.User, 1)
	user := new(model.User)
	user.ID = "1"
	user.Name = "Vlad"
	users[0] = user
	return users, nil
}

// Company is the resolver for the company field.
func (r *queryResolver) Company(ctx context.Context) (*entity.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
