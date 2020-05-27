package presenter

import (
	"belajar-gcp/domain/usecase"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	usecase usecase.UseCase
}

func NewHTTPHandler(usecase usecase.UseCase) *HTTPHandler {
	return &HTTPHandler{
		usecase: usecase,
	}
}

type (

	// Response - structure response
	Response struct {
		Ctx     echo.Context           `json:"-"`
		Status  int                    `json:"-"`
		Code    int                    `json:"code"`
		Success bool                   `json:"success"`
		Data    map[string]interface{} `json:"data"`
		Errors  []string               `json:"errors"`
	}
)

func newResponse(status, code int, success bool, data map[string]interface{}, errors []string) *Response {
	return &Response{
		Status:  status,
		Success: success,
		Code:    code,
		Data:    data,
		Errors:  errors,
	}
}

// WriteResponse - write response to the client
func (r *Response) WriteResponse(ctx echo.Context) error {
	return ctx.JSON(r.Status, r)
}

func (h *HTTPHandler) FindAllUser(c echo.Context) error {
	// logCtx := fmt.Sprintf("%T.FindAllUser", *h)
	// ctx := c.Request().Context()
	data := make(map[string]interface{})
	errs := make([]string, 0)

	resp, err := h.usecase.FindAllUser()
	if err != nil {
		errs = append(errs, err.Error())
		// helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "register_lender", ctx.Value("Correlation-ID"))
		return newResponse(http.StatusInternalServerError, http.StatusInternalServerError, false, data, errs).WriteResponse(c)
	}
	data["user"] = resp
	return newResponse(http.StatusInternalServerError, http.StatusInternalServerError, true, data, errs).WriteResponse(c)
}
