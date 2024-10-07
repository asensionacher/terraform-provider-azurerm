package vpngateways

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DhGroup string

const (
	DhGroupDHGroupOne              DhGroup = "DHGroup1"
	DhGroupDHGroupOneFour          DhGroup = "DHGroup14"
	DhGroupDHGroupTwo              DhGroup = "DHGroup2"
	DhGroupDHGroupTwoFour          DhGroup = "DHGroup24"
	DhGroupDHGroupTwoZeroFourEight DhGroup = "DHGroup2048"
	DhGroupECPThreeEightFour       DhGroup = "ECP384"
	DhGroupECPTwoFiveSix           DhGroup = "ECP256"
	DhGroupNone                    DhGroup = "None"
)

func PossibleValuesForDhGroup() []string {
	return []string{
		string(DhGroupDHGroupOne),
		string(DhGroupDHGroupOneFour),
		string(DhGroupDHGroupTwo),
		string(DhGroupDHGroupTwoFour),
		string(DhGroupDHGroupTwoZeroFourEight),
		string(DhGroupECPThreeEightFour),
		string(DhGroupECPTwoFiveSix),
		string(DhGroupNone),
	}
}

func (s *DhGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDhGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDhGroup(input string) (*DhGroup, error) {
	vals := map[string]DhGroup{
		"dhgroup1":    DhGroupDHGroupOne,
		"dhgroup14":   DhGroupDHGroupOneFour,
		"dhgroup2":    DhGroupDHGroupTwo,
		"dhgroup24":   DhGroupDHGroupTwoFour,
		"dhgroup2048": DhGroupDHGroupTwoZeroFourEight,
		"ecp384":      DhGroupECPThreeEightFour,
		"ecp256":      DhGroupECPTwoFiveSix,
		"none":        DhGroupNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DhGroup(input)
	return &out, nil
}

type IPsecEncryption string

const (
	IPsecEncryptionAESOneNineTwo     IPsecEncryption = "AES192"
	IPsecEncryptionAESOneTwoEight    IPsecEncryption = "AES128"
	IPsecEncryptionAESTwoFiveSix     IPsecEncryption = "AES256"
	IPsecEncryptionDES               IPsecEncryption = "DES"
	IPsecEncryptionDESThree          IPsecEncryption = "DES3"
	IPsecEncryptionGCMAESOneNineTwo  IPsecEncryption = "GCMAES192"
	IPsecEncryptionGCMAESOneTwoEight IPsecEncryption = "GCMAES128"
	IPsecEncryptionGCMAESTwoFiveSix  IPsecEncryption = "GCMAES256"
	IPsecEncryptionNone              IPsecEncryption = "None"
)

func PossibleValuesForIPsecEncryption() []string {
	return []string{
		string(IPsecEncryptionAESOneNineTwo),
		string(IPsecEncryptionAESOneTwoEight),
		string(IPsecEncryptionAESTwoFiveSix),
		string(IPsecEncryptionDES),
		string(IPsecEncryptionDESThree),
		string(IPsecEncryptionGCMAESOneNineTwo),
		string(IPsecEncryptionGCMAESOneTwoEight),
		string(IPsecEncryptionGCMAESTwoFiveSix),
		string(IPsecEncryptionNone),
	}
}

func (s *IPsecEncryption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIPsecEncryption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIPsecEncryption(input string) (*IPsecEncryption, error) {
	vals := map[string]IPsecEncryption{
		"aes192":    IPsecEncryptionAESOneNineTwo,
		"aes128":    IPsecEncryptionAESOneTwoEight,
		"aes256":    IPsecEncryptionAESTwoFiveSix,
		"des":       IPsecEncryptionDES,
		"des3":      IPsecEncryptionDESThree,
		"gcmaes192": IPsecEncryptionGCMAESOneNineTwo,
		"gcmaes128": IPsecEncryptionGCMAESOneTwoEight,
		"gcmaes256": IPsecEncryptionGCMAESTwoFiveSix,
		"none":      IPsecEncryptionNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPsecEncryption(input)
	return &out, nil
}

type IPsecIntegrity string

const (
	IPsecIntegrityGCMAESOneNineTwo  IPsecIntegrity = "GCMAES192"
	IPsecIntegrityGCMAESOneTwoEight IPsecIntegrity = "GCMAES128"
	IPsecIntegrityGCMAESTwoFiveSix  IPsecIntegrity = "GCMAES256"
	IPsecIntegrityMDFive            IPsecIntegrity = "MD5"
	IPsecIntegritySHAOne            IPsecIntegrity = "SHA1"
	IPsecIntegritySHATwoFiveSix     IPsecIntegrity = "SHA256"
)

func PossibleValuesForIPsecIntegrity() []string {
	return []string{
		string(IPsecIntegrityGCMAESOneNineTwo),
		string(IPsecIntegrityGCMAESOneTwoEight),
		string(IPsecIntegrityGCMAESTwoFiveSix),
		string(IPsecIntegrityMDFive),
		string(IPsecIntegritySHAOne),
		string(IPsecIntegritySHATwoFiveSix),
	}
}

func (s *IPsecIntegrity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIPsecIntegrity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIPsecIntegrity(input string) (*IPsecIntegrity, error) {
	vals := map[string]IPsecIntegrity{
		"gcmaes192": IPsecIntegrityGCMAESOneNineTwo,
		"gcmaes128": IPsecIntegrityGCMAESOneTwoEight,
		"gcmaes256": IPsecIntegrityGCMAESTwoFiveSix,
		"md5":       IPsecIntegrityMDFive,
		"sha1":      IPsecIntegritySHAOne,
		"sha256":    IPsecIntegritySHATwoFiveSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPsecIntegrity(input)
	return &out, nil
}

type IkeEncryption string

const (
	IkeEncryptionAESOneNineTwo     IkeEncryption = "AES192"
	IkeEncryptionAESOneTwoEight    IkeEncryption = "AES128"
	IkeEncryptionAESTwoFiveSix     IkeEncryption = "AES256"
	IkeEncryptionDES               IkeEncryption = "DES"
	IkeEncryptionDESThree          IkeEncryption = "DES3"
	IkeEncryptionGCMAESOneTwoEight IkeEncryption = "GCMAES128"
	IkeEncryptionGCMAESTwoFiveSix  IkeEncryption = "GCMAES256"
)

func PossibleValuesForIkeEncryption() []string {
	return []string{
		string(IkeEncryptionAESOneNineTwo),
		string(IkeEncryptionAESOneTwoEight),
		string(IkeEncryptionAESTwoFiveSix),
		string(IkeEncryptionDES),
		string(IkeEncryptionDESThree),
		string(IkeEncryptionGCMAESOneTwoEight),
		string(IkeEncryptionGCMAESTwoFiveSix),
	}
}

func (s *IkeEncryption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIkeEncryption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIkeEncryption(input string) (*IkeEncryption, error) {
	vals := map[string]IkeEncryption{
		"aes192":    IkeEncryptionAESOneNineTwo,
		"aes128":    IkeEncryptionAESOneTwoEight,
		"aes256":    IkeEncryptionAESTwoFiveSix,
		"des":       IkeEncryptionDES,
		"des3":      IkeEncryptionDESThree,
		"gcmaes128": IkeEncryptionGCMAESOneTwoEight,
		"gcmaes256": IkeEncryptionGCMAESTwoFiveSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IkeEncryption(input)
	return &out, nil
}

type IkeIntegrity string

const (
	IkeIntegrityGCMAESOneTwoEight IkeIntegrity = "GCMAES128"
	IkeIntegrityGCMAESTwoFiveSix  IkeIntegrity = "GCMAES256"
	IkeIntegrityMDFive            IkeIntegrity = "MD5"
	IkeIntegritySHAOne            IkeIntegrity = "SHA1"
	IkeIntegritySHAThreeEightFour IkeIntegrity = "SHA384"
	IkeIntegritySHATwoFiveSix     IkeIntegrity = "SHA256"
)

func PossibleValuesForIkeIntegrity() []string {
	return []string{
		string(IkeIntegrityGCMAESOneTwoEight),
		string(IkeIntegrityGCMAESTwoFiveSix),
		string(IkeIntegrityMDFive),
		string(IkeIntegritySHAOne),
		string(IkeIntegritySHAThreeEightFour),
		string(IkeIntegritySHATwoFiveSix),
	}
}

func (s *IkeIntegrity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIkeIntegrity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIkeIntegrity(input string) (*IkeIntegrity, error) {
	vals := map[string]IkeIntegrity{
		"gcmaes128": IkeIntegrityGCMAESOneTwoEight,
		"gcmaes256": IkeIntegrityGCMAESTwoFiveSix,
		"md5":       IkeIntegrityMDFive,
		"sha1":      IkeIntegritySHAOne,
		"sha384":    IkeIntegritySHAThreeEightFour,
		"sha256":    IkeIntegritySHATwoFiveSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IkeIntegrity(input)
	return &out, nil
}

type PfsGroup string

const (
	PfsGroupECPThreeEightFour   PfsGroup = "ECP384"
	PfsGroupECPTwoFiveSix       PfsGroup = "ECP256"
	PfsGroupNone                PfsGroup = "None"
	PfsGroupPFSMM               PfsGroup = "PFSMM"
	PfsGroupPFSOne              PfsGroup = "PFS1"
	PfsGroupPFSOneFour          PfsGroup = "PFS14"
	PfsGroupPFSTwo              PfsGroup = "PFS2"
	PfsGroupPFSTwoFour          PfsGroup = "PFS24"
	PfsGroupPFSTwoZeroFourEight PfsGroup = "PFS2048"
)

func PossibleValuesForPfsGroup() []string {
	return []string{
		string(PfsGroupECPThreeEightFour),
		string(PfsGroupECPTwoFiveSix),
		string(PfsGroupNone),
		string(PfsGroupPFSMM),
		string(PfsGroupPFSOne),
		string(PfsGroupPFSOneFour),
		string(PfsGroupPFSTwo),
		string(PfsGroupPFSTwoFour),
		string(PfsGroupPFSTwoZeroFourEight),
	}
}

func (s *PfsGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePfsGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePfsGroup(input string) (*PfsGroup, error) {
	vals := map[string]PfsGroup{
		"ecp384":  PfsGroupECPThreeEightFour,
		"ecp256":  PfsGroupECPTwoFiveSix,
		"none":    PfsGroupNone,
		"pfsmm":   PfsGroupPFSMM,
		"pfs1":    PfsGroupPFSOne,
		"pfs14":   PfsGroupPFSOneFour,
		"pfs2":    PfsGroupPFSTwo,
		"pfs24":   PfsGroupPFSTwoFour,
		"pfs2048": PfsGroupPFSTwoZeroFourEight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PfsGroup(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
		"updating":  ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type VirtualNetworkGatewayConnectionProtocol string

const (
	VirtualNetworkGatewayConnectionProtocolIKEvOne VirtualNetworkGatewayConnectionProtocol = "IKEv1"
	VirtualNetworkGatewayConnectionProtocolIKEvTwo VirtualNetworkGatewayConnectionProtocol = "IKEv2"
)

func PossibleValuesForVirtualNetworkGatewayConnectionProtocol() []string {
	return []string{
		string(VirtualNetworkGatewayConnectionProtocolIKEvOne),
		string(VirtualNetworkGatewayConnectionProtocolIKEvTwo),
	}
}

func (s *VirtualNetworkGatewayConnectionProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualNetworkGatewayConnectionProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualNetworkGatewayConnectionProtocol(input string) (*VirtualNetworkGatewayConnectionProtocol, error) {
	vals := map[string]VirtualNetworkGatewayConnectionProtocol{
		"ikev1": VirtualNetworkGatewayConnectionProtocolIKEvOne,
		"ikev2": VirtualNetworkGatewayConnectionProtocolIKEvTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualNetworkGatewayConnectionProtocol(input)
	return &out, nil
}

type VnetLocalRouteOverrideCriteria string

const (
	VnetLocalRouteOverrideCriteriaContains VnetLocalRouteOverrideCriteria = "Contains"
	VnetLocalRouteOverrideCriteriaEqual    VnetLocalRouteOverrideCriteria = "Equal"
)

func PossibleValuesForVnetLocalRouteOverrideCriteria() []string {
	return []string{
		string(VnetLocalRouteOverrideCriteriaContains),
		string(VnetLocalRouteOverrideCriteriaEqual),
	}
}

func (s *VnetLocalRouteOverrideCriteria) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVnetLocalRouteOverrideCriteria(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVnetLocalRouteOverrideCriteria(input string) (*VnetLocalRouteOverrideCriteria, error) {
	vals := map[string]VnetLocalRouteOverrideCriteria{
		"contains": VnetLocalRouteOverrideCriteriaContains,
		"equal":    VnetLocalRouteOverrideCriteriaEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VnetLocalRouteOverrideCriteria(input)
	return &out, nil
}

type VpnConnectionStatus string

const (
	VpnConnectionStatusConnected    VpnConnectionStatus = "Connected"
	VpnConnectionStatusConnecting   VpnConnectionStatus = "Connecting"
	VpnConnectionStatusNotConnected VpnConnectionStatus = "NotConnected"
	VpnConnectionStatusUnknown      VpnConnectionStatus = "Unknown"
)

func PossibleValuesForVpnConnectionStatus() []string {
	return []string{
		string(VpnConnectionStatusConnected),
		string(VpnConnectionStatusConnecting),
		string(VpnConnectionStatusNotConnected),
		string(VpnConnectionStatusUnknown),
	}
}

func (s *VpnConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnConnectionStatus(input string) (*VpnConnectionStatus, error) {
	vals := map[string]VpnConnectionStatus{
		"connected":    VpnConnectionStatusConnected,
		"connecting":   VpnConnectionStatusConnecting,
		"notconnected": VpnConnectionStatusNotConnected,
		"unknown":      VpnConnectionStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnConnectionStatus(input)
	return &out, nil
}

type VpnLinkConnectionMode string

const (
	VpnLinkConnectionModeDefault       VpnLinkConnectionMode = "Default"
	VpnLinkConnectionModeInitiatorOnly VpnLinkConnectionMode = "InitiatorOnly"
	VpnLinkConnectionModeResponderOnly VpnLinkConnectionMode = "ResponderOnly"
)

func PossibleValuesForVpnLinkConnectionMode() []string {
	return []string{
		string(VpnLinkConnectionModeDefault),
		string(VpnLinkConnectionModeInitiatorOnly),
		string(VpnLinkConnectionModeResponderOnly),
	}
}

func (s *VpnLinkConnectionMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnLinkConnectionMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnLinkConnectionMode(input string) (*VpnLinkConnectionMode, error) {
	vals := map[string]VpnLinkConnectionMode{
		"default":       VpnLinkConnectionModeDefault,
		"initiatoronly": VpnLinkConnectionModeInitiatorOnly,
		"responderonly": VpnLinkConnectionModeResponderOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnLinkConnectionMode(input)
	return &out, nil
}

type VpnNatRuleMode string

const (
	VpnNatRuleModeEgressSnat  VpnNatRuleMode = "EgressSnat"
	VpnNatRuleModeIngressSnat VpnNatRuleMode = "IngressSnat"
)

func PossibleValuesForVpnNatRuleMode() []string {
	return []string{
		string(VpnNatRuleModeEgressSnat),
		string(VpnNatRuleModeIngressSnat),
	}
}

func (s *VpnNatRuleMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnNatRuleMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnNatRuleMode(input string) (*VpnNatRuleMode, error) {
	vals := map[string]VpnNatRuleMode{
		"egresssnat":  VpnNatRuleModeEgressSnat,
		"ingresssnat": VpnNatRuleModeIngressSnat,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnNatRuleMode(input)
	return &out, nil
}

type VpnNatRuleType string

const (
	VpnNatRuleTypeDynamic VpnNatRuleType = "Dynamic"
	VpnNatRuleTypeStatic  VpnNatRuleType = "Static"
)

func PossibleValuesForVpnNatRuleType() []string {
	return []string{
		string(VpnNatRuleTypeDynamic),
		string(VpnNatRuleTypeStatic),
	}
}

func (s *VpnNatRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnNatRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnNatRuleType(input string) (*VpnNatRuleType, error) {
	vals := map[string]VpnNatRuleType{
		"dynamic": VpnNatRuleTypeDynamic,
		"static":  VpnNatRuleTypeStatic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnNatRuleType(input)
	return &out, nil
}
