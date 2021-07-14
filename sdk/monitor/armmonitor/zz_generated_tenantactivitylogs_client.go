// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// TenantActivityLogsClient contains the methods for the TenantActivityLogs group.
// Don't use this type directly, use NewTenantActivityLogsClient() instead.
type TenantActivityLogsClient struct {
	con *armcore.Connection
}

// NewTenantActivityLogsClient creates a new instance of TenantActivityLogsClient with the specified values.
func NewTenantActivityLogsClient(con *armcore.Connection) *TenantActivityLogsClient {
	return &TenantActivityLogsClient{con: con}
}

// List - Gets the Activity Logs for the Tenant. Everything that is applicable to the API to get the Activity Logs for the subscription is applicable to
// this API (the parameters, $filter, etc.). One thing to
// point out here is that this API does not retrieve the logs at the individual subscription of the tenant but only surfaces the logs that were generated
// at the tenant level.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TenantActivityLogsClient) List(options *TenantActivityLogsListOptions) EventDataCollectionPager {
	return &eventDataCollectionPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp EventDataCollectionResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.EventDataCollection.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *TenantActivityLogsClient) listCreateRequest(ctx context.Context, options *TenantActivityLogsListOptions) (*azcore.Request, error) {
	urlPath := "/providers/Microsoft.Insights/eventtypes/management/values"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2015-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TenantActivityLogsClient) listHandleResponse(resp *azcore.Response) (EventDataCollectionResponse, error) {
	var val *EventDataCollection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return EventDataCollectionResponse{}, err
	}
	return EventDataCollectionResponse{RawResponse: resp.Response, EventDataCollection: val}, nil
}

// listHandleError handles the List error response.
func (client *TenantActivityLogsClient) listHandleError(resp *azcore.Response) error {
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
