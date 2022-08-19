package repository

import (
	"context"
	"errors"

	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/internal/flag"
	"github.com/blastertwist/flag-dash/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type flagRepo struct {
	cfg *config.Config
	db 	*sqlx.DB
}

func NewFlagRepo(cfg *config.Config, db *sqlx.DB) flag.Repository {
	return &flagRepo{cfg:cfg, db:db}
}

func (fr *flagRepo) GetAllFlags(ctx context.Context, projectID uint32) ([]*dao.Flag, error) {
	var flags []*dao.Flag

	rows, err := fr.db.QueryxContext(ctx, GetAllFlagsQuery, projectID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		f := &dao.Flag{}
		if err = rows.StructScan(f); err != nil {
			return nil, err
		}
		flags = append(flags, f)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return flags, nil
}

func (fr *flagRepo) GetFlagsCount(ctx context.Context, projectID uint32) (uint32, error) {
	var itemCount uint32
	if err := fr.db.QueryRowxContext(ctx, CountFlagItemQuery).Scan(&itemCount); err != nil {
		return 0, err
	}
	return itemCount, nil
}

func (fr *flagRepo) GetFlags(ctx context.Context, pq *utils.PaginationQuery, projectID uint32) ([]*dao.Flag, error) {
	var flags []*dao.Flag

	rows, err := fr.db.QueryxContext(ctx, GetFlagsQuery, projectID, pq.Filter, pq.Offset, pq.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		f := &dao.Flag{}
		if err = rows.StructScan(f); err != nil {
			return nil, err
		}
		flags = append(flags, f)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return flags, nil
}

func (fr *flagRepo) GetFlag(ctx context.Context, flag *dao.Flag) (*dao.Flag, error) {
	f := &dao.Flag{}
	if err := fr.db.QueryRowxContext(ctx, GetFlagQuery, flag.ID).StructScan(f); err != nil {
		return nil, err
	}
	return f, nil
}

func (fr *flagRepo) NewFlag(ctx context.Context, flag *dao.Flag) (*dao.Flag, error) {
	f := &dao.Flag{}
	if err := fr.db.QueryRowxContext(ctx, NewFlagQuery, flag.ProjectID, flag.Name, false).StructScan(f); err != nil {
		return nil, err
	}
	return f, nil
}

func (fr *flagRepo) EditFlag(ctx context.Context, flag *dao.Flag) (*dao.Flag, error) {
	f := &dao.Flag{}
	if err:= fr.db.QueryRowxContext(ctx, EditFlagQuery, flag.Name, flag.Active, flag.ID).StructScan(f); err != nil {
		return nil, err
	}
	return f, nil
}

func (fr *flagRepo) DeleteFlag(ctx context.Context, flag *dao.Flag) error {
	result, err := fr.db.ExecContext(ctx, DeleteFlagQuery, flag.ID)

	if err != nil {
		return err
	}

	num, err := result.RowsAffected()


	if err != nil {
		return err
	}

	if num == 0 {
		return errors.New("SQL Error: Failed to delete row")
	}
	return nil	
}

// Flag Context Repo Functions
func (fr *flagRepo) GetFlagContextsCount(ctx context.Context, flagID uint32) (uint32, error) {
	var itemCount uint32

	if err := fr.db.QueryRowxContext(ctx, GetFlagContextsCountQuery, flagID).Scan(&itemCount); err != nil {
		return 0, err
	}

	return itemCount, nil
}

func (fr *flagRepo) GetFlagContexts(ctx context.Context, pq *utils.PaginationQuery, flagID uint32) ([]*dao.FlagContext, error) {
	var flagContexts []*dao.FlagContext
	
	rows, err := fr.db.QueryxContext(ctx, GetFlagContextsQuery, flagID, pq.Filter, pq.Offset, pq.Limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		fc := &dao.FlagContext{}

		if err = rows.StructScan(fc); err != nil {
			return nil, err
		}

		flagContexts = append(flagContexts, fc)
	}
	return flagContexts, nil
}

func (fr *flagRepo) GetFlagContext(ctx context.Context, flagContext *dao.FlagContext) (*dao.FlagContext, error) {
	fc := &dao.FlagContext{}
	if err := fr.db.QueryRowxContext(ctx, GetFlagContextQuery, flagContext.ID).StructScan(fc); err != nil {
		return nil, err
	}
	return fc, nil
}

func (fr *flagRepo) NewFlagContext(ctx context.Context, flagContext *dao.FlagContext) (*dao.FlagContext, error) {
	fc := &dao.FlagContext{}
	if err := fr.db.QueryRowxContext(ctx, NewFlagContextQuery, flagContext.FlagID, flagContext.Name, flagContext.Condition, flagContext.Value).StructScan(fc); err != nil {
		return nil, err
	}
	return fc, nil
}

func (fr *flagRepo) EditFlagContext(ctx context.Context, flagContext *dao.FlagContext) (*dao.FlagContext, error) {
	fc := &dao.FlagContext{}
	if err := fr.db.QueryRowxContext(ctx, EditFlagContextQuery, flagContext.Name, flagContext.Condition, flagContext.Value, flagContext.ID).StructScan(fc); err != nil {
		return nil, err
	}
	return fc, nil
}

func (fr *flagRepo) DeleteFlagContext(ctx context.Context, flagContext *dao.FlagContext) error {
	result, err := fr.db.ExecContext(ctx, DeleteFlagContextQuery, flagContext.ID)

	if err != nil {
		return err
	}

	num, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if num == 0 {
		return errors.New("SQL Error: Failed to delete row")
	}

	return nil
}