package bookmarks

import (
	"net/http"

	commons "github.com/3cognito/library/app/common"
	"github.com/3cognito/library/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewController(
	service BookmarkServiceInterface,
) BookmarkControllerInterface {
	return &bookmarkController{
		service: service,
	}
}

func (c *bookmarkController) AddToBookmark(ctx *gin.Context) {
	parsedBookID, parseErr := uuid.Parse(ctx.Param("bookId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, commons.BadRequest, parseErr.Error())
		return
	}

	parsedUserID, parseErr := uuid.Parse(ctx.GetString("userId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, commons.BadRequest, parseErr.Error())
		return
	}

	if bookmarkErr := c.service.AddToBookmark(parsedUserID, parsedBookID); bookmarkErr != nil {
		if bookmarkErr.Error() == commons.ErrResourceNotFound.Error() {
			utils.JsonErrorResponse(ctx, http.StatusBadRequest, commons.BadRequest, commons.ErrBookAlreadyBookmarked.Error())
			return
		}
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, commons.BadRequest, bookmarkErr.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusCreated, "book added to bookmark", nil)
}

func (c *bookmarkController) RemoveFromBookmark(ctx *gin.Context) {
	parsedBookID, parseErr := uuid.Parse(ctx.Param("bookId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, commons.BadRequest, parseErr.Error())
		return
	}

	parsedUserID, parseErr := uuid.Parse(ctx.GetString("userId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, commons.BadRequest, parseErr.Error())
		return
	}

	if bookmarkErr := c.service.RemoveFromBookmark(parsedUserID, parsedBookID); bookmarkErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, commons.BadRequest, bookmarkErr.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusOK, "book removed from bookmark", nil)
}

func (c *bookmarkController) GetUserBookMarks(ctx *gin.Context) {
	parsedUserID, parseErr := uuid.Parse(ctx.GetString("userId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, commons.BadRequest, parseErr.Error())
		return
	}

	bookmarks, bookmarkErr := c.service.GetUserBookMarks(parsedUserID)
	if bookmarkErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, commons.BadRequest, bookmarkErr.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusOK, "user bookmarks", bookmarks)
}
