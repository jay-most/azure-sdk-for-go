//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkScopesServer is a fake server for instances of the armmonitor.PrivateLinkScopesClient type.
type PrivateLinkScopesServer struct {
	// CreateOrUpdate is the fake for method PrivateLinkScopesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, scopeName string, azureMonitorPrivateLinkScopePayload armmonitor.AzureMonitorPrivateLinkScope, options *armmonitor.PrivateLinkScopesClientCreateOrUpdateOptions) (resp azfake.Responder[armmonitor.PrivateLinkScopesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateLinkScopesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, scopeName string, options *armmonitor.PrivateLinkScopesClientBeginDeleteOptions) (resp azfake.PollerResponder[armmonitor.PrivateLinkScopesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateLinkScopesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, scopeName string, options *armmonitor.PrivateLinkScopesClientGetOptions) (resp azfake.Responder[armmonitor.PrivateLinkScopesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PrivateLinkScopesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armmonitor.PrivateLinkScopesClientListOptions) (resp azfake.PagerResponder[armmonitor.PrivateLinkScopesClientListResponse])

	// NewListByResourceGroupPager is the fake for method PrivateLinkScopesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmonitor.PrivateLinkScopesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmonitor.PrivateLinkScopesClientListByResourceGroupResponse])

	// UpdateTags is the fake for method PrivateLinkScopesClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, scopeName string, privateLinkScopeTags armmonitor.TagsResource, options *armmonitor.PrivateLinkScopesClientUpdateTagsOptions) (resp azfake.Responder[armmonitor.PrivateLinkScopesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewPrivateLinkScopesServerTransport creates a new instance of PrivateLinkScopesServerTransport with the provided implementation.
// The returned PrivateLinkScopesServerTransport instance is connected to an instance of armmonitor.PrivateLinkScopesClient by way of the
// undefined.Transporter field.
func NewPrivateLinkScopesServerTransport(srv *PrivateLinkScopesServer) *PrivateLinkScopesServerTransport {
	return &PrivateLinkScopesServerTransport{srv: srv}
}

// PrivateLinkScopesServerTransport connects instances of armmonitor.PrivateLinkScopesClient to instances of PrivateLinkScopesServer.
// Don't use this type directly, use NewPrivateLinkScopesServerTransport instead.
type PrivateLinkScopesServerTransport struct {
	srv                         *PrivateLinkScopesServer
	beginDelete                 *azfake.PollerResponder[armmonitor.PrivateLinkScopesClientDeleteResponse]
	newListPager                *azfake.PagerResponder[armmonitor.PrivateLinkScopesClientListResponse]
	newListByResourceGroupPager *azfake.PagerResponder[armmonitor.PrivateLinkScopesClientListByResourceGroupResponse]
}

// Do implements the policy.Transporter interface for PrivateLinkScopesServerTransport.
func (p *PrivateLinkScopesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateLinkScopesClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "PrivateLinkScopesClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PrivateLinkScopesClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateLinkScopesClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PrivateLinkScopesClient.NewListByResourceGroupPager":
		resp, err = p.dispatchNewListByResourceGroupPager(req)
	case "PrivateLinkScopesClient.UpdateTags":
		resp, err = p.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateLinkScopesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/microsoft.insights/privateLinkScopes/(?P<scopeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.AzureMonitorPrivateLinkScope](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	scopeNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("scopeName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), resourceGroupNameUnescaped, scopeNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureMonitorPrivateLinkScope, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateLinkScopesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if p.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/microsoft.insights/privateLinkScopes/(?P<scopeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		scopeNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("scopeName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, scopeNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		p.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(p.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(p.beginDelete) {
		p.beginDelete = nil
	}

	return resp, nil
}

func (p *PrivateLinkScopesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/microsoft.insights/privateLinkScopes/(?P<scopeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	scopeNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("scopeName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameUnescaped, scopeNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureMonitorPrivateLinkScope, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateLinkScopesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if p.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/microsoft.insights/privateLinkScopes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListPager(nil)
		p.newListPager = &resp
		server.PagerResponderInjectNextLinks(p.newListPager, req, func(page *armmonitor.PrivateLinkScopesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListPager) {
		p.newListPager = nil
	}
	return resp, nil
}

func (p *PrivateLinkScopesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	if p.newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/microsoft.insights/privateLinkScopes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		p.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(p.newListByResourceGroupPager, req, func(page *armmonitor.PrivateLinkScopesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(p.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(p.newListByResourceGroupPager) {
		p.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (p *PrivateLinkScopesServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if p.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/microsoft.insights/privateLinkScopes/(?P<scopeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.TagsResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	scopeNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("scopeName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, scopeNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureMonitorPrivateLinkScope, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
