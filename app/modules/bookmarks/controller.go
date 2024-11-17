package bookmarks

import (
	"github.com/gin-gonic/gin"
)

func NewController(
	service BookmarkServiceInterface,
) BookmarkControllerInterface {
	return &bookmarkController{
		service: service,
	}
}

func (c *bookmarkController) AddToBookmark(ctx *gin.Context) {

}

func (c *bookmarkController) RemoveFromBookmark(ctx *gin.Context) {

}

func (c *bookmarkController) GetUserBookMarks(ctx *gin.Context) {}
