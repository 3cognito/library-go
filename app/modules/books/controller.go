package books

import (
	"net/http"

	"github.com/3cognito/library/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewController(
	bookService ServiceInterface,
) ControllerInterface {
	return &controller{
		bookService: bookService,
	}
}

func (c *controller) CreateBook(ctx *gin.Context) {
	var params CreateBookRequest
	bookFile, bookFileErr := ctx.FormFile("bookFile")
	if bookFileErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, BadRequest, bookFileErr.Error())
		return
	}

	params.BookFile = bookFile

	imageFile, imageFileErr := ctx.FormFile("imageFile")
	if imageFileErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, BadRequest, imageFileErr.Error())
		return
	}

	params.ImageFile = imageFile

	userId, parseErr := uuid.Parse(ctx.GetString("userId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, BadRequest, parseErr.Error())
		return
	}

	if !utils.ValidParams(ctx, &params) {
		return
	}

	book, err := c.bookService.AddBook(userId, params)
	if err != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, RequestSuccessful, err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusCreated, RequestSuccessful, book)
}