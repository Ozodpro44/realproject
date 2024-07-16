package postgres

import (
	"context"
	"real_project/models"
	log "real_project/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ownerRepo struct {
	db *pgxpool.Pool
	log log.Log
}

func NewOwnerRepo(db *pgxpool.Pool, log log.Log) OwnerRepoI {
	return &ownerRepo{db, log}
}

func (o *ownerRepo) Login(ctx context.Context, login *models.LoginOwn) (*models.Owner, error) {
	return nil, nil
}