package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"

	"github.com/strahe/curio-dashboard/graph"
	"github.com/strahe/curio-dashboard/graph/model"
)

// ID is the resolver for the id field.
func (r *sectorMetaResolver) ID(ctx context.Context, obj *model.SectorMeta) (string, error) {
	return fmt.Sprintf("%d-%d", obj.SpID, obj.SectorNum), nil
}

// SectorMeta returns graph.SectorMetaResolver implementation.
func (r *Resolver) SectorMeta() graph.SectorMetaResolver { return &sectorMetaResolver{r} }

type sectorMetaResolver struct{ *Resolver }
