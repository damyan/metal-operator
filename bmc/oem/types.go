// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package oem

import "github.com/damyan/gofish/redfish"

type SimpleUpdateRequestBody struct {
	redfish.SimpleUpdateParameters
	RedfishOperationApplyTime redfish.OperationApplyTime `json:"@Redfish.OperationApplyTime,omitempty"`
}
