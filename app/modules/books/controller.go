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

func (c *controller) AddBook(ctx *gin.Context) {
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

	pages, pageErr := utils.StringToInt(ctx.PostForm("pages"))
	if pageErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, BadRequest, pageErr.Error())
		return
	}
	params.Title = ctx.PostForm("title")
	params.ISBN = ctx.PostForm("isbn")
	params.Publisher = ctx.PostForm("publisher")
	params.PublicationDate = ctx.PostForm("publication_date")
	params.Pages = pages
	params.Language = ctx.PostForm("language")
	params.Description = ctx.PostForm("description")
	params.Genres = ctx.PostFormArray("genres")

	if !utils.NoEmptyFields(params) {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, BadRequest, "All fields are required")
		return
	}

	userId, parseErr := uuid.Parse(ctx.GetString("userId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, BadRequest, parseErr.Error())
		return
	}

	book, err := c.bookService.AddBook(userId, params)
	if err != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, RequestFailed, err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusCreated, RequestSuccessful, book)
}

func (c *controller) DeleteBook(ctx *gin.Context) {
	userId, parseErr := uuid.Parse(ctx.GetString("userId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, BadRequest, parseErr.Error())
		return
	}

	bookId, parseErr := uuid.Parse(ctx.Param("bookId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, BadRequest, parseErr.Error())
		return
	}

	err := c.bookService.DeleteBook(userId, bookId)
	if err != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, RequestFailed, err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusOK, RequestSuccessful, nil)
}

func (c *controller) GetAuthorBooks(ctx *gin.Context) {
	userId, parseErr := uuid.Parse(ctx.GetString("userId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, BadRequest, parseErr.Error())
		return
	}

	books, err := c.bookService.GetAuthorBooks(userId)
	if err != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, RequestFailed, err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusOK, RequestSuccessful, books)
}

func (c *controller) GetBook(ctx *gin.Context) {
	bookId, parseErr := uuid.Parse(ctx.Param("bookId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, BadRequest, parseErr.Error())
		return
	}

	book, err := c.bookService.GetBookByID(bookId)
	if err != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, RequestFailed, err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusOK, RequestSuccessful, book)
}
