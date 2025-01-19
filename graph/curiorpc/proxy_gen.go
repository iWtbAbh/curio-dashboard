// Code generated by github.com/filecoin-project/lotus/gen/api. DO NOT EDIT.

package curiorpc

import (
	"context"

	"github.com/filecoin-project/curio/web/api/webrpc"
	"github.com/filecoin-project/go-address"
	"golang.org/x/xerrors"
)

var ErrNotSupported = xerrors.New("method not supported")

type WebRPCStruct struct {
	Internal WebRPCMethods
}

type WebRPCMethods struct {
	ActorList func(p0 context.Context) ([]string, error) ``

	ActorSectorExpirations func(p0 context.Context, p1 address.Address) (*webrpc.SectorExpirations, error) ``

	ActorSummary func(p0 context.Context) ([]webrpc.ActorSummary, error) ``

	ClusterMachines func(p0 context.Context) ([]webrpc.MachineSummary, error) ``

	ClusterNodeInfo func(p0 context.Context, p1 int64) (*webrpc.MachineInfo, error) ``

	ClusterTaskHistory func(p0 context.Context) ([]webrpc.TaskHistorySummary, error) ``

	ClusterTaskSummary func(p0 context.Context) ([]webrpc.TaskSummary, error) ``

	DealsPending func(p0 context.Context) ([]webrpc.OpenDealInfo, error) ``

	DealsSealNow func(p0 context.Context, p1 uint64, p2 uint64) error ``

	HarmonyTaskHistory func(p0 context.Context, p1 string) ([]webrpc.HarmonyTaskHistory, error) ``

	HarmonyTaskMachines func(p0 context.Context, p1 string) ([]webrpc.HarmonyMachineDesc, error) ``

	HarmonyTaskStats func(p0 context.Context) ([]webrpc.HarmonyTaskStats, error) ``

	MarketBalance func(p0 context.Context) ([]webrpc.MarketBalanceStatus, error) ``

	MoveBalanceToEscrow func(p0 context.Context, p1 string, p2 string, p3 string) (string, error) ``

	PipelinePorepRestartAll func(p0 context.Context) error ``

	PipelinePorepSectors func(p0 context.Context) ([]sectorListEntry, error) ``

	PorepPipelineSummary func(p0 context.Context) ([]webrpc.PorepPipelineSummary, error) ``

	SectorInfo func(p0 context.Context, p1 string, p2 int64) (*webrpc.SectorInfo, error) ``

	SectorRemove func(p0 context.Context, p1 int, p2 int) error ``

	SectorResume func(p0 context.Context, p1 int64, p2 int64) error ``

	SetStorageAsk func(p0 context.Context, p1 *webrpc.StorageAsk) error ``

	StorageGCApprove func(p0 context.Context, p1 int64, p2 int64, p3 int64, p4 string) error ``

	StorageGCApproveAll func(p0 context.Context) error ``

	StorageGCMarks func(p0 context.Context) ([]webrpc.StorageGCMarks, error) ``

	StorageGCUnapproveAll func(p0 context.Context) error ``

	StorageUseStats func(p0 context.Context) ([]webrpc.StorageUseStats, error) ``

	SyncerState func(p0 context.Context) ([]webrpc.RpcInfo, error) ``

	UpgradeDelete func(p0 context.Context, p1 uint64, p2 uint64) error ``

	UpgradeResetTaskIDs func(p0 context.Context, p1 uint64, p2 uint64) error ``

	UpgradeSectors func(p0 context.Context) ([]webrpc.UpgradeSector, error) ``

	WinStats func(p0 context.Context) ([]webrpc.WinStats, error) ``
}

type WebRPCStub struct {
}

func (s *WebRPCStruct) ActorList(p0 context.Context) ([]string, error) {
	if s.Internal.ActorList == nil {
		return *new([]string), ErrNotSupported
	}
	return s.Internal.ActorList(p0)
}

func (s *WebRPCStub) ActorList(p0 context.Context) ([]string, error) {
	return *new([]string), ErrNotSupported
}

func (s *WebRPCStruct) ActorSectorExpirations(p0 context.Context, p1 address.Address) (*webrpc.SectorExpirations, error) {
	if s.Internal.ActorSectorExpirations == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.ActorSectorExpirations(p0, p1)
}

func (s *WebRPCStub) ActorSectorExpirations(p0 context.Context, p1 address.Address) (*webrpc.SectorExpirations, error) {
	return nil, ErrNotSupported
}

func (s *WebRPCStruct) ActorSummary(p0 context.Context) ([]webrpc.ActorSummary, error) {
	if s.Internal.ActorSummary == nil {
		return *new([]webrpc.ActorSummary), ErrNotSupported
	}
	return s.Internal.ActorSummary(p0)
}

func (s *WebRPCStub) ActorSummary(p0 context.Context) ([]webrpc.ActorSummary, error) {
	return *new([]webrpc.ActorSummary), ErrNotSupported
}

func (s *WebRPCStruct) ClusterMachines(p0 context.Context) ([]webrpc.MachineSummary, error) {
	if s.Internal.ClusterMachines == nil {
		return *new([]webrpc.MachineSummary), ErrNotSupported
	}
	return s.Internal.ClusterMachines(p0)
}

func (s *WebRPCStub) ClusterMachines(p0 context.Context) ([]webrpc.MachineSummary, error) {
	return *new([]webrpc.MachineSummary), ErrNotSupported
}

func (s *WebRPCStruct) ClusterNodeInfo(p0 context.Context, p1 int64) (*webrpc.MachineInfo, error) {
	if s.Internal.ClusterNodeInfo == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.ClusterNodeInfo(p0, p1)
}

func (s *WebRPCStub) ClusterNodeInfo(p0 context.Context, p1 int64) (*webrpc.MachineInfo, error) {
	return nil, ErrNotSupported
}

func (s *WebRPCStruct) ClusterTaskHistory(p0 context.Context) ([]webrpc.TaskHistorySummary, error) {
	if s.Internal.ClusterTaskHistory == nil {
		return *new([]webrpc.TaskHistorySummary), ErrNotSupported
	}
	return s.Internal.ClusterTaskHistory(p0)
}

func (s *WebRPCStub) ClusterTaskHistory(p0 context.Context) ([]webrpc.TaskHistorySummary, error) {
	return *new([]webrpc.TaskHistorySummary), ErrNotSupported
}

func (s *WebRPCStruct) ClusterTaskSummary(p0 context.Context) ([]webrpc.TaskSummary, error) {
	if s.Internal.ClusterTaskSummary == nil {
		return *new([]webrpc.TaskSummary), ErrNotSupported
	}
	return s.Internal.ClusterTaskSummary(p0)
}

func (s *WebRPCStub) ClusterTaskSummary(p0 context.Context) ([]webrpc.TaskSummary, error) {
	return *new([]webrpc.TaskSummary), ErrNotSupported
}

func (s *WebRPCStruct) DealsPending(p0 context.Context) ([]webrpc.OpenDealInfo, error) {
	if s.Internal.DealsPending == nil {
		return *new([]webrpc.OpenDealInfo), ErrNotSupported
	}
	return s.Internal.DealsPending(p0)
}

func (s *WebRPCStub) DealsPending(p0 context.Context) ([]webrpc.OpenDealInfo, error) {
	return *new([]webrpc.OpenDealInfo), ErrNotSupported
}

func (s *WebRPCStruct) DealsSealNow(p0 context.Context, p1 uint64, p2 uint64) error {
	if s.Internal.DealsSealNow == nil {
		return ErrNotSupported
	}
	return s.Internal.DealsSealNow(p0, p1, p2)
}

func (s *WebRPCStub) DealsSealNow(p0 context.Context, p1 uint64, p2 uint64) error {
	return ErrNotSupported
}

func (s *WebRPCStruct) HarmonyTaskHistory(p0 context.Context, p1 string) ([]webrpc.HarmonyTaskHistory, error) {
	if s.Internal.HarmonyTaskHistory == nil {
		return *new([]webrpc.HarmonyTaskHistory), ErrNotSupported
	}
	return s.Internal.HarmonyTaskHistory(p0, p1)
}

func (s *WebRPCStub) HarmonyTaskHistory(p0 context.Context, p1 string) ([]webrpc.HarmonyTaskHistory, error) {
	return *new([]webrpc.HarmonyTaskHistory), ErrNotSupported
}

func (s *WebRPCStruct) HarmonyTaskMachines(p0 context.Context, p1 string) ([]webrpc.HarmonyMachineDesc, error) {
	if s.Internal.HarmonyTaskMachines == nil {
		return *new([]webrpc.HarmonyMachineDesc), ErrNotSupported
	}
	return s.Internal.HarmonyTaskMachines(p0, p1)
}

func (s *WebRPCStub) HarmonyTaskMachines(p0 context.Context, p1 string) ([]webrpc.HarmonyMachineDesc, error) {
	return *new([]webrpc.HarmonyMachineDesc), ErrNotSupported
}

func (s *WebRPCStruct) HarmonyTaskStats(p0 context.Context) ([]webrpc.HarmonyTaskStats, error) {
	if s.Internal.HarmonyTaskStats == nil {
		return *new([]webrpc.HarmonyTaskStats), ErrNotSupported
	}
	return s.Internal.HarmonyTaskStats(p0)
}

func (s *WebRPCStub) HarmonyTaskStats(p0 context.Context) ([]webrpc.HarmonyTaskStats, error) {
	return *new([]webrpc.HarmonyTaskStats), ErrNotSupported
}

func (s *WebRPCStruct) MarketBalance(p0 context.Context) ([]webrpc.MarketBalanceStatus, error) {
	if s.Internal.MarketBalance == nil {
		return *new([]webrpc.MarketBalanceStatus), ErrNotSupported
	}
	return s.Internal.MarketBalance(p0)
}

func (s *WebRPCStub) MarketBalance(p0 context.Context) ([]webrpc.MarketBalanceStatus, error) {
	return *new([]webrpc.MarketBalanceStatus), ErrNotSupported
}

func (s *WebRPCStruct) MoveBalanceToEscrow(p0 context.Context, p1 string, p2 string, p3 string) (string, error) {
	if s.Internal.MoveBalanceToEscrow == nil {
		return "", ErrNotSupported
	}
	return s.Internal.MoveBalanceToEscrow(p0, p1, p2, p3)
}

func (s *WebRPCStub) MoveBalanceToEscrow(p0 context.Context, p1 string, p2 string, p3 string) (string, error) {
	return "", ErrNotSupported
}

func (s *WebRPCStruct) PipelinePorepRestartAll(p0 context.Context) error {
	if s.Internal.PipelinePorepRestartAll == nil {
		return ErrNotSupported
	}
	return s.Internal.PipelinePorepRestartAll(p0)
}

func (s *WebRPCStub) PipelinePorepRestartAll(p0 context.Context) error {
	return ErrNotSupported
}

func (s *WebRPCStruct) PipelinePorepSectors(p0 context.Context) ([]sectorListEntry, error) {
	if s.Internal.PipelinePorepSectors == nil {
		return *new([]sectorListEntry), ErrNotSupported
	}
	return s.Internal.PipelinePorepSectors(p0)
}

func (s *WebRPCStub) PipelinePorepSectors(p0 context.Context) ([]sectorListEntry, error) {
	return *new([]sectorListEntry), ErrNotSupported
}

func (s *WebRPCStruct) PorepPipelineSummary(p0 context.Context) ([]webrpc.PorepPipelineSummary, error) {
	if s.Internal.PorepPipelineSummary == nil {
		return *new([]webrpc.PorepPipelineSummary), ErrNotSupported
	}
	return s.Internal.PorepPipelineSummary(p0)
}

func (s *WebRPCStub) PorepPipelineSummary(p0 context.Context) ([]webrpc.PorepPipelineSummary, error) {
	return *new([]webrpc.PorepPipelineSummary), ErrNotSupported
}

func (s *WebRPCStruct) SectorInfo(p0 context.Context, p1 string, p2 int64) (*webrpc.SectorInfo, error) {
	if s.Internal.SectorInfo == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.SectorInfo(p0, p1, p2)
}

func (s *WebRPCStub) SectorInfo(p0 context.Context, p1 string, p2 int64) (*webrpc.SectorInfo, error) {
	return nil, ErrNotSupported
}

func (s *WebRPCStruct) SectorRemove(p0 context.Context, p1 int, p2 int) error {
	if s.Internal.SectorRemove == nil {
		return ErrNotSupported
	}
	return s.Internal.SectorRemove(p0, p1, p2)
}

func (s *WebRPCStub) SectorRemove(p0 context.Context, p1 int, p2 int) error {
	return ErrNotSupported
}

func (s *WebRPCStruct) SectorResume(p0 context.Context, p1 int64, p2 int64) error {
	if s.Internal.SectorResume == nil {
		return ErrNotSupported
	}
	return s.Internal.SectorResume(p0, p1, p2)
}

func (s *WebRPCStub) SectorResume(p0 context.Context, p1 int64, p2 int64) error {
	return ErrNotSupported
}

func (s *WebRPCStruct) SetStorageAsk(p0 context.Context, p1 *webrpc.StorageAsk) error {
	if s.Internal.SetStorageAsk == nil {
		return ErrNotSupported
	}
	return s.Internal.SetStorageAsk(p0, p1)
}

func (s *WebRPCStub) SetStorageAsk(p0 context.Context, p1 *webrpc.StorageAsk) error {
	return ErrNotSupported
}

func (s *WebRPCStruct) StorageGCApprove(p0 context.Context, p1 int64, p2 int64, p3 int64, p4 string) error {
	if s.Internal.StorageGCApprove == nil {
		return ErrNotSupported
	}
	return s.Internal.StorageGCApprove(p0, p1, p2, p3, p4)
}

func (s *WebRPCStub) StorageGCApprove(p0 context.Context, p1 int64, p2 int64, p3 int64, p4 string) error {
	return ErrNotSupported
}

func (s *WebRPCStruct) StorageGCApproveAll(p0 context.Context) error {
	if s.Internal.StorageGCApproveAll == nil {
		return ErrNotSupported
	}
	return s.Internal.StorageGCApproveAll(p0)
}

func (s *WebRPCStub) StorageGCApproveAll(p0 context.Context) error {
	return ErrNotSupported
}

func (s *WebRPCStruct) StorageGCMarks(p0 context.Context) ([]webrpc.StorageGCMarks, error) {
	if s.Internal.StorageGCMarks == nil {
		return *new([]webrpc.StorageGCMarks), ErrNotSupported
	}
	return s.Internal.StorageGCMarks(p0)
}

func (s *WebRPCStub) StorageGCMarks(p0 context.Context) ([]webrpc.StorageGCMarks, error) {
	return *new([]webrpc.StorageGCMarks), ErrNotSupported
}

func (s *WebRPCStruct) StorageGCUnapproveAll(p0 context.Context) error {
	if s.Internal.StorageGCUnapproveAll == nil {
		return ErrNotSupported
	}
	return s.Internal.StorageGCUnapproveAll(p0)
}

func (s *WebRPCStub) StorageGCUnapproveAll(p0 context.Context) error {
	return ErrNotSupported
}

func (s *WebRPCStruct) StorageUseStats(p0 context.Context) ([]webrpc.StorageUseStats, error) {
	if s.Internal.StorageUseStats == nil {
		return *new([]webrpc.StorageUseStats), ErrNotSupported
	}
	return s.Internal.StorageUseStats(p0)
}

func (s *WebRPCStub) StorageUseStats(p0 context.Context) ([]webrpc.StorageUseStats, error) {
	return *new([]webrpc.StorageUseStats), ErrNotSupported
}

func (s *WebRPCStruct) SyncerState(p0 context.Context) ([]webrpc.RpcInfo, error) {
	if s.Internal.SyncerState == nil {
		return *new([]webrpc.RpcInfo), ErrNotSupported
	}
	return s.Internal.SyncerState(p0)
}

func (s *WebRPCStub) SyncerState(p0 context.Context) ([]webrpc.RpcInfo, error) {
	return *new([]webrpc.RpcInfo), ErrNotSupported
}

func (s *WebRPCStruct) UpgradeDelete(p0 context.Context, p1 uint64, p2 uint64) error {
	if s.Internal.UpgradeDelete == nil {
		return ErrNotSupported
	}
	return s.Internal.UpgradeDelete(p0, p1, p2)
}

func (s *WebRPCStub) UpgradeDelete(p0 context.Context, p1 uint64, p2 uint64) error {
	return ErrNotSupported
}

func (s *WebRPCStruct) UpgradeResetTaskIDs(p0 context.Context, p1 uint64, p2 uint64) error {
	if s.Internal.UpgradeResetTaskIDs == nil {
		return ErrNotSupported
	}
	return s.Internal.UpgradeResetTaskIDs(p0, p1, p2)
}

func (s *WebRPCStub) UpgradeResetTaskIDs(p0 context.Context, p1 uint64, p2 uint64) error {
	return ErrNotSupported
}

func (s *WebRPCStruct) UpgradeSectors(p0 context.Context) ([]webrpc.UpgradeSector, error) {
	if s.Internal.UpgradeSectors == nil {
		return *new([]webrpc.UpgradeSector), ErrNotSupported
	}
	return s.Internal.UpgradeSectors(p0)
}

func (s *WebRPCStub) UpgradeSectors(p0 context.Context) ([]webrpc.UpgradeSector, error) {
	return *new([]webrpc.UpgradeSector), ErrNotSupported
}

func (s *WebRPCStruct) WinStats(p0 context.Context) ([]webrpc.WinStats, error) {
	if s.Internal.WinStats == nil {
		return *new([]webrpc.WinStats), ErrNotSupported
	}
	return s.Internal.WinStats(p0)
}

func (s *WebRPCStub) WinStats(p0 context.Context) ([]webrpc.WinStats, error) {
	return *new([]webrpc.WinStats), ErrNotSupported
}

var _ WebRPC = new(WebRPCStruct)
