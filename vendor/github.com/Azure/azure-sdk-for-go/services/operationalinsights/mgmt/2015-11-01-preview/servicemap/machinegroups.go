package servicemap

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// MachineGroupsClient is the service Map API Reference
type MachineGroupsClient struct {
	BaseClient
}

// NewMachineGroupsClient creates an instance of the MachineGroupsClient client.
func NewMachineGroupsClient(subscriptionID string) MachineGroupsClient {
	return NewMachineGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMachineGroupsClientWithBaseURI creates an instance of the MachineGroupsClient client.
func NewMachineGroupsClientWithBaseURI(baseURI string, subscriptionID string) MachineGroupsClient {
	return MachineGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new machine group.
//
// resourceGroupName is resource group name within the specified subscriptionId. workspaceName is OMS workspace
// containing the resources of interest. machineGroup is machine Group resource to create.
func (client MachineGroupsClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, machineGroup MachineGroup) (result MachineGroup, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineGroup,
			Constraints: []validation.Constraint{{Target: "machineGroup.MachineGroupProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "machineGroup.MachineGroupProperties.DisplayName", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "machineGroup.MachineGroupProperties.DisplayName", Name: validation.MaxLength, Rule: 256, Chain: nil},
						{Target: "machineGroup.MachineGroupProperties.DisplayName", Name: validation.MinLength, Rule: 1, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("servicemap.MachineGroupsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, workspaceName, machineGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client MachineGroupsClient) CreatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineGroup MachineGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machineGroups", pathParameters),
		autorest.WithJSON(machineGroup),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client MachineGroupsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client MachineGroupsClient) CreateResponder(resp *http.Response) (result MachineGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified Machine Group.
//
// resourceGroupName is resource group name within the specified subscriptionId. workspaceName is OMS workspace
// containing the resources of interest. machineGroupName is machine Group resource name.
func (client MachineGroupsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, machineGroupName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineGroupName,
			Constraints: []validation.Constraint{{Target: "machineGroupName", Name: validation.MaxLength, Rule: 36, Chain: nil},
				{Target: "machineGroupName", Name: validation.MinLength, Rule: 36, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.MachineGroupsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, machineGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MachineGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineGroupName":  autorest.Encode("path", machineGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machineGroups/{machineGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MachineGroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MachineGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns the specified machine group.
//
// resourceGroupName is resource group name within the specified subscriptionId. workspaceName is OMS workspace
// containing the resources of interest. machineGroupName is machine Group resource name.
func (client MachineGroupsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, machineGroupName string) (result MachineGroup, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineGroupName,
			Constraints: []validation.Constraint{{Target: "machineGroupName", Name: validation.MaxLength, Rule: 36, Chain: nil},
				{Target: "machineGroupName", Name: validation.MinLength, Rule: 36, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.MachineGroupsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, machineGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MachineGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineGroupName":  autorest.Encode("path", machineGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machineGroups/{machineGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MachineGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MachineGroupsClient) GetResponder(resp *http.Response) (result MachineGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByWorkspace returns all machine groups.
//
// resourceGroupName is resource group name within the specified subscriptionId. workspaceName is OMS workspace
// containing the resources of interest.
func (client MachineGroupsClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result MachineGroupCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.MachineGroupsClient", "ListByWorkspace", err.Error())
	}

	result.fn = client.listByWorkspaceNextResults
	req, err := client.ListByWorkspacePreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "ListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.mgc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "ListByWorkspace", resp, "Failure sending request")
		return
	}

	result.mgc, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "ListByWorkspace", resp, "Failure responding to request")
	}

	return
}

// ListByWorkspacePreparer prepares the ListByWorkspace request.
func (client MachineGroupsClient) ListByWorkspacePreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machineGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkspaceSender sends the ListByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client MachineGroupsClient) ListByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkspaceResponder handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (client MachineGroupsClient) ListByWorkspaceResponder(resp *http.Response) (result MachineGroupCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByWorkspaceNextResults retrieves the next set of results, if any.
func (client MachineGroupsClient) listByWorkspaceNextResults(lastResults MachineGroupCollection) (result MachineGroupCollection, err error) {
	req, err := lastResults.machineGroupCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "listByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "listByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "listByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client MachineGroupsClient) ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result MachineGroupCollectionIterator, err error) {
	result.page, err = client.ListByWorkspace(ctx, resourceGroupName, workspaceName)
	return
}

// Update updates a machine group.
//
// resourceGroupName is resource group name within the specified subscriptionId. workspaceName is OMS workspace
// containing the resources of interest. machineGroupName is machine Group resource name. machineGroup is machine
// Group resource to update.
func (client MachineGroupsClient) Update(ctx context.Context, resourceGroupName string, workspaceName string, machineGroupName string, machineGroup MachineGroup) (result MachineGroup, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineGroupName,
			Constraints: []validation.Constraint{{Target: "machineGroupName", Name: validation.MaxLength, Rule: 36, Chain: nil},
				{Target: "machineGroupName", Name: validation.MinLength, Rule: 36, Chain: nil}}},
		{TargetValue: machineGroup,
			Constraints: []validation.Constraint{{Target: "machineGroup.MachineGroupProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "machineGroup.MachineGroupProperties.DisplayName", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "machineGroup.MachineGroupProperties.DisplayName", Name: validation.MaxLength, Rule: 256, Chain: nil},
						{Target: "machineGroup.MachineGroupProperties.DisplayName", Name: validation.MinLength, Rule: 1, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("servicemap.MachineGroupsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, workspaceName, machineGroupName, machineGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MachineGroupsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client MachineGroupsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineGroupName string, machineGroup MachineGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineGroupName":  autorest.Encode("path", machineGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machineGroups/{machineGroupName}", pathParameters),
		autorest.WithJSON(machineGroup),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client MachineGroupsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client MachineGroupsClient) UpdateResponder(resp *http.Response) (result MachineGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
