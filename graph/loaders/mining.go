package loaders

import (
	"context"
	"fmt"
	"time"

	"github.com/strahe/curio-dashboard/types"

	"github.com/strahe/curio-dashboard/graph/model"
)

type MiningLoader interface {
	MiningSummaryByDay(ctx context.Context, start, end time.Time) ([]*model.MiningSummaryDay, error)
	MiningCount(ctx context.Context, start, end time.Time, actor *types.ActorID) (*model.MiningCount, error)
	MiningTasks(ctx context.Context, start time.Time, end time.Time, actor *types.ActorID, won *bool) ([]*model.MiningTask, error)
}

func (l *Loader) MiningSummaryByDay(ctx context.Context, start, end time.Time) ([]*model.MiningSummaryDay, error) {
	if end.IsZero() {
		end = time.Now()
	}
	if end.Before(start) {
		return nil, fmt.Errorf("end time is before start time")
	}

	var result []*model.MiningSummaryDay
	err := l.db.Select(ctx, &result,
		`SELECT
    DATE_TRUNC('day', base_compute_time) AS day,
    sp_id as miner,
    COUNT(*) AS won_block
FROM
    mining_tasks
WHERE
    won = true
  AND base_compute_time BETWEEN $1 AND $2
GROUP BY
    day, sp_id
ORDER BY
    day, sp_id;`, start, end)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (l *Loader) MiningCount(ctx context.Context, start, end time.Time, actor *types.ActorID) (*model.MiningCount, error) {
	if end.IsZero() {
		end = time.Now()
	}
	if end.Before(start) {
		return nil, fmt.Errorf("end time is before start time")
	}
	type mm struct {
		Included bool `db:"included"`
		Count    int  `db:"count"`
	}
	var res []mm

	err := l.db.Select(ctx, &res, `
SELECT
    included,
    COUNT(*) AS count
FROM
    mining_tasks
WHERE
    won = true AND
    ($1::int IS NULL OR sp_id = $1) AND
    base_compute_time BETWEEN $2 AND $3
GROUP BY
    included;`, actor, start, end)
	if err != nil {
		return nil, err
	}
	result := &model.MiningCount{}

	for _, r := range res {
		if r.Included {
			result.Include = r.Count
		} else {
			result.Exclude = r.Count
		}
	}
	return result, nil
}

func (l *Loader) MiningTasks(ctx context.Context, actor *types.ActorID, won *bool, offset int, limit int) ([]*model.MiningTask, error) {
	var result []*model.MiningTask

	err := l.db.Select(ctx, &result, `
SELECT
    task_id,
    sp_id,
    epoch,
    base_compute_time,
    won,
    mined_cid,
    mined_header,
    mined_at,
    submitted_at,
    included
FROM
    mining_tasks
WHERE
    ($1::bool IS NULL OR won = $1) AND
    ($2::int IS NULL OR sp_id = $2)
ORDER BY
    base_compute_time DESC 
LIMIT $3 OFFSET $4;`, won, actor, limit, offset)

	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return result, nil
}
