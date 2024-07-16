package handlers

import (
	"real_project/models"
	"real_project/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *handlers) CreateCategory(ctx *gin.Context) {
	var reqBody models.CreateCategoryReq

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		h.log.Error("error in binding req body", logger.Error(err))
		return
	}
	category := &models.Category{}

	helpers.DataParser(reqBody, &category)

	category, err = h.storage.GetContentRepo().CreateCategory(ctx, category)
	if err != nil{
		h.log.Error("error on creating new category.", logger.Error(err))
		return
	}

	ctx.JSON(201, category)
}

func (h *handlers) GetCategoriesList(ctx *gin.Context) {
	pageS := ctx.Query("page")
	limitS := ctx.Query("limit")

	categories, err := h.storage.GetContentRepo().GetCategores(ctx, helpers.GetPage(pageS), helpers.GetLimit(limitS))

	if err != nil{
		h.log.Error("error on creating new category.", logger.Error(err))
		return
	}

	ctx.JSON(200, categories)
}

func (h *handlers) GetCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	category := &models.Category{}

	category, err := h.storage.GetContentRepo().GetCategory(ctx, id)
	if err != nil{
		h.log.Error("error on getting category.", logger.Error(err))
		return
	}

	ctx.JSON(201, category)
}
