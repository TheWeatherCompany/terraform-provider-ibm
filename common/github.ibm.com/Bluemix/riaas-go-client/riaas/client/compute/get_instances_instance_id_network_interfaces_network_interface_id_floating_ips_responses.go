// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsReader is a Reader for the GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIps structure.
type GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOK creates a GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOK with default headers values
func NewGetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOK() *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOK {
	return &GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOK{}
}

/*GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOK handles this case with default header values.

The associated floating IPs were retrieved successfully.
*/
type GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOK struct {
	Payload *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBody
}

func (o *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOK) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips][%d] getInstancesInstanceIdNetworkInterfacesNetworkInterfaceIdFloatingIpsOK  %+v", 200, o.Payload)
}

func (o *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsNotFound creates a GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsNotFound with default headers values
func NewGetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsNotFound() *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsNotFound {
	return &GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsNotFound{}
}

/*GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsNotFound handles this case with default header values.

A network interface with the specified identifier could not be found.
*/
type GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips][%d] getInstancesInstanceIdNetworkInterfacesNetworkInterfaceIdFloatingIpsNotFound  %+v", 404, o.Payload)
}

func (o *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBody UnpaginatedFloatingIPCollection
//
// Collection of floating IPs
swagger:model GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBody
*/
type GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBody struct {

	// Collection of floating IPs
	FloatingIps []*models.FloatingIP `json:"floating_ips,omitempty"`
}

// Validate validates this get instances instance ID network interfaces network interface ID floating ips o k body
func (o *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFloatingIps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBody) validateFloatingIps(formats strfmt.Registry) error {

	if swag.IsZero(o.FloatingIps) { // not required
		return nil
	}

	for i := 0; i < len(o.FloatingIps); i++ {
		if swag.IsZero(o.FloatingIps[i]) { // not required
			continue
		}

		if o.FloatingIps[i] != nil {
			if err := o.FloatingIps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getInstancesInstanceIdNetworkInterfacesNetworkInterfaceIdFloatingIpsOK" + "." + "floating_ips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBody) UnmarshalBinary(b []byte) error {
	var res GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}