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

package apiresponses

import "code.cloudfoundry.org/brokerapi/v13/domain"

type EmptyResponse struct{}

type ErrorResponse struct {
	Error       string `json:"error,omitempty"`
	Description string `json:"description"`
}

type CatalogResponse struct {
	Services []domain.Service `json:"services"`
}

type ProvisioningResponse struct {
	DashboardURL  string `json:"dashboard_url,omitempty"`
	OperationData string `json:"operation,omitempty"`
	Metadata      any    `json:"metadata,omitempty"`
}

type GetInstanceResponse struct {
	ServiceID    string `json:"service_id"`
	PlanID       string `json:"plan_id"`
	DashboardURL string `json:"dashboard_url,omitempty"`
	Parameters   any    `json:"parameters,omitempty"`
	Metadata     any    `json:"metadata,omitempty"`
}

type UpdateResponse struct {
	DashboardURL  string `json:"dashboard_url,omitempty"`
	OperationData string `json:"operation,omitempty"`
	Metadata      any    `json:"metadata,omitempty"`
}

type DeprovisionResponse struct {
	OperationData string `json:"operation,omitempty"`
}

type LastOperationResponse struct {
	State       domain.LastOperationState `json:"state"`
	Description string                    `json:"description,omitempty"`
}

type AsyncBindResponse struct {
	OperationData string `json:"operation,omitempty"`
}

type BindingResponse struct {
	Credentials     any                  `json:"credentials,omitempty"`
	SyslogDrainURL  string               `json:"syslog_drain_url,omitempty"`
	RouteServiceURL string               `json:"route_service_url,omitempty"`
	VolumeMounts    []domain.VolumeMount `json:"volume_mounts,omitempty"`
	BackupAgentURL  string               `json:"backup_agent_url,omitempty"`
	Endpoints       []domain.Endpoint    `json:"endpoints,omitempty"`
	Metadata        any                  `json:"metadata,omitempty"`
}

type GetBindingResponse struct {
	BindingResponse
	Parameters any `json:"parameters,omitempty"`
}

type UnbindResponse struct {
	OperationData string `json:"operation,omitempty"`
}

type ExperimentalVolumeMountBindingResponse struct {
	Credentials     any                              `json:"credentials,omitempty"`
	SyslogDrainURL  string                           `json:"syslog_drain_url,omitempty"`
	RouteServiceURL string                           `json:"route_service_url,omitempty"`
	VolumeMounts    []domain.ExperimentalVolumeMount `json:"volume_mounts,omitempty"`
	BackupAgentURL  string                           `json:"backup_agent_url,omitempty"`
}
