package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/mathemartins/fraudEngine/db/sqlc"
)

type createUserWatchRequest struct {
	user_id      string `json:"user_id" binding:"required"`
	watch_reason string `json:"watch_reason" binding:"required"`
}

func (server *Server) onboardWatchTrigger(ctx *gin.Context) {
	var req createUserWatchRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user_uid, err := uuid.Parse(req.user_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserWatchParams{
		UserID:      user_uid,
		WatchReason: req.watch_reason,
	}

	userWatch, err := server.store.CreateUserWatch(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userWatch)
}
