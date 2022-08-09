package service

import (
	"context"
	"fmt"

	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/internal/dto"
	"github.com/blastertwist/flag-dash/internal/project"
	"github.com/blastertwist/flag-dash/pkg/utils"
)

type projectService struct {
	cfg *config.Config
	pr project.Repository
}

func NewProjectService(cfg *config.Config, pr project.Repository) project.Service{
	return &projectService{cfg:cfg, pr:pr}
}

func (ps *projectService) GetProjects(ctx context.Context, getProjectsReq *dto.GetProjectsRequest) (*dto.GetProjectsResponse, error) {
	itemCount, errCount := ps.pr.GetProjectCount(ctx)
	
	if errCount != nil {
		return nil, errCount
	}

	// TODO: Set Default Value of Pagination
	pq, errGenPQ := utils.NewPagination(getProjectsReq.Filter, getProjectsReq.Limit, getProjectsReq.PageNum, float32(itemCount))

	fmt.Print(pq)

	if errGenPQ != nil {
		return nil, errGenPQ
	}

	projects, errRepo := ps.pr.GetProjects(ctx, pq)

	if errRepo != nil {
		return nil, errRepo
	}
	
	return &dto.GetProjectsResponse{
		Projects: projects,
		Limit: pq.Limit,
		PageNum: pq.PageNum,
		MaxPage: pq.MaxNum,
	}, nil
}

func (ps *projectService) GetProject(ctx context.Context, getProjectReq *dto.GetProjectRequest) (*dto.GetProjectResponse, error) {

	project, err := ps.pr.GetProject(ctx, &dao.Project{ID:uint32(getProjectReq.ID)})
	if err != nil {
		return nil, err
	}
	return &dto.GetProjectResponse{
		Status: "200",
		Project: project,
	}, nil
}

func (ps *projectService) NewProject(ctx context.Context, newProjectReq *dto.NewProjectRequest) (*dto.NewProjectReponse, error) {
	
	newProject, errCreate := ps.pr.NewProject(ctx, &dao.Project{Name: newProjectReq.Name})
	if errCreate != nil {
		return nil, errCreate
	}

	return &dto.NewProjectReponse{
		Status: "201",
		Project: newProject,
	}, nil
}

func (ps *projectService) EditProject(ctx context.Context, editProjectReq *dto.EditProjectRequest) (*dto.EditProjectResponse, error) {
	editedProject, errEdit := ps.pr.EditProject(ctx, &dao.Project{ID:editProjectReq.ID, Name: editProjectReq.Name})

	if errEdit != nil {
		return nil, errEdit
	}
	return &dto.EditProjectResponse{
		Status: "201",
		Project: editedProject,
	}, nil
}

func (ps *projectService) DeleteProject(ctx context.Context, deleteProjectReq *dto.DeleteProjectRequest) error {

	err := ps.pr.DeleteProject(ctx, &dao.Project{ID: deleteProjectReq.ID})

	if err != nil {
		return err
	}
	
	return nil
}