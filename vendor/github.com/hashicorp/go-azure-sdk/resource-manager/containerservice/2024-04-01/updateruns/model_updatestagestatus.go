package updateruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateStageStatus struct {
	AfterStageWaitStatus *WaitStatus          `json:"afterStageWaitStatus,omitempty"`
	Groups               *[]UpdateGroupStatus `json:"groups,omitempty"`
	Name                 *string              `json:"name,omitempty"`
	Status               *UpdateStatus        `json:"status,omitempty"`
}
