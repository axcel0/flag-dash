package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/auth"
	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/pkg/logger"
	"github.com/blastertwist/flag-dash/pkg/utils"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type authRepo struct {
	cfg    *config.Config
	db     *sqlx.DB
	logger logger.Logger
}

func NewAuthRepository(cfg *config.Config, db *sqlx.DB, logger logger.Logger) auth.Repository {
	return &authRepo{cfg: cfg, db: db, logger: logger}
}

func (r *authRepo) GetUsersCount(ctx context.Context) (uint8, error) {
	var n int

	if err := r.db.QueryRowxContext(ctx, getUsersCountQuery).Scan(&n); err != nil {
		return 0, err
	}
	return uint8(n), nil
}

func (r *authRepo) GetUsers(ctx context.Context, pq *utils.PaginationQuery) ([]*dao.User, error) {
	var users []*dao.User

	rows, err := r.db.QueryxContext(ctx, getUsersQuery, pq.Filter, pq.Offset, pq.Limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		u := &dao.User{}
		if err := rows.StructScan(u); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *authRepo) CreateUser(ctx context.Context, user *dao.User, userProfile *dao.UserProfile) (*dao.User, *dao.UserProfile, error) {

	u := &dao.User{}
	up := &dao.UserProfile{}

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{})
	defer tx.Rollback()

	if err != nil {
		return u, up, err
	}

	if err := tx.QueryRowxContext(ctx, createUserQuery, &user.Email, &user.Password).StructScan(u); err != nil {
		r.logger.Error("[DB Transaction]", zap.String("Error", err.Error()), zap.String("Message", "Transaction error at createUserQuery."))
		tx.Rollback()
		return u, up, err
	}

	if err := tx.QueryRowxContext(ctx, createUserProfileQuery, &u.ID, &userProfile.FirstName, &userProfile.LastName, &userProfile.PhoneNumber).StructScan(up); err != nil {
		r.logger.Error("[DB Transaction]", zap.String("Error", err.Error()), zap.String("Message", "Transaction error at createUserProfileQuery."))
		return u, up, err
	}

	tx.MustExecContext(ctx, createUserRoleQuery, &u.ID)

	tx.Commit()
	return u, up, nil

}

func (r *authRepo) Update(ctx context.Context, user *dao.User) (*dao.User, error) {
	u := &dao.User{}

	if err := r.db.QueryRowxContext(ctx, updateUserQuery, user.Email, user.Password, user.FirstName, user.LastName, user.PhoneNumber, user.RoleLevel, user.ID).StructScan(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *authRepo) Delete(ctx context.Context, user *dao.User) error {
	res, err := r.db.ExecContext(ctx, deleteUserQuery, user.ID)

	if err != nil {
		return err
	}

	rowAffected, err := res.RowsAffected()

	if rowAffected == 0 {
		return errors.New("SQL Error: Failed to delete intended row")
	}
	return nil
}

func (r *authRepo) FindByID(ctx context.Context, user *dao.User) (*dao.User, error) {
	u := &dao.User{}

	if err := r.db.QueryRowxContext(ctx, findUserByIDQuery, user.ID).StructScan(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *authRepo) FindByEmail(ctx context.Context, user *dao.User) (*dao.User, error) {
	u := &dao.User{}
	if err := r.db.QueryRowxContext(ctx, findUserByEmailQuery, user.Email).StructScan(u); err != nil {
		r.logger.Error("[DB Query]", zap.String("Error", err.Error()), zap.String("Message", "Query error at findUserByEmailQuery"))
		return nil, err
	}
	return u, nil
}
