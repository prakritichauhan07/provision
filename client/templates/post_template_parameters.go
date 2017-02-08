package templates

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

	"github.com/rackn/rocket-skates/models"
)

// NewPostTemplateParams creates a new PostTemplateParams object
// with the default values initialized.
func NewPostTemplateParams() *PostTemplateParams {
	var ()
	return &PostTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTemplateParamsWithTimeout creates a new PostTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTemplateParamsWithTimeout(timeout time.Duration) *PostTemplateParams {
	var ()
	return &PostTemplateParams{

		timeout: timeout,
	}
}

// NewPostTemplateParamsWithContext creates a new PostTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTemplateParamsWithContext(ctx context.Context) *PostTemplateParams {
	var ()
	return &PostTemplateParams{

		Context: ctx,
	}
}

/*PostTemplateParams contains all the parameters to send to the API endpoint
for the post template operation typically these are written to a http.Request
*/
type PostTemplateParams struct {

	/*Body*/
	Body *models.TemplateInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post template params
func (o *PostTemplateParams) WithTimeout(timeout time.Duration) *PostTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post template params
func (o *PostTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post template params
func (o *PostTemplateParams) WithContext(ctx context.Context) *PostTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post template params
func (o *PostTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post template params
func (o *PostTemplateParams) WithBody(body *models.TemplateInput) *PostTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post template params
func (o *PostTemplateParams) SetBody(body *models.TemplateInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.TemplateInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
