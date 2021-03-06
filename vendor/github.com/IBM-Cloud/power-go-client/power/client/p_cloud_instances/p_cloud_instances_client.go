// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new p cloud instances API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p cloud instances API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PcloudCloudinstancesDelete deletes a power cloud instance
*/
func (a *Client) PcloudCloudinstancesDelete(params *PcloudCloudinstancesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudCloudinstancesDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudinstancesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.cloudinstances.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudinstancesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudCloudinstancesDeleteOK), nil

}

/*
PcloudCloudinstancesGet gets a cloud instance s current state information
*/
func (a *Client) PcloudCloudinstancesGet(params *PcloudCloudinstancesGetParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudCloudinstancesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudinstancesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.cloudinstances.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudinstancesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudCloudinstancesGetOK), nil

}

/*
PcloudCloudinstancesPut updates upgrade a cloud instance
*/
func (a *Client) PcloudCloudinstancesPut(params *PcloudCloudinstancesPutParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudCloudinstancesPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudinstancesPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.cloudinstances.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudinstancesPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudCloudinstancesPutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
