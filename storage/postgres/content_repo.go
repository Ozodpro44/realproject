package postgres

import (
	"context"
	"real_project/models"
	log "real_project/pkg/logger"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/saidamir98/udevs_pkg/logger"
)

type contentRepo struct {
	db *pgxpool.Pool
	log log.Log
}

func NewContentRepo(db *pgxpool.Pool, log log.Log) ContentRepoI {
	return &contentRepo{
		db, log, 
	}
}

//categoty
func (c *contentRepo) CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error){
	c.log.Debug("Request in CreateCategory.")

	category.CategoryID = uuid.New()
	//
	query := `
		INSERT INTO
			categories (
				category_id,
				name
			) VALUES (
				$1, $2 
			)`

	_, err := c.db.Exec(ctx, query, category.CategoryID, category.Name)
	if err != nil {
		c.log.Error("error on creating new category.", logger.Error(err))
		return nil, err
	}		

	return c.GetCategory(ctx, category.CategoryID.String())
}
func (c *contentRepo) GetCategores(ctx context.Context, page, Limit int32) (*models.GetCategoresListResp, error){
	c.log.Debug("request in creatin category.")

	query := `
		SELECT
			*
		FROM
			categories
		LIMIT
			$1
		OFFSET
			$2`

	offset := (page - 1) * Limit
	rows, err := c.db.Query(ctx,query, Limit, offset)

	if err != nil {
		c.log.Error("error on getting all.", logger.Error(err))
		return nil, err
	}

	defer rows.Close()

	var categories []*models.Category

	for rows.Next() {

		var category models.Category

		err := rows.Scan(
			&category.CategoryID,
			&category.Name,
			&category.CreatedAt,
		)

		if err != nil {
			c.log.Error("error on scan all.", logger.Error(err))
			return nil ,err
		}

		categories = append(categories, &category)
	}

	var count int32

	c.db.QueryRow(ctx, "SELECT count(*) FROM categories").Scan(&count)
	if err != nil {
		c.log.Error("error on scaning count.", logger.Error(err))
		return nil, err
	}

	return &models.GetCategoresListResp{
		Categories: categories,
		Count: count,
		}, nil
}
func (c *contentRepo) GetCategory(ctx context.Context, id string) (*models.Category, error){
	c.log.Debug("request in creatin categories.")

	
	var category models.Category

	query := `
		SELECT 
			* 
		FROM 
			categories 
		WHERE 
			category_id = $1 `

	err := c.db.QueryRow(ctx, query, id).Scan(
		&category.CategoryID,
		&category.Name,
		&category.CreatedAt,
	)

	if err != nil {
		c.log.Error("error on getting category on GetCategory.", logger.Error(err))
		return nil, err
	}

	return &category, nil
}
func (c *contentRepo) UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error){
	return nil, nil
}
func (c *contentRepo) DeleteCategory(ctx context.Context, id string) error{
	return nil
}

//sub category
func (c *contentRepo) CreateSubCategory(ctx context.Context, category *models.Subcategory) (*models.Subcategory, error) {
	return nil, nil
}
func (c *contentRepo) GetSubCategores(ctx context.Context, page, Limit int32) ([]*models.Subcategory, error) {
	return nil, nil
}
func (c *contentRepo) GetSubCategory(ctx context.Context, id string) (*models.Subcategory, error) {
	return nil, nil
}
func (c *contentRepo) UpdateSubCategory(ctx context.Context, category *models.Subcategory) (*models.Subcategory, error) {
	return nil, nil
}
func (c *contentRepo) DeleteSubCategory(ctx context.Context, id string) error {
	return nil
}

//article
func (c *contentRepo) CreateArticle(ctx context.Context, category *models.Category) (*models.Article, error) {
	return nil, nil
}
func (c *contentRepo) GetArticles(ctx context.Context, page, Limit int32) ([]*models.Article, error) {
	return nil, nil
}
func (c *contentRepo) GetArticle(ctx context.Context, id string) (*models.Article, error) {
	return nil, nil
}
func (c *contentRepo) UpdateArticle(ctx context.Context, category *models.Article) (*models.Article, error) {
	return nil, nil
}
func (c *contentRepo) DeleteArticle(ctx context.Context, id string) error {
	return nil
}