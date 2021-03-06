// Code generated by go-swagger; DO NOT EDIT.

package ticket_projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewTicketProjectsListParams creates a new TicketProjectsListParams object
// with the default values initialized.
func NewTicketProjectsListParams() *TicketProjectsListParams {
	var ()
	return &TicketProjectsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTicketProjectsListParamsWithTimeout creates a new TicketProjectsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTicketProjectsListParamsWithTimeout(timeout time.Duration) *TicketProjectsListParams {
	var ()
	return &TicketProjectsListParams{

		timeout: timeout,
	}
}

// NewTicketProjectsListParamsWithContext creates a new TicketProjectsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTicketProjectsListParamsWithContext(ctx context.Context) *TicketProjectsListParams {
	var ()
	return &TicketProjectsListParams{

		Context: ctx,
	}
}

// NewTicketProjectsListParamsWithHTTPClient creates a new TicketProjectsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTicketProjectsListParamsWithHTTPClient(client *http.Client) *TicketProjectsListParams {
	var ()
	return &TicketProjectsListParams{
		HTTPClient: client,
	}
}

/*TicketProjectsListParams contains all the parameters to send to the API endpoint
for the ticket projects list operation typically these are written to a http.Request
*/
type TicketProjectsListParams struct {

	/*AddedBy
	  Filter-traversable object

	*/
	AddedBy *string
	/*Filters
	  Complex filter that handles filtering with groups of conditions connected with AND and OR conjunctions. The query can be arbitrarily grouped and nested.

	*/
	Filters *string
	/*ID*/
	ID *float64
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *float64
	/*IntegrationName*/
	IntegrationName *string
	/*IntegrationNameIcontains*/
	IntegrationNameIcontains *string
	/*IntegrationNameStartswith*/
	IntegrationNameStartswith *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*NameIcontains*/
	NameIcontains *string
	/*NameStartswith*/
	NameStartswith *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Ordering
	  Valid ordering fields are: `id`, `parent_id`, `vendor_type`, `name`, `processing__updated_at`, `integration__name`, `processing__processing_status__id`

	*/
	Ordering *string
	/*Repo
	  Filter-traversable object

	*/
	Repo *string
	/*RepoID
	  Records matching the provided id

	*/
	RepoID *float64
	/*RepoIDIn
	  Match on repo_id in a comma separated list

	*/
	RepoIDIn *string
	/*Search
	  A search term.

	*/
	Search *string
	/*UpdatedAt
	  Records == date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	UpdatedAt *string
	/*UpdatedAtGt
	  Records > date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	UpdatedAtGt *string
	/*UpdatedAtGte
	  Records >= date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	UpdatedAtGte *string
	/*UpdatedAtLt
	  Records < date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	UpdatedAtLt *string
	/*UpdatedAtLte
	  Records <= date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	UpdatedAtLte *string
	/*URL*/
	URL *string
	/*URLIcontains*/
	URLIcontains *string
	/*URLStartswith*/
	URLStartswith *string
	/*VendorType*/
	VendorType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ticket projects list params
func (o *TicketProjectsListParams) WithTimeout(timeout time.Duration) *TicketProjectsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ticket projects list params
func (o *TicketProjectsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ticket projects list params
func (o *TicketProjectsListParams) WithContext(ctx context.Context) *TicketProjectsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ticket projects list params
func (o *TicketProjectsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ticket projects list params
func (o *TicketProjectsListParams) WithHTTPClient(client *http.Client) *TicketProjectsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ticket projects list params
func (o *TicketProjectsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddedBy adds the addedBy to the ticket projects list params
func (o *TicketProjectsListParams) WithAddedBy(addedBy *string) *TicketProjectsListParams {
	o.SetAddedBy(addedBy)
	return o
}

// SetAddedBy adds the addedBy to the ticket projects list params
func (o *TicketProjectsListParams) SetAddedBy(addedBy *string) {
	o.AddedBy = addedBy
}

// WithFilters adds the filters to the ticket projects list params
func (o *TicketProjectsListParams) WithFilters(filters *string) *TicketProjectsListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the ticket projects list params
func (o *TicketProjectsListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithID adds the id to the ticket projects list params
func (o *TicketProjectsListParams) WithID(id *float64) *TicketProjectsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ticket projects list params
func (o *TicketProjectsListParams) SetID(id *float64) {
	o.ID = id
}

// WithIDIn adds the iDIn to the ticket projects list params
func (o *TicketProjectsListParams) WithIDIn(iDIn *float64) *TicketProjectsListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the ticket projects list params
func (o *TicketProjectsListParams) SetIDIn(iDIn *float64) {
	o.IDIn = iDIn
}

// WithIntegrationName adds the integrationName to the ticket projects list params
func (o *TicketProjectsListParams) WithIntegrationName(integrationName *string) *TicketProjectsListParams {
	o.SetIntegrationName(integrationName)
	return o
}

// SetIntegrationName adds the integrationName to the ticket projects list params
func (o *TicketProjectsListParams) SetIntegrationName(integrationName *string) {
	o.IntegrationName = integrationName
}

// WithIntegrationNameIcontains adds the integrationNameIcontains to the ticket projects list params
func (o *TicketProjectsListParams) WithIntegrationNameIcontains(integrationNameIcontains *string) *TicketProjectsListParams {
	o.SetIntegrationNameIcontains(integrationNameIcontains)
	return o
}

// SetIntegrationNameIcontains adds the integrationNameIcontains to the ticket projects list params
func (o *TicketProjectsListParams) SetIntegrationNameIcontains(integrationNameIcontains *string) {
	o.IntegrationNameIcontains = integrationNameIcontains
}

// WithIntegrationNameStartswith adds the integrationNameStartswith to the ticket projects list params
func (o *TicketProjectsListParams) WithIntegrationNameStartswith(integrationNameStartswith *string) *TicketProjectsListParams {
	o.SetIntegrationNameStartswith(integrationNameStartswith)
	return o
}

// SetIntegrationNameStartswith adds the integrationNameStartswith to the ticket projects list params
func (o *TicketProjectsListParams) SetIntegrationNameStartswith(integrationNameStartswith *string) {
	o.IntegrationNameStartswith = integrationNameStartswith
}

// WithLimit adds the limit to the ticket projects list params
func (o *TicketProjectsListParams) WithLimit(limit *int64) *TicketProjectsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ticket projects list params
func (o *TicketProjectsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ticket projects list params
func (o *TicketProjectsListParams) WithName(name *string) *TicketProjectsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ticket projects list params
func (o *TicketProjectsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIcontains adds the nameIcontains to the ticket projects list params
func (o *TicketProjectsListParams) WithNameIcontains(nameIcontains *string) *TicketProjectsListParams {
	o.SetNameIcontains(nameIcontains)
	return o
}

// SetNameIcontains adds the nameIcontains to the ticket projects list params
func (o *TicketProjectsListParams) SetNameIcontains(nameIcontains *string) {
	o.NameIcontains = nameIcontains
}

// WithNameStartswith adds the nameStartswith to the ticket projects list params
func (o *TicketProjectsListParams) WithNameStartswith(nameStartswith *string) *TicketProjectsListParams {
	o.SetNameStartswith(nameStartswith)
	return o
}

// SetNameStartswith adds the nameStartswith to the ticket projects list params
func (o *TicketProjectsListParams) SetNameStartswith(nameStartswith *string) {
	o.NameStartswith = nameStartswith
}

// WithOffset adds the offset to the ticket projects list params
func (o *TicketProjectsListParams) WithOffset(offset *int64) *TicketProjectsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ticket projects list params
func (o *TicketProjectsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the ticket projects list params
func (o *TicketProjectsListParams) WithOrdering(ordering *string) *TicketProjectsListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the ticket projects list params
func (o *TicketProjectsListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithRepo adds the repo to the ticket projects list params
func (o *TicketProjectsListParams) WithRepo(repo *string) *TicketProjectsListParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the ticket projects list params
func (o *TicketProjectsListParams) SetRepo(repo *string) {
	o.Repo = repo
}

// WithRepoID adds the repoID to the ticket projects list params
func (o *TicketProjectsListParams) WithRepoID(repoID *float64) *TicketProjectsListParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the ticket projects list params
func (o *TicketProjectsListParams) SetRepoID(repoID *float64) {
	o.RepoID = repoID
}

// WithRepoIDIn adds the repoIDIn to the ticket projects list params
func (o *TicketProjectsListParams) WithRepoIDIn(repoIDIn *string) *TicketProjectsListParams {
	o.SetRepoIDIn(repoIDIn)
	return o
}

// SetRepoIDIn adds the repoIdIn to the ticket projects list params
func (o *TicketProjectsListParams) SetRepoIDIn(repoIDIn *string) {
	o.RepoIDIn = repoIDIn
}

// WithSearch adds the search to the ticket projects list params
func (o *TicketProjectsListParams) WithSearch(search *string) *TicketProjectsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the ticket projects list params
func (o *TicketProjectsListParams) SetSearch(search *string) {
	o.Search = search
}

// WithUpdatedAt adds the updatedAt to the ticket projects list params
func (o *TicketProjectsListParams) WithUpdatedAt(updatedAt *string) *TicketProjectsListParams {
	o.SetUpdatedAt(updatedAt)
	return o
}

// SetUpdatedAt adds the updatedAt to the ticket projects list params
func (o *TicketProjectsListParams) SetUpdatedAt(updatedAt *string) {
	o.UpdatedAt = updatedAt
}

// WithUpdatedAtGt adds the updatedAtGt to the ticket projects list params
func (o *TicketProjectsListParams) WithUpdatedAtGt(updatedAtGt *string) *TicketProjectsListParams {
	o.SetUpdatedAtGt(updatedAtGt)
	return o
}

// SetUpdatedAtGt adds the updatedAtGt to the ticket projects list params
func (o *TicketProjectsListParams) SetUpdatedAtGt(updatedAtGt *string) {
	o.UpdatedAtGt = updatedAtGt
}

// WithUpdatedAtGte adds the updatedAtGte to the ticket projects list params
func (o *TicketProjectsListParams) WithUpdatedAtGte(updatedAtGte *string) *TicketProjectsListParams {
	o.SetUpdatedAtGte(updatedAtGte)
	return o
}

// SetUpdatedAtGte adds the updatedAtGte to the ticket projects list params
func (o *TicketProjectsListParams) SetUpdatedAtGte(updatedAtGte *string) {
	o.UpdatedAtGte = updatedAtGte
}

// WithUpdatedAtLt adds the updatedAtLt to the ticket projects list params
func (o *TicketProjectsListParams) WithUpdatedAtLt(updatedAtLt *string) *TicketProjectsListParams {
	o.SetUpdatedAtLt(updatedAtLt)
	return o
}

// SetUpdatedAtLt adds the updatedAtLt to the ticket projects list params
func (o *TicketProjectsListParams) SetUpdatedAtLt(updatedAtLt *string) {
	o.UpdatedAtLt = updatedAtLt
}

// WithUpdatedAtLte adds the updatedAtLte to the ticket projects list params
func (o *TicketProjectsListParams) WithUpdatedAtLte(updatedAtLte *string) *TicketProjectsListParams {
	o.SetUpdatedAtLte(updatedAtLte)
	return o
}

// SetUpdatedAtLte adds the updatedAtLte to the ticket projects list params
func (o *TicketProjectsListParams) SetUpdatedAtLte(updatedAtLte *string) {
	o.UpdatedAtLte = updatedAtLte
}

// WithURL adds the url to the ticket projects list params
func (o *TicketProjectsListParams) WithURL(url *string) *TicketProjectsListParams {
	o.SetURL(url)
	return o
}

// SetURL adds the url to the ticket projects list params
func (o *TicketProjectsListParams) SetURL(url *string) {
	o.URL = url
}

// WithURLIcontains adds the uRLIcontains to the ticket projects list params
func (o *TicketProjectsListParams) WithURLIcontains(uRLIcontains *string) *TicketProjectsListParams {
	o.SetURLIcontains(uRLIcontains)
	return o
}

// SetURLIcontains adds the urlIcontains to the ticket projects list params
func (o *TicketProjectsListParams) SetURLIcontains(uRLIcontains *string) {
	o.URLIcontains = uRLIcontains
}

// WithURLStartswith adds the uRLStartswith to the ticket projects list params
func (o *TicketProjectsListParams) WithURLStartswith(uRLStartswith *string) *TicketProjectsListParams {
	o.SetURLStartswith(uRLStartswith)
	return o
}

// SetURLStartswith adds the urlStartswith to the ticket projects list params
func (o *TicketProjectsListParams) SetURLStartswith(uRLStartswith *string) {
	o.URLStartswith = uRLStartswith
}

// WithVendorType adds the vendorType to the ticket projects list params
func (o *TicketProjectsListParams) WithVendorType(vendorType *string) *TicketProjectsListParams {
	o.SetVendorType(vendorType)
	return o
}

// SetVendorType adds the vendorType to the ticket projects list params
func (o *TicketProjectsListParams) SetVendorType(vendorType *string) {
	o.VendorType = vendorType
}

// WriteToRequest writes these params to a swagger request
func (o *TicketProjectsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddedBy != nil {

		// query param added_by
		var qrAddedBy string
		if o.AddedBy != nil {
			qrAddedBy = *o.AddedBy
		}
		qAddedBy := qrAddedBy
		if qAddedBy != "" {
			if err := r.SetQueryParam("added_by", qAddedBy); err != nil {
				return err
			}
		}

	}

	if o.Filters != nil {

		// query param filters
		var qrFilters string
		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {
			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID float64
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatFloat64(qrID)
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn float64
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := swag.FormatFloat64(qrIDIn)
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
				return err
			}
		}

	}

	if o.IntegrationName != nil {

		// query param integration__name
		var qrIntegrationName string
		if o.IntegrationName != nil {
			qrIntegrationName = *o.IntegrationName
		}
		qIntegrationName := qrIntegrationName
		if qIntegrationName != "" {
			if err := r.SetQueryParam("integration__name", qIntegrationName); err != nil {
				return err
			}
		}

	}

	if o.IntegrationNameIcontains != nil {

		// query param integration__name__icontains
		var qrIntegrationNameIcontains string
		if o.IntegrationNameIcontains != nil {
			qrIntegrationNameIcontains = *o.IntegrationNameIcontains
		}
		qIntegrationNameIcontains := qrIntegrationNameIcontains
		if qIntegrationNameIcontains != "" {
			if err := r.SetQueryParam("integration__name__icontains", qIntegrationNameIcontains); err != nil {
				return err
			}
		}

	}

	if o.IntegrationNameStartswith != nil {

		// query param integration__name__startswith
		var qrIntegrationNameStartswith string
		if o.IntegrationNameStartswith != nil {
			qrIntegrationNameStartswith = *o.IntegrationNameStartswith
		}
		qIntegrationNameStartswith := qrIntegrationNameStartswith
		if qIntegrationNameStartswith != "" {
			if err := r.SetQueryParam("integration__name__startswith", qIntegrationNameStartswith); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.NameIcontains != nil {

		// query param name__icontains
		var qrNameIcontains string
		if o.NameIcontains != nil {
			qrNameIcontains = *o.NameIcontains
		}
		qNameIcontains := qrNameIcontains
		if qNameIcontains != "" {
			if err := r.SetQueryParam("name__icontains", qNameIcontains); err != nil {
				return err
			}
		}

	}

	if o.NameStartswith != nil {

		// query param name__startswith
		var qrNameStartswith string
		if o.NameStartswith != nil {
			qrNameStartswith = *o.NameStartswith
		}
		qNameStartswith := qrNameStartswith
		if qNameStartswith != "" {
			if err := r.SetQueryParam("name__startswith", qNameStartswith); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string
		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {
			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}

	}

	if o.Repo != nil {

		// query param repo
		var qrRepo string
		if o.Repo != nil {
			qrRepo = *o.Repo
		}
		qRepo := qrRepo
		if qRepo != "" {
			if err := r.SetQueryParam("repo", qRepo); err != nil {
				return err
			}
		}

	}

	if o.RepoID != nil {

		// query param repo_id
		var qrRepoID float64
		if o.RepoID != nil {
			qrRepoID = *o.RepoID
		}
		qRepoID := swag.FormatFloat64(qrRepoID)
		if qRepoID != "" {
			if err := r.SetQueryParam("repo_id", qRepoID); err != nil {
				return err
			}
		}

	}

	if o.RepoIDIn != nil {

		// query param repo_id__in
		var qrRepoIDIn string
		if o.RepoIDIn != nil {
			qrRepoIDIn = *o.RepoIDIn
		}
		qRepoIDIn := qrRepoIDIn
		if qRepoIDIn != "" {
			if err := r.SetQueryParam("repo_id__in", qRepoIDIn); err != nil {
				return err
			}
		}

	}

	if o.Search != nil {

		// query param search
		var qrSearch string
		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {
			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}

	}

	if o.UpdatedAt != nil {

		// query param updated_at
		var qrUpdatedAt string
		if o.UpdatedAt != nil {
			qrUpdatedAt = *o.UpdatedAt
		}
		qUpdatedAt := qrUpdatedAt
		if qUpdatedAt != "" {
			if err := r.SetQueryParam("updated_at", qUpdatedAt); err != nil {
				return err
			}
		}

	}

	if o.UpdatedAtGt != nil {

		// query param updated_at__gt
		var qrUpdatedAtGt string
		if o.UpdatedAtGt != nil {
			qrUpdatedAtGt = *o.UpdatedAtGt
		}
		qUpdatedAtGt := qrUpdatedAtGt
		if qUpdatedAtGt != "" {
			if err := r.SetQueryParam("updated_at__gt", qUpdatedAtGt); err != nil {
				return err
			}
		}

	}

	if o.UpdatedAtGte != nil {

		// query param updated_at__gte
		var qrUpdatedAtGte string
		if o.UpdatedAtGte != nil {
			qrUpdatedAtGte = *o.UpdatedAtGte
		}
		qUpdatedAtGte := qrUpdatedAtGte
		if qUpdatedAtGte != "" {
			if err := r.SetQueryParam("updated_at__gte", qUpdatedAtGte); err != nil {
				return err
			}
		}

	}

	if o.UpdatedAtLt != nil {

		// query param updated_at__lt
		var qrUpdatedAtLt string
		if o.UpdatedAtLt != nil {
			qrUpdatedAtLt = *o.UpdatedAtLt
		}
		qUpdatedAtLt := qrUpdatedAtLt
		if qUpdatedAtLt != "" {
			if err := r.SetQueryParam("updated_at__lt", qUpdatedAtLt); err != nil {
				return err
			}
		}

	}

	if o.UpdatedAtLte != nil {

		// query param updated_at__lte
		var qrUpdatedAtLte string
		if o.UpdatedAtLte != nil {
			qrUpdatedAtLte = *o.UpdatedAtLte
		}
		qUpdatedAtLte := qrUpdatedAtLte
		if qUpdatedAtLte != "" {
			if err := r.SetQueryParam("updated_at__lte", qUpdatedAtLte); err != nil {
				return err
			}
		}

	}

	if o.URL != nil {

		// query param url
		var qrURL string
		if o.URL != nil {
			qrURL = *o.URL
		}
		qURL := qrURL
		if qURL != "" {
			if err := r.SetQueryParam("url", qURL); err != nil {
				return err
			}
		}

	}

	if o.URLIcontains != nil {

		// query param url__icontains
		var qrURLIcontains string
		if o.URLIcontains != nil {
			qrURLIcontains = *o.URLIcontains
		}
		qURLIcontains := qrURLIcontains
		if qURLIcontains != "" {
			if err := r.SetQueryParam("url__icontains", qURLIcontains); err != nil {
				return err
			}
		}

	}

	if o.URLStartswith != nil {

		// query param url__startswith
		var qrURLStartswith string
		if o.URLStartswith != nil {
			qrURLStartswith = *o.URLStartswith
		}
		qURLStartswith := qrURLStartswith
		if qURLStartswith != "" {
			if err := r.SetQueryParam("url__startswith", qURLStartswith); err != nil {
				return err
			}
		}

	}

	if o.VendorType != nil {

		// query param vendor_type
		var qrVendorType string
		if o.VendorType != nil {
			qrVendorType = *o.VendorType
		}
		qVendorType := qrVendorType
		if qVendorType != "" {
			if err := r.SetQueryParam("vendor_type", qVendorType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
