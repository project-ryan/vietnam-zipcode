package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetLocationRequest struct {
	ZipCode string `uri:"zip_code"`
}

func (s *Server) getLocation(ctx *gin.Context) {
	var req GetLocationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println(req)

	location, err := s.store.GetLocation(req.ZipCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if location == nil {
		err := errors.New("zip code not found")
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, location)
}
