// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// ActionGroupsClient contains the methods for the ActionGroups group.
// Don't use this type directly, use NewActionGroupsClient() instead.
type ActionGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewActionGroupsClient creates a new instance of ActionGroupsClient with the specified values.
func NewActionGroupsClient(con *armcore.Connection, subscriptionID string) *ActionGroupsClient {
	return &ActionGroupsClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create a new action group or update an existing one.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActionGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroup ActionGroupResource, options *ActionGroupsCreateOrUpdateOptions) (ActionGroupResourceResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, actionGroupName, actionGroup, options)
	if err != nil {
		return ActionGroupResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ActionGroupResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return ActionGroupResourceResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ActionGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroup ActionGroupResource, options *ActionGroupsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if actionGroupName == "" {
		return nil, errors.New("parameter actionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionGroupName}", url.PathEscape(actionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(actionGroup)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ActionGroupsClient) createOrUpdateHandleResponse(resp *azcore.Response) (ActionGroupResourceResponse, error) {
	var val *ActionGroupResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ActionGroupResourceResponse{}, err
	}
	return ActionGroupResourceResponse{RawResponse: resp.Response, ActionGroupResource: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ActionGroupsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Delete an action group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActionGroupsClient) Delete(ctx context.Context, resourceGroupName string, actionGroupName string, options *ActionGroupsDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, actionGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ActionGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, actionGroupName string, options *ActionGroupsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if actionGroupName == "" {
		return nil, errors.New("parameter actionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionGroupName}", url.PathEscape(actionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ActionGroupsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// EnableReceiver - Enable a receiver in an action group. This changes the receiver's status from Disabled to Enabled. This operation is only supported
// for Email or SMS receivers.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActionGroupsClient) EnableReceiver(ctx context.Context, resourceGroupName string, actionGroupName string, enableRequest EnableRequest, options *ActionGroupsEnableReceiverOptions) (*http.Response, error) {
	req, err := client.enableReceiverCreateRequest(ctx, resourceGroupName, actionGroupName, enableRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusConflict) {
		return nil, client.enableReceiverHandleError(resp)
	}
	return resp.Response, nil
}

// enableReceiverCreateRequest creates the EnableReceiver request.
func (client *ActionGroupsClient) enableReceiverCreateRequest(ctx context.Context, resourceGroupName string, actionGroupName string, enableRequest EnableRequest, options *ActionGroupsEnableReceiverOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}/subscribe"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if actionGroupName == "" {
		return nil, errors.New("parameter actionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionGroupName}", url.PathEscape(actionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(enableRequest)
}

// enableReceiverHandleError handles the EnableReceiver error response.
func (client *ActionGroupsClient) enableReceiverHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Get an action group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActionGroupsClient) Get(ctx context.Context, resourceGroupName string, actionGroupName string, options *ActionGroupsGetOptions) (ActionGroupResourceResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, actionGroupName, options)
	if err != nil {
		return ActionGroupResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ActionGroupResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ActionGroupResourceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ActionGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, actionGroupName string, options *ActionGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if actionGroupName == "" {
		return nil, errors.New("parameter actionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionGroupName}", url.PathEscape(actionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ActionGroupsClient) getHandleResponse(resp *azcore.Response) (ActionGroupResourceResponse, error) {
	var val *ActionGroupResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ActionGroupResourceResponse{}, err
	}
	return ActionGroupResourceResponse{RawResponse: resp.Response, ActionGroupResource: val}, nil
}

// getHandleError handles the Get error response.
func (client *ActionGroupsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByResourceGroup - Get a list of all action groups in a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActionGroupsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *ActionGroupsListByResourceGroupOptions) (ActionGroupListResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ActionGroupListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ActionGroupListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ActionGroupListResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ActionGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ActionGroupsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ActionGroupsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ActionGroupListResponse, error) {
	var val *ActionGroupList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ActionGroupListResponse{}, err
	}
	return ActionGroupListResponse{RawResponse: resp.Response, ActionGroupList: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ActionGroupsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListBySubscriptionID - Get a list of all action groups in a subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActionGroupsClient) ListBySubscriptionID(ctx context.Context, options *ActionGroupsListBySubscriptionIDOptions) (ActionGroupListResponse, error) {
	req, err := client.listBySubscriptionIDCreateRequest(ctx, options)
	if err != nil {
		return ActionGroupListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ActionGroupListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ActionGroupListResponse{}, client.listBySubscriptionIDHandleError(resp)
	}
	return client.listBySubscriptionIDHandleResponse(resp)
}

// listBySubscriptionIDCreateRequest creates the ListBySubscriptionID request.
func (client *ActionGroupsClient) listBySubscriptionIDCreateRequest(ctx context.Context, options *ActionGroupsListBySubscriptionIDOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/microsoft.insights/actionGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionIDHandleResponse handles the ListBySubscriptionID response.
func (client *ActionGroupsClient) listBySubscriptionIDHandleResponse(resp *azcore.Response) (ActionGroupListResponse, error) {
	var val *ActionGroupList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ActionGroupListResponse{}, err
	}
	return ActionGroupListResponse{RawResponse: resp.Response, ActionGroupList: val}, nil
}

// listBySubscriptionIDHandleError handles the ListBySubscriptionID error response.
func (client *ActionGroupsClient) listBySubscriptionIDHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Update - Updates an existing action group's tags. To update other fields use the CreateOrUpdate method.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActionGroupsClient) Update(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroupPatch ActionGroupPatchBody, options *ActionGroupsUpdateOptions) (ActionGroupResourceResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, actionGroupName, actionGroupPatch, options)
	if err != nil {
		return ActionGroupResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ActionGroupResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ActionGroupResourceResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ActionGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroupPatch ActionGroupPatchBody, options *ActionGroupsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if actionGroupName == "" {
		return nil, errors.New("parameter actionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionGroupName}", url.PathEscape(actionGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(actionGroupPatch)
}

// updateHandleResponse handles the Update response.
func (client *ActionGroupsClient) updateHandleResponse(resp *azcore.Response) (ActionGroupResourceResponse, error) {
	var val *ActionGroupResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ActionGroupResourceResponse{}, err
	}
	return ActionGroupResourceResponse{RawResponse: resp.Response, ActionGroupResource: val}, nil
}

// updateHandleError handles the Update error response.
func (client *ActionGroupsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
