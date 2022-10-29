package handler

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/youtube-api-golang/internal/apis/operations/health"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/shared"
)

func (h *handler) GetHealthHandler() health.GetHealthHandlerFunc {
	return func(params health.GetHealthParams, tokenData interface{}) middleware.Responder {

		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		fmt.Println("CLAIM", tokenClaim)

		_, err = h.useCase.GetHealthcheck(context.Background())
		if err != nil {
			return health.NewGetHealthBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return health.NewGetHealthOK().WithPayload(&health.GetHealthOKBody{
			Data:    &models.Health{Status: "OK"},
			Message: "Success Get Health",
		}).WithAccessControlAllowOrigin(os.Getenv("CLIENT_URL"))
	}
}
