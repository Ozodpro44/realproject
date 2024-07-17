package postgres

import (
	"context"
	"real_project/models"
)

type ContentRepoI interface {
	//categoty
	CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	GetCategores(ctx context.Context, page, Limit int32) (*models.GetCategoresListResp, error)
	GetCategory(ctx context.Context, id string) (*models.Category, error)
	UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	DeleteCategory(ctx context.Context, id string) error

	//Catogory
	CreateSubCategory(ctx context.Context, category *models.Subcategory) (*models.Subcategory, error)
	GetSubCategores(ctx context.Context, page, Limit int32) ([]*models.Subcategory, error)
	GetSubCategory(ctx context.Context, id string) (*models.Subcategory, error)
	UpdateSubCategory(ctx context.Context, category *models.Subcategory) (*models.Subcategory, error)
	DeleteSubCategory(ctx context.Context, id string) error

	//Article
	CreateArticle(ctx context.Context, category *models.Category) (*models.Article, error)
	GetArticles(ctx context.Context, page, Limit int32) ([]*models.Article, error)
	GetArticle(ctx context.Context, id string) (*models.Article, error)
	UpdateArticle(ctx context.Context, category *models.Article) (*models.Article, error)
	DeleteArticle(ctx context.Context, id string) error
}

type OwnerRepoI interface {
	Login(ctx context.Context, login *models.LoginOwn) (*models.Owner, error)
}

type CommonRepoI interface {
	CheckIsExists(ctx context.Context, req *models.Common) (bool, error)
}
