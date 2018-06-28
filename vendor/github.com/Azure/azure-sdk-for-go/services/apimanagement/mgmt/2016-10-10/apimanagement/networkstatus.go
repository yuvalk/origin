package apimanagement

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

// NetworkStatusClient is the apiManagement Client
type NetworkStatusClient struct {
	BaseClient
}

// NewNetworkStatusClient creates an instance of the NetworkStatusClient client.
func NewNetworkStatusClient(subscriptionID string) NetworkStatusClient {
	return NewNetworkStatusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNetworkStatusClientWithBaseURI creates an instance of the NetworkStatusClient client.
func NewNetworkStatusClientWithBaseURI(baseURI string, subscriptionID string) NetworkStatusClient {
	return NetworkStatusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByService gets the Connectivity Status to the external resources on which the Api Management service depends from
// inside the Cloud Service. This also returns the DNS Servers as visible to the CloudService.
//
// resourceGroupName is the name of the resource group. serviceName is the name of the API Management service.
func (client NetworkStatusClient) GetByService(ctx context.Context, resourceGroupName string, serviceName string) (result NetworkStatusContract, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.NetworkStatusClient", "GetByService", err.Error())
	}

	req, err := client.GetByServicePreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NetworkStatusClient", "GetByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.NetworkStatusClient", "GetByService", resp, "Failure sending request")
		return
	}

	result, err = client.GetByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NetworkStatusClient", "GetByService", resp, "Failure responding to request")
	}

	return
}

// GetByServicePreparer prepares the GetByService request.
func (client NetworkStatusClient) GetByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/networkstatus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByServiceSender sends the GetByService request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkStatusClient) GetByServiceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetByServiceResponder handles the response to the GetByService request. The method always
// closes the http.Response Body.
func (client NetworkStatusClient) GetByServiceResponder(resp *http.Response) (result NetworkStatusContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
