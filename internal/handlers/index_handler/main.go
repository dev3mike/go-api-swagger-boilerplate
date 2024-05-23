package indexhandler

import (
	"errors"
	"net/http"

	"github.com/dev3mike/go-api-swagger-boilerplate/internal/dtos"
	errorhandler "github.com/dev3mike/go-api-swagger-boilerplate/internal/handlers/error_handler"
	"github.com/dev3mike/go-api-swagger-boilerplate/internal/interfaces"
	profileservice "github.com/dev3mike/go-api-swagger-boilerplate/internal/services/profile"
	internalerrors "github.com/dev3mike/go-api-swagger-boilerplate/pkg/constants/internal_errors"
	"github.com/dev3mike/go-xmapper"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type IndexHandler struct {
}

func NewIndexHandler() interfaces.IndexHandler {
	return &IndexHandler{}
}

// HandleGetProfile godoc
// @Summary Get user profile
// @Description Retrieve the profile of a user by their username
// @Tags Profiles
// @Produce  json
// @Param username path string true "Username" minlength(4) maxlength(12)
// @Success 200 {object} dtos.ProfileResponseDTO
// @Failure 400 {object} apierror.ErrResponse "Invalid username"
// @Failure 404 {object} apierror.ErrResponse "Profile not found"
// @Router /profiles/{username} [get]
func (h *IndexHandler) GetRootHandler(w http.ResponseWriter, r *http.Request) {
	responseDto := dtos.ResponseDTO{
		Message: "Welcome to the API",
	}
	render.Render(w, r, &responseDto)
}

// HandleGetProfile godoc
// @Summary Get user profile
// @Description Retrieve the profile of a user by their username
// @Tags Profiles
// @Produce  json
// @Param username path string true "Username" minlength(4) maxlength(12)
// @Success 200 {object} dtos.ProfileResponseDTO
// @Failure 400 {object} apierror.ErrResponse "Invalid username"
// @Failure 404 {object} apierror.ErrResponse "Profile not found"
// @Router /profiles/{username} [get]
func (h *IndexHandler) PostRootHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	_, err := xmapper.ValidateSingleField(username, "validators:'required,minLength:4,maxLength:12'")
	if err != nil {

		if errors.Is(err, xmapper.ErrValidation) {
			errorhandler.HandleInternalError(w, r, internalerrors.ErrBadRequest)
			return
		}

		errorhandler.HandleInternalError(w, r, err)
		return
	}

	profile, error := profileservice.GetProfile(username)

	if error != nil {
		errorhandler.HandleInternalError(w, r, error)
		return
	}

	profileDto := dtos.ProfileResponseDTO{}

	err = xmapper.MapStructs(profile, &profileDto)

	if err != nil {
		errorhandler.HandleInternalError(w, r, err)
		return
	}

	render.Render(w, r, &profileDto)
}
