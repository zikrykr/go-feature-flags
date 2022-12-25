package rest

import (
	"net/http"

	"github.com/go-feature-flag/rest/request"
	"github.com/go-feature-flag/rest/response"
	"github.com/labstack/echo/v4"
)

type RestFeatureFlagController struct {
	e                  *echo.Echo
	restFeatureFlagSvc RestFeatureFlagItf
}

func NewRestFeatureFlagController(e *echo.Echo, restFeatureFlagSvc RestFeatureFlagItf) *RestFeatureFlagController {
	return &RestFeatureFlagController{
		e,
		restFeatureFlagSvc,
	}
}

func (ctl *RestFeatureFlagController) CreateFeatureFlag(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload request.CreateFeatureFlagReq
	)
	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	result, err := ctl.restFeatureFlagSvc.CreateFeatureFlag(ctx, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Data:    nil,
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.Response{
		Data:    result,
		Success: true,
		Message: "Feature Flag Created Successfully",
	})
}

func (ctl *RestFeatureFlagController) UpdateFeatureFlag(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload request.CreateFeatureFlagReq
	)
	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	result, err := ctl.restFeatureFlagSvc.UpdateFeatureFlag(ctx, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Data:    nil,
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Data:    result,
		Success: true,
		Message: "Feature Flag Updated Successfully",
	})
}

func (ctl *RestFeatureFlagController) GetFeatureFlags(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	result, err := ctl.restFeatureFlagSvc.GetFeatureFlags(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Data:    nil,
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Data:    result,
		Success: true,
		Message: "Get Feature Flag Successful",
	})
}

func (ctl *RestFeatureFlagController) DeleteFeatureFlag(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	flagName := c.Param("flagName")

	if err := ctl.restFeatureFlagSvc.DeleteFeatureFlag(ctx, flagName); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Data:    nil,
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Data:    nil,
		Success: true,
		Message: "Feature Flag Deleted Successfully",
	})
}
