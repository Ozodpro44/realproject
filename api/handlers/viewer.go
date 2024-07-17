package handlers

import (
	"real_project/models"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *handlers) CheckUser(ctx *gin.Context) {
	var reqBody models.CheckViewer

	gin.Bind(&reqBody)

	isExists, err := h.storage.GetCommonRepo().CheckIsExists(ctx, &models.Common{
		TableName:  "viewers",
		ColumnName: "gmail",
		ExpValue:   reqBody.Gmail,
	})
	if err != nil {
		h.log.Error("error on checking viewer.", logger.Error(err))
		return
	}

	if isExists {
		ctx.JSON(201, models.CheckExistsResp{
			IatsExists: isExists,
			Status:     "log-in",
		})
		return
	}
}
