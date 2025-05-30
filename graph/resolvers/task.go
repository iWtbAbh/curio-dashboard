package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"
	"fmt"
	"time"

	"github.com/strahe/curio-dashboard/graph"
	"github.com/strahe/curio-dashboard/graph/cachecontrol"
	"github.com/strahe/curio-dashboard/graph/model"
)

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id int) (*model.Task, error) {
	return r.loader.Task(ctx, id)
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	return r.loader.Tasks(ctx)
}

// TasksCount is the resolver for the tasksCount field.
func (r *queryResolver) TasksCount(ctx context.Context) (int, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	return r.loader.TasksCount(ctx)
}

// TaskHistories is the resolver for the taskHistories field.
func (r *queryResolver) TaskHistories(ctx context.Context, start *time.Time, end *time.Time, hostPort *string, name *string, result *bool, offset int, limit int) ([]*model.TaskHistory, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute*5)
	return r.loader.TaskHistories(ctx, start, end, hostPort, name, result, offset, limit)
}

// TaskHistoriesCount is the resolver for the taskHistoriesCount field.
func (r *queryResolver) TaskHistoriesCount(ctx context.Context, start *time.Time, end *time.Time, hostPort *string, name *string, result *bool) (int, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute*5)
	return r.loader.TaskHistoriesCount(ctx, start, end, hostPort, name, result)
}

// TaskHistoriesAggregate is the resolver for the taskHistoriesAggregate field.
func (r *queryResolver) TaskHistoriesAggregate(ctx context.Context, start time.Time, end time.Time, interval model.TaskHistoriesAggregateInterval) ([]*model.TaskAggregate, error) {
	if start.IsZero() || end.IsZero() {
		return nil, fmt.Errorf("start and end time must be provided")
	}
	if start.After(end) {
		return nil, fmt.Errorf("start time must be before end time")
	}
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute*5)
	return r.loader.TaskHistoriesAggregate(ctx, start, end, interval)
}

// TasksStats is the resolver for the tasksStats field.
func (r *queryResolver) TasksStats(ctx context.Context, start time.Time, end time.Time, machine *string) ([]*model.TaskStats, error) {
	stats, err := r.loader.TasksStats(ctx, start, end, machine)
	if err != nil {
		return nil, err
	}
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute*5)
	return stats, nil
}

// TaskNames is the resolver for the taskNames field.
func (r *queryResolver) TaskNames(ctx context.Context) ([]string, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Hour)
	return r.loader.TaskNames(ctx)
}

// TaskSuccessRate is the resolver for the taskSuccessRate field.
func (r *queryResolver) TaskSuccessRate(ctx context.Context, name *string, start time.Time, end time.Time) (*model.TaskSuccessRate, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute*10)
	return r.loader.TaskSuccessRate(ctx, name, start, end)
}

// RunningTaskSummary is the resolver for the runningTaskSummary field.
func (r *queryResolver) RunningTaskSummary(ctx context.Context) (*model.RunningTaskSummary, error) {
	var res model.RunningTaskSummary
	end := time.Now()
	start := end.Add(-time.Hour * 24)
	var err error
	res.Running, err = r.loader.TaskRunningCount(ctx)
	if err != nil {
		return nil, err
	}
	res.Queued, err = r.loader.TaskQueuedCount(ctx)
	if err != nil {
		return nil, err
	}
	avg, err := r.loader.TaskAvgWaitTime(ctx, start, end, nil)
	if err != nil {
		return nil, err
	}
	res.AverageWaitTime = avg.Seconds()
	return &res, nil
}

// TaskDurationStats is the resolver for the taskDurationStats field.
func (r *queryResolver) TaskDurationStats(ctx context.Context, name string, start time.Time, end time.Time) (*model.TaskDurationStats, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute*10)
	return r.loader.TaskDurationStats(ctx, name, start, end)
}

// TasksDurationStats is the resolver for the tasksDurationStats field.
func (r *queryResolver) TasksDurationStats(ctx context.Context, start time.Time, end time.Time) ([]*model.TaskDurationStats, error) {
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute*10)
	return r.loader.TasksDurationStats(ctx, start, end)
}

// CompletedTask is the resolver for the completedTask field.
func (r *subscriptionResolver) CompletedTask(ctx context.Context, machine *string, last int) (<-chan *model.TaskHistory, error) {
	return r.loader.SubCompletedTask(ctx, machine, last)
}

// NewTask is the resolver for the newTask field.
func (r *subscriptionResolver) NewTask(ctx context.Context, machineID *int, last int) (<-chan *model.Task, error) {
	return r.loader.SubNewTask(ctx, machineID, last)
}

// InitiatedBy is the resolver for the initiatedBy field.
func (r *taskResolver) InitiatedBy(ctx context.Context, obj *model.Task) (*model.Machine, error) {
	if obj.InitiatedByID == nil {
		return nil, nil
	}
	var out model.Machine
	if err := r.db.QueryRow(ctx, "SELECT id,last_contact,host_and_port,cpu,gpu,ram FROM harmony_machines WHERE id = $1", obj.InitiatedByID).
		Scan(&out.ID, &out.LastContact, &out.HostAndPort, &out.CPU, &out.Gpu, &out.RAM); err != nil {
		return nil, err
	}
	return &out, nil
}

// Owner is the resolver for the owner field.
func (r *taskResolver) Owner(ctx context.Context, obj *model.Task) (*model.Machine, error) {
	if obj.OwnerID == nil {
		return nil, nil
	}
	var out model.Machine
	if err := r.db.QueryRow(ctx, `SELECT
    id,
    last_contact,
    host_and_port,
    cpu,
    gpu,
    ram
FROM harmony_machines
WHERE id = $1`, obj.OwnerID).
		Scan(&out.ID, &out.LastContact, &out.HostAndPort, &out.CPU, &out.Gpu, &out.RAM); err != nil {
		return nil, err
	}
	return &out, nil
}

// AddedBy is the resolver for the addedBy field.
func (r *taskResolver) AddedBy(ctx context.Context, obj *model.Task) (*model.Machine, error) {
	var out model.Machine
	if err := r.db.QueryRow(ctx, `SELECT
    id,
    last_contact,
    host_and_port,
    cpu,
    gpu,
    ram
FROM harmony_machines
WHERE id = $1`, obj.AddedByID).
		Scan(&out.ID, &out.LastContact, &out.HostAndPort, &out.CPU, &out.Gpu, &out.RAM); err != nil {
		log.Errorf("failed to fetch machine %d: %s", obj.AddedByID, err) // This should never happen, unless there's a problem somewhere
		return &model.Machine{
			ID: obj.AddedByID,
		}, nil
	}
	return &out, nil
}

// PreviousTask is the resolver for the previousTask field.
func (r *taskResolver) PreviousTask(ctx context.Context, obj *model.Task) (*model.TaskHistory, error) {
	if obj.PreviousTaskID == nil {
		return nil, nil
	}
	var out model.TaskHistory
	if err := r.db.QueryRow(ctx, "SELECT id,task_id,name,posted,work_start,work_end, result, err FROM harmony_task_history WHERE task_id = $1", obj.PreviousTaskID).
		Scan(&out.ID, &out.TaskID, &out.Name, &out.Posted, &out.WorkStart, &out.WorkEnd, &out.Result, &out.Err); err != nil {
		return nil, err
	}
	return &out, nil
}

// Histories is the resolver for the histories field.
func (r *taskResolver) Histories(ctx context.Context, obj *model.Task) ([]*model.TaskHistory, error) {
	var out []*model.TaskHistory
	if err := r.db.Select(ctx, &out, `SELECT
    id,
    task_id,
    name,
    posted,
    work_start,
    work_end,
    result,
    err,
    completed_by_host_and_port
FROM
    harmony_task_history
WHERE task_id = $1`, obj.ID); err != nil {
		return nil, err
	}
	return out, nil
}

// CompletedBy is the resolver for the completedBy field.
func (r *taskHistoryResolver) CompletedBy(ctx context.Context, obj *model.TaskHistory) (*model.Machine, error) {
	res, err := r.loader.MachineByHostPort(ctx, obj.CompletedByHostAndPort)
	if err != nil {
		return nil, err
	}
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	return res, nil
}

// Task returns graph.TaskResolver implementation.
func (r *Resolver) Task() graph.TaskResolver { return &taskResolver{r} }

// TaskHistory returns graph.TaskHistoryResolver implementation.
func (r *Resolver) TaskHistory() graph.TaskHistoryResolver { return &taskHistoryResolver{r} }

type taskResolver struct{ *Resolver }
type taskHistoryResolver struct{ *Resolver }
