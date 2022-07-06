package repository

import (
	"context"
	"database/sql"

	"github.com/blastertwist/flag-dash/config"
	"github.com/blastertwist/flag-dash/internal/auth"
	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/blastertwist/flag-dash/pkg/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type authRepo struct {
	cfg *config.Config
	db *sqlx.DB
	logger logger.Logger
}

func NewAuthRepository(cfg *config.Config, db *sqlx.DB, logger logger.Logger) auth.Repository{
	return &authRepo{cfg:cfg, db:db, logger: logger}
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

func (r *authRepo)	Update(ctx context.Context, user *dao.User) (*dao.User, error) {
	return nil, nil
}

func (r *authRepo)	Delete(ctx context.Context, userID uint32) error {
	return nil
}

func (r *authRepo)	FindByID(ctx context.Context, userID uint32) (*dao.User, error) {
	return nil, nil
}

func (r *authRepo)	FindByEmail(ctx context.Context, email string) (*dao.User, error) {
	u := &dao.User{}
	if err := r.db.QueryRowxContext(ctx, findUserByEmailQuery, email).StructScan(u); err != nil{
		r.logger.Error("[DB Query]", zap.String("Error", err.Error()), zap.String("Message", "Query error at findUserByEmailQuery"))
		return u, err
	}
	return u, nil
}

