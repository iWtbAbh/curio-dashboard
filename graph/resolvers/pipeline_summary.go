package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"

	"github.com/strahe/curio-dashboard/graph"
	"github.com/strahe/curio-dashboard/graph/model"
)

// Sdr is the resolver for the sdr field.
func (r *pipelineSummaryResolver) Sdr(ctx context.Context, obj *model.PipelineSummary) (int, error) {
	var out int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM sectors_sdr_pipeline WHERE sp_id = $1 AND after_sdr = false", obj.ID).Scan(&out)
	return out, err
}

// Trees is the resolver for the trees field.
func (r *pipelineSummaryResolver) Trees(ctx context.Context, obj *model.PipelineSummary) (int, error) {
	var out int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM sectors_sdr_pipeline WHERE sp_id = $1 AND (after_tree_d = false OR after_tree_c = false OR after_tree_r = false) AND after_sdr = true", obj.ID).Scan(&out)
	return out, err
}

// PrecommitMsg is the resolver for the precommitMsg field.
func (r *pipelineSummaryResolver) PrecommitMsg(ctx context.Context, obj *model.PipelineSummary) (int, error) {
	var out int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM sectors_sdr_pipeline WHERE sp_id = $1 AND after_tree_r = true and after_precommit_msg = false", obj.ID).Scan(&out)
	return out, err
}

// WaitSeed is the resolver for the waitSeed field.
func (r *pipelineSummaryResolver) WaitSeed(ctx context.Context, obj *model.PipelineSummary) (int, error) {
	head, err := r.fullNode.ChainHead(ctx)
	if err != nil {
		return 0, err
	}
	var out int
	err = r.db.QueryRow(ctx, "SELECT COUNT(*) FROM sectors_sdr_pipeline WHERE sp_id = $1 AND after_precommit_msg_success = true AND seed_epoch > $2", obj.ID, head.Height()).Scan(&out)
	return out, err
}

// Porep is the resolver for the porep field.
func (r *pipelineSummaryResolver) Porep(ctx context.Context, obj *model.PipelineSummary) (int, error) {
	var out int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM sectors_sdr_pipeline WHERE sp_id = $1 AND after_porep = false AND after_precommit_msg_success = true", obj.ID).Scan(&out)
	return out, err
}

// CommitMsg is the resolver for the commitMsg field.
func (r *pipelineSummaryResolver) CommitMsg(ctx context.Context, obj *model.PipelineSummary) (int, error) {
	var out int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM sectors_sdr_pipeline WHERE sp_id = $1 AND after_commit_msg_success = false AND after_porep = true", obj.ID).Scan(&out)
	return out, err
}

// Done is the resolver for the done field.
func (r *pipelineSummaryResolver) Done(ctx context.Context, obj *model.PipelineSummary) (int, error) {
	var out int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM sectors_sdr_pipeline WHERE sp_id = $1 AND after_commit_msg_success = true", obj.ID).Scan(&out)
	return out, err
}

// Failed is the resolver for the failed field.
func (r *pipelineSummaryResolver) Failed(ctx context.Context, obj *model.PipelineSummary) (int, error) {
	var out int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM sectors_sdr_pipeline WHERE sp_id = $1 AND failed = true", obj.ID).Scan(&out)
	return out, err
}

// PipelineSummary returns graph.PipelineSummaryResolver implementation.
func (r *Resolver) PipelineSummary() graph.PipelineSummaryResolver {
	return &pipelineSummaryResolver{r}
}

type pipelineSummaryResolver struct{ *Resolver }
