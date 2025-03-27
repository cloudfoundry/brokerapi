// Copyright (C) 2015-Present Pivotal Software, Inc. All rights reserved.

// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package brokerapi

import (
	"code.cloudfoundry.org/brokerapi/v13/domain"
	"code.cloudfoundry.org/brokerapi/v13/domain/apiresponses"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate -o fakes/auto_fake_service_broker.go -fake-name AutoFakeServiceBroker . ServiceBroker

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
// Each method of the ServiceBroker interface maps to an individual endpoint of the Open Service Broker API.
//
// The specification is available here: https://github.com/openservicebrokerapi/servicebroker/blob/v2.14/spec.md
//
// The OpenAPI documentation is available here: http://petstore.swagger.io/?url=https://raw.githubusercontent.com/openservicebrokerapi/servicebroker/v2.14/openapi.yaml
type ServiceBroker interface {
	domain.ServiceBroker
}

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type DetailsWithRawParameters interface {
	domain.DetailsWithRawParameters
}

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type DetailsWithRawContext interface {
	domain.DetailsWithRawContext
}

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type ProvisionDetails = domain.ProvisionDetails

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type ProvisionedServiceSpec = domain.ProvisionedServiceSpec

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type GetInstanceDetailsSpec = domain.GetInstanceDetailsSpec

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type UnbindSpec = domain.UnbindSpec

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type BindDetails = domain.BindDetails

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type BindResource = domain.BindResource

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type UnbindDetails = domain.UnbindDetails

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type UpdateServiceSpec = domain.UpdateServiceSpec

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type DeprovisionServiceSpec = domain.DeprovisionServiceSpec

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type DeprovisionDetails = domain.DeprovisionDetails

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type UpdateDetails = domain.UpdateDetails

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type PreviousValues = domain.PreviousValues

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type PollDetails = domain.PollDetails

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type LastOperation = domain.LastOperation

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type LastOperationState = domain.LastOperationState

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
const (
	InProgress LastOperationState = "in progress"
	Succeeded  LastOperationState = "succeeded"
	Failed     LastOperationState = "failed"
)

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type Binding = domain.Binding

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type GetBindingSpec = domain.GetBindingSpec

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type VolumeMount = domain.VolumeMount

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type SharedDevice = domain.SharedDevice

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain/apiresponses
var (
	ErrInstanceAlreadyExists = apiresponses.ErrInstanceAlreadyExists

	ErrInstanceDoesNotExist = apiresponses.ErrInstanceDoesNotExist

	ErrInstanceNotFound = apiresponses.ErrInstanceNotFound

	ErrInstanceLimitMet = apiresponses.ErrInstanceLimitMet

	ErrBindingAlreadyExists = apiresponses.ErrBindingAlreadyExists

	ErrBindingDoesNotExist = apiresponses.ErrBindingDoesNotExist

	ErrBindingNotFound = apiresponses.ErrBindingNotFound

	ErrAsyncRequired = apiresponses.ErrAsyncRequired

	ErrPlanChangeNotSupported = apiresponses.ErrPlanChangeNotSupported

	ErrRawParamsInvalid = apiresponses.ErrRawParamsInvalid

	ErrAppGuidNotProvided = apiresponses.ErrAppGuidNotProvided

	ErrPlanQuotaExceeded = apiresponses.ErrPlanQuotaExceeded

	ErrServiceQuotaExceeded = apiresponses.ErrServiceQuotaExceeded

	ErrConcurrentInstanceAccess = apiresponses.ErrConcurrentInstanceAccess

	ErrMaintenanceInfoConflict = apiresponses.ErrMaintenanceInfoConflict

	ErrMaintenanceInfoNilConflict = apiresponses.ErrMaintenanceInfoNilConflict
)
