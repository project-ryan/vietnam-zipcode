package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents an API error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    string `json:"code,omitempty"`
	Details string `json:"details,omitempty"`
}

func (s *Server) getLocation(ctx *gin.Context) {
	zipcode := ctx.DefaultQuery("zipcode", "")

	if zipcode == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(
			"zipcode parameter is required",
			"MISSING_ZIPCODE",
			"Please provide a zipcode query parameter",
		))
		return
	}

	loc, err := s.service.GetLocation(zipcode)
	if err != nil {
		if err.Error() == "invalid zipcode format: "+zipcode {
			ctx.JSON(http.StatusBadRequest, errorResponse(
				"Invalid zipcode format",
				"INVALID_FORMAT",
				"Zipcode must be exactly 5 digits",
			))
			return
		}

		ctx.JSON(http.StatusNotFound, errorResponse(
			"Zipcode not found",
			"NOT_FOUND",
			"The provided zipcode does not exist in our database",
		))
		return
	}

	ctx.JSON(http.StatusOK, loc)
}
