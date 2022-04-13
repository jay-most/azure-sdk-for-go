//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer

import "encoding/json"

func unmarshalAudioEncoderBaseClassification(rawMsg json.RawMessage) (AudioEncoderBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AudioEncoderBaseClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.AudioEncoderAac":
		b = &AudioEncoderAac{}
	default:
		b = &AudioEncoderBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAuthenticationBaseClassification(rawMsg json.RawMessage) (AuthenticationBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AuthenticationBaseClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.JwtAuthentication":
		b = &JwtAuthentication{}
	default:
		b = &AuthenticationBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalCertificateSourceClassification(rawMsg json.RawMessage) (CertificateSourceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CertificateSourceClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.PemCertificateList":
		b = &PemCertificateList{}
	default:
		b = &CertificateSource{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalCredentialsBaseClassification(rawMsg json.RawMessage) (CredentialsBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CredentialsBaseClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.UsernamePasswordCredentials":
		b = &UsernamePasswordCredentials{}
	default:
		b = &CredentialsBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalEncoderPresetBaseClassification(rawMsg json.RawMessage) (EncoderPresetBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EncoderPresetBaseClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.EncoderCustomPreset":
		b = &EncoderCustomPreset{}
	case "#Microsoft.VideoAnalyzer.EncoderSystemPreset":
		b = &EncoderSystemPreset{}
	default:
		b = &EncoderPresetBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalEndpointBaseClassification(rawMsg json.RawMessage) (EndpointBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EndpointBaseClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.TlsEndpoint":
		b = &TLSEndpoint{}
	case "#Microsoft.VideoAnalyzer.UnsecuredEndpoint":
		b = &UnsecuredEndpoint{}
	default:
		b = &EndpointBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalProcessorNodeBaseClassification(rawMsg json.RawMessage) (ProcessorNodeBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProcessorNodeBaseClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.EncoderProcessor":
		b = &EncoderProcessor{}
	default:
		b = &ProcessorNodeBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalProcessorNodeBaseClassificationArray(rawMsg json.RawMessage) ([]ProcessorNodeBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ProcessorNodeBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalProcessorNodeBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalSinkNodeBaseClassification(rawMsg json.RawMessage) (SinkNodeBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SinkNodeBaseClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.VideoSink":
		b = &VideoSink{}
	default:
		b = &SinkNodeBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSinkNodeBaseClassificationArray(rawMsg json.RawMessage) ([]SinkNodeBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]SinkNodeBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalSinkNodeBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalSourceNodeBaseClassification(rawMsg json.RawMessage) (SourceNodeBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SourceNodeBaseClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.RtspSource":
		b = &RtspSource{}
	case "#Microsoft.VideoAnalyzer.VideoSource":
		b = &VideoSource{}
	default:
		b = &SourceNodeBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSourceNodeBaseClassificationArray(rawMsg json.RawMessage) ([]SourceNodeBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]SourceNodeBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalSourceNodeBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalTimeSequenceBaseClassification(rawMsg json.RawMessage) (TimeSequenceBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TimeSequenceBaseClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.VideoSequenceAbsoluteTimeMarkers":
		b = &VideoSequenceAbsoluteTimeMarkers{}
	default:
		b = &TimeSequenceBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTokenKeyClassification(rawMsg json.RawMessage) (TokenKeyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TokenKeyClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.EccTokenKey":
		b = &EccTokenKey{}
	case "#Microsoft.VideoAnalyzer.RsaTokenKey":
		b = &RsaTokenKey{}
	default:
		b = &TokenKey{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTokenKeyClassificationArray(rawMsg json.RawMessage) ([]TokenKeyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]TokenKeyClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalTokenKeyClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalTunnelBaseClassification(rawMsg json.RawMessage) (TunnelBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TunnelBaseClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.SecureIotDeviceRemoteTunnel":
		b = &SecureIotDeviceRemoteTunnel{}
	default:
		b = &TunnelBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalVideoEncoderBaseClassification(rawMsg json.RawMessage) (VideoEncoderBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b VideoEncoderBaseClassification
	switch m["@type"] {
	case "#Microsoft.VideoAnalyzer.VideoEncoderH264":
		b = &VideoEncoderH264{}
	default:
		b = &VideoEncoderBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
