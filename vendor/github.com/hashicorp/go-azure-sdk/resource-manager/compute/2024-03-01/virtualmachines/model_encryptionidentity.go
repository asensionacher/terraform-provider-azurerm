package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionIdentity struct {
	UserAssignedIdentityResourceId *string `json:"userAssignedIdentityResourceId,omitempty"`
}
