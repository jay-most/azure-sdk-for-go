//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LogProfilesClient contains the methods for the LogProfiles group.
// Don't use this type directly, use NewLogProfilesClient() instead.
type LogProfilesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLogProfilesClient creates a new instance of LogProfilesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLogProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LogProfilesClient, error) {
	cl, err := arm.NewClient(moduleName+".LogProfilesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LogProfilesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a log profile in Azure Monitoring REST API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
//   - logProfileName - The name of the log profile.
//   - parameters - Parameters supplied to the operation.
//   - options - LogProfilesClientCreateOrUpdateOptions contains the optional parameters for the LogProfilesClient.CreateOrUpdate
//     method.
func (client *LogProfilesClient) CreateOrUpdate(ctx context.Context, logProfileName string, parameters LogProfileResource, options *LogProfilesClientCreateOrUpdateOptions) (LogProfilesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "LogProfilesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, logProfileName, parameters, options)
	if err != nil {
		return LogProfilesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogProfilesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogProfilesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LogProfilesClient) createOrUpdateCreateRequest(ctx context.Context, logProfileName string, parameters LogProfileResource, options *LogProfilesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles/{logProfileName}"
	if logProfileName == "" {
		return nil, errors.New("parameter logProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logProfileName}", url.PathEscape(logProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LogProfilesClient) createOrUpdateHandleResponse(resp *http.Response) (LogProfilesClientCreateOrUpdateResponse, error) {
	result := LogProfilesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogProfileResource); err != nil {
		return LogProfilesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the log profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
//   - logProfileName - The name of the log profile.
//   - options - LogProfilesClientDeleteOptions contains the optional parameters for the LogProfilesClient.Delete method.
func (client *LogProfilesClient) Delete(ctx context.Context, logProfileName string, options *LogProfilesClientDeleteOptions) (LogProfilesClientDeleteResponse, error) {
	var err error
	const operationName = "LogProfilesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, logProfileName, options)
	if err != nil {
		return LogProfilesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogProfilesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogProfilesClientDeleteResponse{}, err
	}
	return LogProfilesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LogProfilesClient) deleteCreateRequest(ctx context.Context, logProfileName string, options *LogProfilesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles/{logProfileName}"
	if logProfileName == "" {
		return nil, errors.New("parameter logProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logProfileName}", url.PathEscape(logProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the log profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
//   - logProfileName - The name of the log profile.
//   - options - LogProfilesClientGetOptions contains the optional parameters for the LogProfilesClient.Get method.
func (client *LogProfilesClient) Get(ctx context.Context, logProfileName string, options *LogProfilesClientGetOptions) (LogProfilesClientGetResponse, error) {
	var err error
	const operationName = "LogProfilesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, logProfileName, options)
	if err != nil {
		return LogProfilesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogProfilesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LogProfilesClient) getCreateRequest(ctx context.Context, logProfileName string, options *LogProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles/{logProfileName}"
	if logProfileName == "" {
		return nil, errors.New("parameter logProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logProfileName}", url.PathEscape(logProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LogProfilesClient) getHandleResponse(resp *http.Response) (LogProfilesClientGetResponse, error) {
	result := LogProfilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogProfileResource); err != nil {
		return LogProfilesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the log profiles.
//
// Generated from API version 2016-03-01
//   - options - LogProfilesClientListOptions contains the optional parameters for the LogProfilesClient.NewListPager method.
func (client *LogProfilesClient) NewListPager(options *LogProfilesClientListOptions) *runtime.Pager[LogProfilesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LogProfilesClientListResponse]{
		More: func(page LogProfilesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *LogProfilesClientListResponse) (LogProfilesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LogProfilesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, options)
			if err != nil {
				return LogProfilesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LogProfilesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LogProfilesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *LogProfilesClient) listCreateRequest(ctx context.Context, options *LogProfilesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LogProfilesClient) listHandleResponse(resp *http.Response) (LogProfilesClientListResponse, error) {
	result := LogProfilesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogProfileCollection); err != nil {
		return LogProfilesClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing LogProfilesResource. To update other fields use the CreateOrUpdate method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
//   - logProfileName - The name of the log profile.
//   - logProfilesResource - Parameters supplied to the operation.
//   - options - LogProfilesClientUpdateOptions contains the optional parameters for the LogProfilesClient.Update method.
func (client *LogProfilesClient) Update(ctx context.Context, logProfileName string, logProfilesResource LogProfileResourcePatch, options *LogProfilesClientUpdateOptions) (LogProfilesClientUpdateResponse, error) {
	var err error
	const operationName = "LogProfilesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, logProfileName, logProfilesResource, options)
	if err != nil {
		return LogProfilesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogProfilesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogProfilesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *LogProfilesClient) updateCreateRequest(ctx context.Context, logProfileName string, logProfilesResource LogProfileResourcePatch, options *LogProfilesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles/{logProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if logProfileName == "" {
		return nil, errors.New("parameter logProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logProfileName}", url.PathEscape(logProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, logProfilesResource); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *LogProfilesClient) updateHandleResponse(resp *http.Response) (LogProfilesClientUpdateResponse, error) {
	result := LogProfilesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogProfileResource); err != nil {
		return LogProfilesClientUpdateResponse{}, err
	}
	return result, nil
}
