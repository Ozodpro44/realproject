package postgres

import (
	"context"
	"fmt"
	"real_project/models"
	log "real_project/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/saidamir98/udevs_pkg/logger"
)


type CommonRepo struct {
	db *pgxpool.Pool
	log log.Log
}

func NewCommonRepo(db *pgxpool.Pool, log log.Log) CommonRepoI {
	return &CommonRepo{db, log}
}

func (c *CommonRepo) CheckIsExists(ctx context.Context, req *models.Common) (bool, error) {
	var isExists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s = %s)", req.TableName, req.ColumnName, req.ExpValue)
	
	err := c.db.QueryRow(ctx, query).Scan(&isExists)
	if err != nil {
		c.log.Error("error on checking is exists.", logger.Error(err))
		return false , err
	}

	return isExists ,nil
}
