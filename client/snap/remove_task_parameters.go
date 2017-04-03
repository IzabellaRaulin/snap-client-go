package snap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRemoveTaskParams creates a new RemoveTaskParams object
// with the default values initialized.
func NewRemoveTaskParams() *RemoveTaskParams {
	var ()
	return &RemoveTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveTaskParamsWithTimeout creates a new RemoveTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveTaskParamsWithTimeout(timeout time.Duration) *RemoveTaskParams {
	var ()
	return &RemoveTaskParams{

		timeout: timeout,
	}
}

// NewRemoveTaskParamsWithContext creates a new RemoveTaskParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveTaskParamsWithContext(ctx context.Context) *RemoveTaskParams {
	var ()
	return &RemoveTaskParams{

		Context: ctx,
	}
}

// NewRemoveTaskParamsWithHTTPClient creates a new RemoveTaskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveTaskParamsWithHTTPClient(client *http.Client) *RemoveTaskParams {
	var ()
	return &RemoveTaskParams{
		HTTPClient: client,
	}
}

/*RemoveTaskParams contains all the parameters to send to the API endpoint
for the remove task operation typically these are written to a http.Request
*/
type RemoveTaskParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove task params
func (o *RemoveTaskParams) WithTimeout(timeout time.Duration) *RemoveTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove task params
func (o *RemoveTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove task params
func (o *RemoveTaskParams) WithContext(ctx context.Context) *RemoveTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove task params
func (o *RemoveTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove task params
func (o *RemoveTaskParams) WithHTTPClient(client *http.Client) *RemoveTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove task params
func (o *RemoveTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the remove task params
func (o *RemoveTaskParams) WithID(id string) *RemoveTaskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove task params
func (o *RemoveTaskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}