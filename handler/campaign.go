package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campsHandler struct {
	campsService campaign.Service
}

func NewCampHandler(campsService campaign.Service) *campsHandler {
	return &campsHandler{campsService}
}

func (h *campsHandler) RegisterCamps(c *gin.Context) {
	var input campaign.CampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Account has been failed1", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newCamps, err := h.campsService.RegisterCampaign(input)
	if err != nil {
		response := helper.APIResponse("Account has been failed2", http.StatusBadRequest, "erorr", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := campaign.FormatCamps(newCamps)
	response := helper.APIResponse("Account has been registered3", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
