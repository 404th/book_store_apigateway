package v1

import (
	"net/http"
	"strconv"

	"github.com/404th/book_store_apigateway/genproto/book_service"
	"github.com/gin-gonic/gin"
)

// Book godoc
// @ID create-book
// @Router /v1/book [POST]
// @Summary create book
// @Description create book
// @Tags book
// @Accept json
// @Produce json
// @Param book body book_service.CreateBookRequest true "book"
// @Success 200 {object} models.ResponseModel{data=book_service.IDTracker} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateBook(c *gin.Context) {
	var body book_service.CreateBookRequest

	if err := c.BindJSON(&body); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "could not bind json with struct (CreateBook)", err)
		return
	}

	resp, err := h.services.BookService().CreateBook(c, &body)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "cannot create book", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "book created", resp)
}

// Book godoc
// @ID get-all-books
// @Router /v1/book [GET]
// @Summary get all books
// @Description get all books
// @Tags book
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Param search query string false "search"
// @Success 200 {object} models.ResponseModel{data=book_service.GetAllBooksResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllBooks(c *gin.Context) {
	var body book_service.GetAllBooksRequest = book_service.GetAllBooksRequest{
		Limit:  10,
		Offset: 0,
		Search: "",
	}

	off_q, exists := c.GetQuery("offset")
	if exists {
		iOff, err := strconv.Atoi(off_q)
		if err != nil {
			h.handleErrorResponse(c, http.StatusBadRequest, "invalid offset (GetAllBooks)", err)
			return
		}
		body.Offset = int32(iOff)
	}

	lim_q, exists := c.GetQuery("limit")
	if exists {
		iLim, err := strconv.Atoi(lim_q)
		if err != nil {
			h.handleErrorResponse(c, http.StatusBadRequest, "invalid limit (GetAllBooks)", err)
			return
		}
		body.Limit = int32(iLim)
	}

	sea_q, exists := c.GetQuery("search")
	if exists {
		body.Search = sea_q
	}

	resp, err := h.services.BookService().GetAllBooks(c, &body)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "could not get all books (GetAllBooks)", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "got all books", resp)
}

// Book godoc
// @ID get-by-id-book
// @Router /v1/book/{id} [GET]
// @Summary get book
// @Description get book
// @Tags book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.ResponseModel{data=book_service.Book} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetBookByID(c *gin.Context) {
	var body book_service.GetBookByIDRequest

	body.Id = c.Param("id")

	resp, err := h.services.BookService().GetBookByID(c, &body)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "could not get book by id (GetBookByID)", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "got book by id", resp)
}

// Book godoc
// @ID update-book
// @Router /v1/book/{id} [PUT]
// @Summary update book
// @Description update book
// @Tags book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param book body book_service.UpdateBookRequest true "desc"
// @Success 200 {object} models.ResponseModel{data=book_service.IDTracker} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateBook(c *gin.Context) {
	var body book_service.UpdateBookRequest

	if err := c.BindJSON(&body); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "could not bind json with struct (UpdateBook)", err)
		return
	}

	body.Id = c.Param("id")

	resp, err := h.services.BookService().UpdateBook(c, &body)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "cannot update book", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "book updated", resp)
}

// Book godoc
// @ID delete-book
// @Router /v1/book/{id} [DELETE]
// @Summary delete book
// @Description delete book
// @Tags book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.ResponseModel{data=book_service.IDTracker} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteBook(c *gin.Context) {
	var body book_service.DeleteBookRequest

	body.Id = c.Param("id")

	resp, err := h.services.BookService().DeleteBook(c, &body)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "could not get book by id (DeleteBook)", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "deleted book by id", resp)
}
