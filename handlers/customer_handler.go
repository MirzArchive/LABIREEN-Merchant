package handlers

import (
	"labireen-merchant/services"
	"labireen-merchant/utilities/jwtx"
	"labireen-merchant/utilities/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MerchantHandler interface {
	GetMe(ctx *gin.Context)
}

type MerchantHandlerImpl struct {
	svc services.MerchantService
}

func NewMerchantHandler(svc services.MerchantService) *MerchantHandlerImpl {
	return &MerchantHandlerImpl{svc}
}

func (cH *MerchantHandlerImpl) GetMe(ctx *gin.Context) {
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
