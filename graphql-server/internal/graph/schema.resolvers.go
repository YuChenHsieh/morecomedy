package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-server/internal/graph/generated"
	"graphql-server/internal/graph/mock"
	"graphql-server/internal/graph/model"
	"graphql-server/pkg/db"
	"graphql-server/pkg/utils"
)

func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	var event *model.Event
	q := `SELECT * from Event`
	rows, err := db.Clickhouse.Query(q)
	if err != nil {
		return nil, err
	}
	utils.Use(rows)
	return event, nil
}

func (r *mutationResolver) CreateShow(ctx context.Context, input model.NewShow) (*model.Show, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	var events []*model.Event
	events = mock.Events
	// q := `SELECT * from Show`
	// rows, err := db.Clickhouse.Query(q)
	// if err != nil {
	// 	return nil, err
	// }

	// defer rows.Close()

	// // Loop through rows, using Scan to assign column data to struct fields.
	// for rows.Next() {
	// 	var event *model.Event
	// 	if err := rows.Scan(&event.ID, &event.Name, &event.Briefly, &event.Shows); err != nil {
	// 		return events, err
	// 	}
	// 	events = append(events, event)
	// }

	return events, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
