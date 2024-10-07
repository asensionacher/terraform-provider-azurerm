package grafanaresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedGrafanaPropertiesUpdateParameters struct {
	ApiKey                   *ApiKey                   `json:"apiKey,omitempty"`
	DeterministicOutboundIP  *DeterministicOutboundIP  `json:"deterministicOutboundIP,omitempty"`
	EnterpriseConfigurations *EnterpriseConfigurations `json:"enterpriseConfigurations,omitempty"`
	GrafanaConfigurations    *GrafanaConfigurations    `json:"grafanaConfigurations,omitempty"`
	GrafanaIntegrations      *GrafanaIntegrations      `json:"grafanaIntegrations,omitempty"`
	GrafanaMajorVersion      *string                   `json:"grafanaMajorVersion,omitempty"`
	GrafanaPlugins           *map[string]GrafanaPlugin `json:"grafanaPlugins,omitempty"`
	PublicNetworkAccess      *PublicNetworkAccess      `json:"publicNetworkAccess,omitempty"`
	ZoneRedundancy           *ZoneRedundancy           `json:"zoneRedundancy,omitempty"`
}
