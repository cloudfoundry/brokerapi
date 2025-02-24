package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"code.cloudfoundry.org/brokerapi/v13/domain"
	"code.cloudfoundry.org/brokerapi/v13/domain/apiresponses"
	"code.cloudfoundry.org/brokerapi/v13/internal/blog"
	"code.cloudfoundry.org/brokerapi/v13/middlewares"
)

const unbindLogKey = "unbind"

func (h APIHandler) Unbind(w http.ResponseWriter, req *http.Request) {
	instanceID := req.PathValue("instance_id")
	bindingID := req.PathValue("binding_id")

	logger := h.logger.Session(req.Context(), unbindLogKey, blog.InstanceID(instanceID), blog.BindingID(bindingID))

	requestId := fmt.Sprintf("%v", req.Context().Value(middlewares.RequestIdentityKey))

	details := domain.UnbindDetails{
		PlanID:    req.FormValue("plan_id"),
		ServiceID: req.FormValue("service_id"),
	}

	if details.ServiceID == "" {
		h.respond(w, http.StatusBadRequest, requestId, apiresponses.ErrorResponse{
			Description: serviceIdError.Error(),
		})
		logger.Error(serviceIdMissingKey, serviceIdError)
		return
	}

	if details.PlanID == "" {
		h.respond(w, http.StatusBadRequest, requestId, apiresponses.ErrorResponse{
			Description: planIdError.Error(),
		})
		logger.Error(planIdMissingKey, planIdError)
		return
	}

	asyncAllowed := req.FormValue("accepts_incomplete") == "true"
	unbindResponse, err := h.serviceBroker.Unbind(req.Context(), instanceID, bindingID, details, asyncAllowed)
	if err != nil {
		switch err := err.(type) {
		case *apiresponses.FailureResponse:
			logger.Error(err.LoggerAction(), err)
			h.respond(w, err.ValidatedStatusCode(slog.New(logger)), requestId, err.ErrorResponse())
		default:
			logger.Error(unknownErrorKey, err)
			h.respond(w, http.StatusInternalServerError, requestId, apiresponses.ErrorResponse{
				Description: err.Error(),
			})
		}
		return
	}

	if unbindResponse.IsAsync {
		h.respond(w, http.StatusAccepted, requestId, apiresponses.UnbindResponse{
			OperationData: unbindResponse.OperationData,
		})
	} else {
		h.respond(w, http.StatusOK, requestId, apiresponses.EmptyResponse{})
	}
}
