package service

import (
	"context"

	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/internal/dto"
	"github.com/blastertwist/flag-dash/internal/flag"
	"github.com/blastertwist/flag-dash/pkg/utils"
)

type flagService struct {
	cfg *config.Config
	fr flag.Repository
}

func NewFlagService(cfg *config.Config, fr flag.Repository) flag.Service{
	return &flagService{cfg:cfg, fr:fr}
}

// Flag
func (fs *flagService) GetAllFlags(ctx context.Context, getAllFlagsReq *dto.GetAllFlagsRequest) (*dto.GetAllFlagsResponse, error){
	res, err := fs.fr.GetAllFlags(ctx, getAllFlagsReq.ProjectID)

	if err != nil {
		return nil, err
	}

	return &dto.GetAllFlagsResponse{
		Status: "200",
		Flags: res,
	}, nil
}

func (fs *flagService) GetFlags(ctx context.Context, getFlagsReq *dto.GetFlagsRequest) (*dto.GetFlagsResponse, error) {
	itemCount, err := fs.fr.GetFlagsCount(ctx, getFlagsReq.ProjectID)
	if err != nil {
		return nil, err
	}
	pq, err := utils.NewPagination(getFlagsReq.Filter, getFlagsReq.Limit, getFlagsReq.PageNum, float32(itemCount))
	if err != nil {
		return nil, err
	}
	flags, err := fs.fr.GetFlags(ctx, pq, getFlagsReq.ProjectID)
	if err != nil {
		return nil, err
	}
	return &dto.GetFlagsResponse{
		Flags: flags,
		Limit: pq.Limit,
		PageNum: pq.PageNum,
		MaxPage: pq.MaxNum,
	}, nil
}

func (fs *flagService) GetFlag(ctx context.Context, getFlagReq *dto.GetFlagRequest) (*dto.GetFlagResponse, error) {
	flag, err := fs.fr.GetFlag(ctx, &dao.Flag{ID: getFlagReq.ID})
	if err != nil {
		return nil, err
	}
	return &dto.GetFlagResponse{Status: "201", Flag: flag}, nil
}

func (fs *flagService) NewFlag(ctx context.Context, newFlagReq *dto.NewFlagRequest) (*dto.NewFlagResponse, error) {
	flag, err := fs.fr.NewFlag(ctx, &dao.Flag{ProjectID: newFlagReq.ProjectID, Name: newFlagReq.Name, Active: false})
	if err != nil {
		return nil, err
	}
	return &dto.NewFlagResponse{
		Status: "201",
		Flag: flag,
	}, nil
}

func (fs *flagService) EditFlag(ctx context.Context, editFlagReq *dto.EditFlagRequest) (*dto.EditFlagResponse, error) {
	flag, err := fs.fr.EditFlag(ctx, &dao.Flag{ID: editFlagReq.ID, Name: editFlagReq.Name, Active: editFlagReq.Active})
	if err != nil {
		return nil, err
	}
	return &dto.EditFlagResponse{
		Status: "200",
		Flag: flag,
	}, nil
}

func (fs *flagService) DeleteFlag(ctx context.Context, deleteFlagReq *dto.DeleteFlagRequest) (*dto.DeleteFlagResponse, error) {
	err := fs.fr.DeleteFlag(ctx, &dao.Flag{ID: deleteFlagReq.ID})
	if err != nil {
		return nil, err
	}
	return &dto.DeleteFlagResponse{
		Status: "200",
		Msg: "Successfully delete flag",
	}, nil
}

// Flag Contexts
func (fs *flagService) GetFlagContexts(ctx context.Context, getFlagContextsReq *dto.GetFlagContextsRequest) (*dto.GetFlagContextsResponse, error) {
	itemCount, err := fs.fr.GetFlagContextsCount(ctx, getFlagContextsReq.FlagID)

	if err != nil {
		return nil, err
	}

	pq, err := utils.NewPagination(getFlagContextsReq.Filter, getFlagContextsReq.Limit, getFlagContextsReq.PageNum, float32(itemCount))
	flagContexts, err := fs.fr.GetFlagContexts(ctx, pq, getFlagContextsReq.FlagID)
	if err != nil {
		return nil, err
	}
	return &dto.GetFlagContextsResponse{
		FlagContexts: flagContexts,
		Limit: pq.Limit,
		PageNum: pq.PageNum,
		MaxPage: pq.MaxNum,
	}, nil
}

func (fs *flagService) GetFlagContext(ctx context.Context, getFlagContextReq *dto.GetFlagContextRequest) (*dto.GetFlagContextResponse, error) {
	flagContext, err := fs.fr.GetFlagContext(ctx, &dao.FlagContext{ID: getFlagContextReq.ID})
	if err != nil {
		return nil, err
	}
	return &dto.GetFlagContextResponse{
		Status: "200", 
		FlagContext: flagContext,
		}, nil	
}

func (fs *flagService) NewFlagContext(ctx context.Context, newFlagContextReq *dto.NewFlagContextRequest) (*dto.NewFlagContextResponse, error) {
	flagContext, err := fs.fr.NewFlagContext(ctx, 
		&dao.FlagContext{FlagID: newFlagContextReq.FlagID, 
		Name: newFlagContextReq.Name, 
		Condition: newFlagContextReq.Condition, 
		Value: newFlagContextReq.Value})
	if err != nil {
		return nil, err
	}
	return &dto.NewFlagContextResponse{
		Status: "201",
		FlagContext: flagContext,
	}, nil
}

func (fs *flagService) EditFlagContext(ctx context.Context, editFlagContextReq *dto.EditFlagContextRequest) (*dto.EditFlagContextResponse, error) {
	flagContext, err := fs.fr.EditFlagContext(ctx, 
		&dao.FlagContext{ID: editFlagContextReq.ID, 
			Name: editFlagContextReq.Name, 
			Condition: editFlagContextReq.Condition, 
			Value: editFlagContextReq.Value})
	if err != nil {
		return nil, err
	}
	return &dto.EditFlagContextResponse{
		Status: "200",
		FlagContext: flagContext,
	},nil
}

func (fs *flagService) DeleteFlagContext(ctx context.Context, deleteFlagContextReq *dto.DeleteFlagContextRequest) (*dto.DeleteFlagContextResponse, error) {
	err := fs.fr.DeleteFlagContext(ctx, &dao.FlagContext{ID:deleteFlagContextReq.ID})
	if err != nil {
		return nil, err
	}
	return &dto.DeleteFlagContextResponse{
		Status: "200",
		Msg: "Flag Context Successfuly Deleted",
	}, nil
}

