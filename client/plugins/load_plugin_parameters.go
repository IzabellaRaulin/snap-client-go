// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLoadPluginParams creates a new LoadPluginParams object
// with the default values initialized.
func NewLoadPluginParams() *LoadPluginParams {
	var ()
	return &LoadPluginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadPluginParamsWithTimeout creates a new LoadPluginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadPluginParamsWithTimeout(timeout time.Duration) *LoadPluginParams {
	var ()
	return &LoadPluginParams{

		timeout: timeout,
	}
}

// NewLoadPluginParamsWithContext creates a new LoadPluginParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadPluginParamsWithContext(ctx context.Context) *LoadPluginParams {
	var ()
	return &LoadPluginParams{

		Context: ctx,
	}
}

// NewLoadPluginParamsWithHTTPClient creates a new LoadPluginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoadPluginParamsWithHTTPClient(client *http.Client) *LoadPluginParams {
	var ()
	return &LoadPluginParams{
		HTTPClient: client,
	}
}

/*LoadPluginParams contains all the parameters to send to the API endpoint
for the load plugin operation typically these are written to a http.Request
*/
type LoadPluginParams struct {

	/*CaCerts
	  CA root certification

	*/
	CaCerts *string
	/*PluginCert
	  Plugin GRPC TLS server certification

	*/
	PluginCert *string
	/*PluginData
	  loads a plugin.

	*/
	PluginData *os.File
	/*PluginKey
	  Plugin GRPC TLS server key

	*/
	PluginKey *string
	/*PluginURI
	  Stand-alone plugin URI

	*/
	PluginURI *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the load plugin params
func (o *LoadPluginParams) WithTimeout(timeout time.Duration) *LoadPluginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load plugin params
func (o *LoadPluginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load plugin params
func (o *LoadPluginParams) WithContext(ctx context.Context) *LoadPluginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load plugin params
func (o *LoadPluginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the load plugin params
func (o *LoadPluginParams) WithHTTPClient(client *http.Client) *LoadPluginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the load plugin params
func (o *LoadPluginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaCerts adds the caCerts to the load plugin params
func (o *LoadPluginParams) WithCaCerts(caCerts *string) *LoadPluginParams {
	o.SetCaCerts(caCerts)
	return o
}

// SetCaCerts adds the caCerts to the load plugin params
func (o *LoadPluginParams) SetCaCerts(caCerts *string) {
	o.CaCerts = caCerts
}

// WithPluginCert adds the pluginCert to the load plugin params
func (o *LoadPluginParams) WithPluginCert(pluginCert *string) *LoadPluginParams {
	o.SetPluginCert(pluginCert)
	return o
}

// SetPluginCert adds the pluginCert to the load plugin params
func (o *LoadPluginParams) SetPluginCert(pluginCert *string) {
	o.PluginCert = pluginCert
}

// WithPluginData adds the pluginData to the load plugin params
func (o *LoadPluginParams) WithPluginData(pluginData *os.File) *LoadPluginParams {
	o.SetPluginData(pluginData)
	return o
}

// SetPluginData adds the pluginData to the load plugin params
func (o *LoadPluginParams) SetPluginData(pluginData *os.File) {
	o.PluginData = pluginData
}

// WithPluginKey adds the pluginKey to the load plugin params
func (o *LoadPluginParams) WithPluginKey(pluginKey *string) *LoadPluginParams {
	o.SetPluginKey(pluginKey)
	return o
}

// SetPluginKey adds the pluginKey to the load plugin params
func (o *LoadPluginParams) SetPluginKey(pluginKey *string) {
	o.PluginKey = pluginKey
}

// WithPluginURI adds the pluginURI to the load plugin params
func (o *LoadPluginParams) WithPluginURI(pluginURI *string) *LoadPluginParams {
	o.SetPluginURI(pluginURI)
	return o
}

// SetPluginURI adds the pluginUri to the load plugin params
func (o *LoadPluginParams) SetPluginURI(pluginURI *string) {
	o.PluginURI = pluginURI
}

// WriteToRequest writes these params to a swagger request
func (o *LoadPluginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CaCerts != nil {

		// form param ca_certs
		var frCaCerts string
		if o.CaCerts != nil {
			frCaCerts = *o.CaCerts
		}
		fCaCerts := frCaCerts
		if fCaCerts != "" {
			if err := r.SetFormParam("ca_certs", fCaCerts); err != nil {
				return err
			}
		}

	}

	if o.PluginCert != nil {

		// form param plugin_cert
		var frPluginCert string
		if o.PluginCert != nil {
			frPluginCert = *o.PluginCert
		}
		fPluginCert := frPluginCert
		if fPluginCert != "" {
			if err := r.SetFormParam("plugin_cert", fPluginCert); err != nil {
				return err
			}
		}

	}

	if o.PluginData != nil {

		if o.PluginData != nil {

			// form file param plugin_data
			if err := r.SetFileParam("plugin_data", o.PluginData); err != nil {
				return err
			}

		}

	}

	if o.PluginKey != nil {

		// form param plugin_key
		var frPluginKey string
		if o.PluginKey != nil {
			frPluginKey = *o.PluginKey
		}
		fPluginKey := frPluginKey
		if fPluginKey != "" {
			if err := r.SetFormParam("plugin_key", fPluginKey); err != nil {
				return err
			}
		}

	}

	if o.PluginURI != nil {

		// form param plugin_uri
		var frPluginURI string
		if o.PluginURI != nil {
			frPluginURI = *o.PluginURI
		}
		fPluginURI := frPluginURI
		if fPluginURI != "" {
			if err := r.SetFormParam("plugin_uri", fPluginURI); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
