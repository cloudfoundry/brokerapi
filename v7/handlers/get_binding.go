package handlers

import (
	"errors"
	"github.com/pivotal-cf/brokerapi/v7/v7/domain"
	"net/http"

	"code.cloudfoundry.org/lager"
	"github.com/gorilla/mux"
	"github.com/pivotal-cf/brokerapi/v7/v7/domain/apiresponses"
	"github.com/pivotal-cf/brokerapi/v7/v7/middlewares"
	"github.com/pivotal-cf/brokerapi/v7/v7/utils"
)

const getBindLogKey = "getBinding"

func (h APIHandler) GetBinding(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	instanceID := vars["instance_id"]
	bindingID := vars["binding_id"]

	logger := h.logger.Session(getBindLogKey, lager.Data{
		instanceIDLogKey: instanceID,
		bindingIDLogKey:  bindingID,
	}, utils.DataForContext(req.Context(), middlewares.CorrelationIDKey))

	version := getAPIVersion(req)
	if version.Minor < 14 {
		err := errors.New("get binding endpoint only supported starting with OSB version 2.14")
		h.respond(w, http.StatusPreconditionFailed, apiresponses.ErrorResponse{
			Description: err.Error(),
		})
		logger.Error(middlewares.ApiVersionInvalidKey, err)
		return
	}

	details := domain.FetchDetails{
		ServiceID: req.URL.Query().Get("service_id"),
		PlanID:    req.URL.Query().Get("plan_id"),
	}

	binding, err := h.serviceBroker.GetBinding(req.Context(), instanceID, bindingID, details)
	if err != nil {
		switch err := err.(type) {
		case *apiresponses.FailureResponse:
			logger.Error(err.LoggerAction(), err)
			h.respond(w, err.ValidatedStatusCode(logger), err.ErrorResponse())
		default:
			logger.Error(unknownErrorKey, err)
			h.respond(w, http.StatusInternalServerError, apiresponses.ErrorResponse{
				Description: err.Error(),
			})
		}
		return
	}

	h.respond(w, http.StatusOK, apiresponses.GetBindingResponse{
		BindingResponse: apiresponses.BindingResponse{
			Credentials:     binding.Credentials,
			SyslogDrainURL:  binding.SyslogDrainURL,
			RouteServiceURL: binding.RouteServiceURL,
			VolumeMounts:    binding.VolumeMounts,
		},
		Parameters: binding.Parameters,
	})
}
