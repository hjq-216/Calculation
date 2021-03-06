package router

import (
	db "Calculation/db/sqlc"
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
)

//createAccountRequest JSON
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//CreateAccount 创建account
func (server *Server) CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Message:  InitMessageFilename(req.Owner),
		Mistakes: InitMistakesFilename(req.Owner),
		Password: req.Password,
	}

	account, err := server.queries.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

//getAccountRequest JSON
type getAccountRequest struct {
	Owner    string `form:"owner"`
	Password string `form:"password"`
}

//LogIn handle login
func (server *Server) LogIn(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.queries.GetAccount(ctx, req.Owner)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if account.Password != req.Password {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

//Register handle register
func (server *Server) Register(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.queries.GetAccount(ctx, req.Owner)
	if err != nil {
		if err == sql.ErrNoRows {
			arg := db.CreateAccountParams{
				Owner:    req.Owner,
				Message:  InitMessageFilename(req.Owner),
				Mistakes: InitMistakesFilename(req.Owner),
				Password: req.Password,
			}
			account1, err1 := server.queries.CreateAccount(ctx, arg)
			if err1 != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errorResponse(err))
					return
				}
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, account1)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

//Getfilename gername
func (server *Server) Getfilename(owner string, ctx *gin.Context) string {
	account, _ := server.queries.GetAccount(ctx, owner)
	return account.Mistakes
}
