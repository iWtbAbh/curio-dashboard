package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"
	"net/http"
	"time"

	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/strahe/curio-dashboard/graph"
	"github.com/strahe/curio-dashboard/graph/cachecontrol"
	"github.com/strahe/curio-dashboard/graph/model"
)

// Detail is the resolver for the detail field.
func (r *machineResolver) Detail(ctx context.Context, obj *model.Machine) (*model.MachineDetail, error) {
	var out model.MachineDetail
	if err := r.db.QueryRow(ctx, `SELECT 
    id,
    machine_name,
    tasks,
    layers,
    startup_time,
    miners,
    machine_id 
FROM harmony_machine_details 
WHERE machine_id = $1`, obj.ID).
		Scan(&out.ID, &out.MachineName, &out.Tasks, &out.Layers, &out.StartupTime, &out.Miners, &out.MachineID); err != nil {
		return nil, err
	}
	return &out, nil
}

// Tasks is the resolver for the tasks field.
func (r *machineResolver) Tasks(ctx context.Context, obj *model.Machine) ([]*model.Task, error) {
	var out []*model.Task
	if err := r.db.Select(ctx, &out, `SELECT
    id,
    initiated_by,
    update_time,
    posted_time,
    owner_id,
    added_by,
    previous_task,
    name
FROM
    harmony_task
WHERE owner_id = $1`, obj.ID); err != nil {
		return nil, err
	}
	return out, nil
}

// TaskHistories is the resolver for the taskHistories field.
func (r *machineResolver) TaskHistories(ctx context.Context, obj *model.Machine, last int) ([]*model.TaskHistory, error) {
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
WHERE work_end > CURRENT_TIMESTAMP - INTERVAL '24 hours' 
  AND completed_by_host_and_port = $1 
ORDER BY work_end DESC LIMIT $2`, obj.HostAndPort, last); err != nil {
		return nil, err
	}
	return out, nil
}

// Storages is the resolver for the storages field.
func (r *machineResolver) Storages(ctx context.Context, obj *model.Machine) ([]*model.StoragePath, error) {
	ss, err := r.loader.MachineStorages(ctx, obj.HostAndPort)
	if err != nil {
		return nil, err
	}
	cachecontrol.SetHint(ctx, cachecontrol.ScopePrivate, time.Minute)
	return ss, nil
}

// Metrics is the resolver for the metrics field.
func (r *machineResolver) Metrics(ctx context.Context, obj *model.Machine) (*model.MachineMetrics, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://%s/debug/metrics", obj.HostAndPort), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error fetching metrics: %v", err)
	}
	defer resp.Body.Close() // nolint: errcheck

	var parser expfmt.TextParser
	metrics, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing metrics: %v", err)
	}

	var result model.MachineMetrics

	parserFunc := func(name string, m *io_prometheus_client.Metric) {
		switch name {
		case "go_info":
			for _, label := range m.Label {
				if label.GetName() == "version" {
					result.GoVersion = label.GetValue()
				}
			}
		case "go_goroutines":
			result.GoRoutines = int(m.GetGauge().GetValue())
		case "process_cpu_seconds_total":
			result.ProcessCPUSecondsTotal = int(m.GetCounter().GetValue())
		case "process_start_time_seconds":
			result.ProcessStartTimeSeconds = int(m.GetGauge().GetValue())
		case "process_virtual_memory_bytes":
			result.ProcessVirtualMemoryBytes = int(m.GetGauge().GetValue())
		case "process_resident_memory_bytes":
			result.ProcessResidentMemoryBytes = int(m.GetGauge().GetValue())
		case "process_open_fds":
			result.ProcessOpenFds = int(m.GetGauge().GetValue())
		case "process_max_fds":
			result.ProcessMaxFds = int(m.GetGauge().GetValue())
		case "go_threads":
			result.GoThreads = int(m.GetGauge().GetValue())
		case "curio_harmonytask_cpu_usage":
			result.CPUUsage = m.GetGauge().GetValue()
		case "curio_harmonytask_ram_usage":
			result.RAMUsage = m.GetGauge().GetValue()
		case "curio_harmonytask_gpu_usage":
			result.GpuUsage = m.GetGauge().GetValue()
		case "curio_harmonytask_active_tasks":
			for _, label := range m.Label {
				result.ActiveTasks = append(result.ActiveTasks, &model.GaugeCountValue{
					Key:   label.GetValue(),
					Value: int(m.GetGauge().GetValue()),
				})
			}
		case "curio_harmonytask_added_tasks":
			for _, label := range m.Label {
				result.AddedTasks = append(result.AddedTasks, &model.GaugeCountValue{
					Key:   label.GetValue(),
					Value: int(m.GetGauge().GetValue()),
				})
			}
		case "curio_harmonytask_tasks_completed":
			for _, label := range m.Label {
				result.TasksCompleted = append(result.TasksCompleted, &model.GaugeCountValue{
					Key:   label.GetValue(),
					Value: int(m.GetGauge().GetValue()),
				})
			}
		case "curio_harmonytask_tasks_started":
			for _, label := range m.Label {
				result.TasksStarted = append(result.TasksStarted, &model.GaugeCountValue{
					Key:   label.GetValue(),
					Value: int(m.GetGauge().GetValue()),
				})
			}
		}
	}

	//var result []metric
	for name, mf := range metrics {
		for _, m := range mf.Metric {
			parserFunc(name, m)
		}
	}
	return &result, nil
}

// Machine returns graph.MachineResolver implementation.
func (r *Resolver) Machine() graph.MachineResolver { return &machineResolver{r} }

type machineResolver struct{ *Resolver }
