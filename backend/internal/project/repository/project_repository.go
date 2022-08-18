package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/internal/project"
	"github.com/blastertwist/flag-dash/pkg/logger"
	"github.com/blastertwist/flag-dash/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type projectRepository struct {
	cfg *config.Config
	db *sqlx.DB
	logger logger.Logger
}

func NewProjectRepository(cfg *config.Config, db *sqlx.DB, logger logger.Logger) project.Repository{
	return &projectRepository{cfg:cfg, db:db, logger:logger}
}

func (pr *projectRepository) GetProjectCount(ctx context.Context) (uint8, error) {
	var rowCount int

	if err := pr.db.QueryRowxContext(ctx, GetProjectCountQuery).Scan(&rowCount); err != nil {
		return 0, err
	}

	return uint8(rowCount), nil
}

func (pr *projectRepository) GetProjects(ctx context.Context, pq *utils.PaginationQuery) ([]*dao.Project, error) {
	
	var projects []*dao.Project
	
	rows, err := pr.db.QueryxContext(ctx, GetProjectsQuery, pq.Filter, pq.Offset, pq.Limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		p := &dao.Project{}
		if err := rows.StructScan(p); err != nil {
			return nil, err
		}

		projects = append(projects, p)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (pr *projectRepository) GetProject(ctx context.Context, project *dao.Project) (*dao.Project, error){

	p := &dao.Project{}
	if err := pr.db.QueryRowxContext(ctx, GetProjectQuery, project.ID).StructScan(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (pr *projectRepository) NewProject(ctx context.Context, project *dao.Project)(*dao.Project, error){
	newProject := &dao.Project{}

	if err := pr.db.QueryRowxContext(ctx, NewProjectQuery, project.Name).StructScan(newProject); err != nil {
		return nil, err
	}

	return newProject, nil
}

func (pr *projectRepository) EditProject(ctx context.Context, project *dao.Project) (*dao.Project, error){
	editProject := &dao.Project{}

	if err := pr.db.QueryRowxContext(ctx, EditProjectQuery, project.Name, project.ID).StructScan(editProject); err != nil {
		fmt.Print(err)
		return nil, err
	}
	return editProject, nil
}

func (pr *projectRepository) DeleteProject(ctx context.Context, project *dao.Project) error{

	result, err := pr.db.ExecContext(ctx, DeleteProjectQuery, project.ID)

	if err != nil {
		return err
	}

	num, err :=result.RowsAffected()

	if err != nil {
		return err
	}

	if num == 0 {
		return errors.New("SQL Error: Failed to delete row")
	}
	
	return nil
}

func (pr *projectRepository) AddProjectAccessKey(ctx context.Context, project *dao.Project) (*dao.Project, error) {

	p := &dao.Project{}

	if err := pr.db.QueryRowxContext(ctx, AddProjectAccessKeyQuery, project.AccessKey).StructScan(p); err != nil {
		return nil, err
	}

	return p, nil
	
}

func (pr *projectRepository) GetProjectAccessKey(ctx context.Context, project *dao.Project) (*dao.Project, error) {
	p := &dao.Project{}

	if err := pr.db.QueryRowxContext(ctx, GetProjectAccessKeyQuery, project.ID).StructScan(p); err != nil {
		return nil, err
	}
	return p, nil
}