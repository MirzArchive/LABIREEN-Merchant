package handlers

import (
	"labireen-merchant/pkg/jwtx"
	"labireen-merchant/pkg/response"
	"labireen-merchant/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MerchantHandler interface {
	GetMe(ctx *gin.Context)
}

type merchantHandlerImpl struct {
	svc services.MerchantService
}

func NewMerchantHandler(svc services.MerchantService) *merchantHandlerImpl {
	return &merchantHandlerImpl{svc}
}

func (cH *merchantHandlerImpl) GetMe(ctx *gin.Context) {
	temp, exist := ctx.Get("currentUser")
	if !exist {
		log := response.ErrorLog{
			Field:   "token",
			Message: "Key token value does not exists",
		}
		response.Error(ctx, http.StatusNotFound, "Key error", log)
	}

	user := temp.(jwtx.MerchantClaims)

	userResp, err := cH.svc.GetMerchant(user.ID)
	if err != nil {
		log := response.ErrorLog{
			Field:   "ID",
			Message: "ID sent is not valid",
		}
		response.Error(ctx, http.StatusNotFound, "Cannot find requested data", log)
		return
	}

	response.Success(ctx, http.StatusOK, "success", userResp)
}
