package handler

import (
	"context"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/youtube-api-golang/internal/apis/operations/user"
	"github.com/youtube-api-golang/internal/apis/operations/video"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/shared"
)

func (h *handler) AddVideoHandler() video.PostVideoHandlerFunc {
	return func(params video.PostVideoParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		result, err := h.useCase.AddVideo(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return user.NewPostVideoBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return video.NewPostVideoOK().WithPayload(result)
	}
}

func (h *handler) DeleteVideoHandler() video.DeleteVideoIDHandlerFunc {
	return func(params video.DeleteVideoIDParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		result, err := h.useCase.DeleteVideo(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return video.NewDeleteVideoIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return video.NewDeleteVideoIDOK().WithPayload(result)
	}
}

func (h *handler) UpdateVideoHandler() video.PutVideoIDHandlerFunc {
	return func(params video.PutVideoIDParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		result, err := h.useCase.UpdateVideo(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return video.NewPutVideoIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return video.NewPutVideoIDOK().WithPayload(result)
	}
}

func (h *handler) GetVideoByIDHandler() video.GetVideoIDHandlerFunc {
	return func(params video.GetVideoIDParams) middleware.Responder {
		res, err := h.useCase.GetVideoByID(context.Background(), params)
		if err != nil {
			return video.NewGetVideoIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return video.NewGetVideoIDOK().WithPayload(res)
	}
}

func (h *handler) UpdateViewHandler() video.PatchVideoViewIDHandlerFunc {
	return func(params video.PatchVideoViewIDParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		res, err := h.useCase.UpdateView(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return video.NewPatchVideoViewIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return video.NewPatchVideoViewIDOK().WithPayload(
			&video.PatchVideoViewIDOKBody{
				Message:    "Success",
				TotalViews: int64(*res),
			},
		)
	}
}

func (h *handler) GetRandomVideosHandler() video.GetVideoRandomHandlerFunc {
	return func(params video.GetVideoRandomParams) middleware.Responder {
		res, err := h.useCase.GetRandomVideos(context.Background())
		if err != nil {
			return video.NewGetVideoRandomBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		var arr []*models.Video
		for i, _ := range res {
			arr = append(arr, &res[i])
		}

		return video.NewGetVideoRandomOK().WithPayload(&video.GetVideoRandomOKBody{
			Message: "Success",
			Videos:  arr,
		})
	}
}

func (h *handler) GetTrendingVideosHandler() video.GetVideoTrendHandlerFunc {
	return func(params video.GetVideoTrendParams) middleware.Responder {
		res, err := h.useCase.GetTrendingVideos(context.Background())
		if err != nil {
			return video.NewGetVideoTrendBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		var arr []*models.Video
		for i, _ := range res {
			arr = append(arr, &res[i])
		}

		return video.NewGetVideoTrendOK().WithPayload(&video.GetVideoTrendOKBody{
			Message: "Success",
			Videos:  arr,
		})
	}
}

func (h *handler) GetVideosFromSubscribedChannelsHandler() video.GetVideoSubHandlerFunc {
	return func(params video.GetVideoSubParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		res, err := h.useCase.GetVideosFromSubscribedChannels(context.Background(), tokenClaim.UserID)

		if err != nil {
			return video.NewGetVideoSubBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		var arr []*models.Video
		for _, v := range res {
			arr = append(arr, &v)
		}

		return video.NewGetVideoSubOK().WithPayload(&video.GetVideoSubOKBody{
			Message: "Success",
			Videos:  arr,
		})
	}
}
