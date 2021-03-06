// Code generated by go-swagger; DO NOT EDIT.

package user_alias

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

// NewUserAliasListParams creates a new UserAliasListParams object
// with the default values initialized.
func NewUserAliasListParams() *UserAliasListParams {
	var ()
	return &UserAliasListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserAliasListParamsWithTimeout creates a new UserAliasListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserAliasListParamsWithTimeout(timeout time.Duration) *UserAliasListParams {
	var ()
	return &UserAliasListParams{

		timeout: timeout,
	}
}

// NewUserAliasListParamsWithContext creates a new UserAliasListParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserAliasListParamsWithContext(ctx context.Context) *UserAliasListParams {
	var ()
	return &UserAliasListParams{

		Context: ctx,
	}
}

// NewUserAliasListParamsWithHTTPClient creates a new UserAliasListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserAliasListParamsWithHTTPClient(client *http.Client) *UserAliasListParams {
	var ()
	return &UserAliasListParams{
		HTTPClient: client,
	}
}

/*UserAliasListParams contains all the parameters to send to the API endpoint
for the user alias list operation typically these are written to a http.Request
*/
type UserAliasListParams struct {

	/*ApexUser
	  Filter-traversable object

	*/
	ApexUser *string
	/*ApexUserIn
	  Match on apex_user_id in a comma separated list

	*/
	ApexUserIn *string
	/*CreatedAt
	  Records == date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	CreatedAt *string
	/*CreatedAtGt
	  Records > date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	CreatedAtGt *string
	/*CreatedAtGte
	  Records >= date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	CreatedAtGte *string
	/*CreatedAtLt
	  Records < date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	CreatedAtLt *string
	/*CreatedAtLte
	  Records <= date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	CreatedAtLte *string
	/*Email*/
	Email *string
	/*EmailIcontains*/
	EmailIcontains *string
	/*EmailStartswith*/
	EmailStartswith *string
	/*Filters
	  Complex filter that handles filtering with groups of conditions connected with AND and OR conjunctions. The query can be arbitrarily grouped and nested.

	*/
	Filters *string
	/*FirstActivityAt
	  Records == date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	FirstActivityAt *string
	/*FirstActivityAtGt
	  Records > date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	FirstActivityAtGt *string
	/*FirstActivityAtGte
	  Records >= date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	FirstActivityAtGte *string
	/*FirstActivityAtLt
	  Records < date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	FirstActivityAtLt *string
	/*FirstActivityAtLte
	  Records <= date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	FirstActivityAtLte *string
	/*ID*/
	ID *float64
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *float64
	/*Integration
	  Filter-traversable object

	*/
	Integration *string
	/*LastActivityAt
	  Records == date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	LastActivityAt *string
	/*LastActivityAtGt
	  Records > date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	LastActivityAtGt *string
	/*LastActivityAtGte
	  Records >= date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	LastActivityAtGte *string
	/*LastActivityAtLt
	  Records < date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	LastActivityAtLt *string
	/*LastActivityAtLte
	  Records <= date or datetime (yyyy-mm-dd hh:mm:ss)

	*/
	LastActivityAtLte *string
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
	  Valid ordering fields are: `id`, `name`, `username`, `email`, `created_at`, `org_id`, `apex_user_id`, `apex_user__name`

	*/
	Ordering *string
	/*Search
	  A search term.

	*/
	Search *string
	/*Username*/
	Username *string
	/*UsernameIcontains*/
	UsernameIcontains *string
	/*UsernameStartswith*/
	UsernameStartswith *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user alias list params
func (o *UserAliasListParams) WithTimeout(timeout time.Duration) *UserAliasListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user alias list params
func (o *UserAliasListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user alias list params
func (o *UserAliasListParams) WithContext(ctx context.Context) *UserAliasListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user alias list params
func (o *UserAliasListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user alias list params
func (o *UserAliasListParams) WithHTTPClient(client *http.Client) *UserAliasListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user alias list params
func (o *UserAliasListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApexUser adds the apexUser to the user alias list params
func (o *UserAliasListParams) WithApexUser(apexUser *string) *UserAliasListParams {
	o.SetApexUser(apexUser)
	return o
}

// SetApexUser adds the apexUser to the user alias list params
func (o *UserAliasListParams) SetApexUser(apexUser *string) {
	o.ApexUser = apexUser
}

// WithApexUserIn adds the apexUserIn to the user alias list params
func (o *UserAliasListParams) WithApexUserIn(apexUserIn *string) *UserAliasListParams {
	o.SetApexUserIn(apexUserIn)
	return o
}

// SetApexUserIn adds the apexUserIn to the user alias list params
func (o *UserAliasListParams) SetApexUserIn(apexUserIn *string) {
	o.ApexUserIn = apexUserIn
}

// WithCreatedAt adds the createdAt to the user alias list params
func (o *UserAliasListParams) WithCreatedAt(createdAt *string) *UserAliasListParams {
	o.SetCreatedAt(createdAt)
	return o
}

// SetCreatedAt adds the createdAt to the user alias list params
func (o *UserAliasListParams) SetCreatedAt(createdAt *string) {
	o.CreatedAt = createdAt
}

// WithCreatedAtGt adds the createdAtGt to the user alias list params
func (o *UserAliasListParams) WithCreatedAtGt(createdAtGt *string) *UserAliasListParams {
	o.SetCreatedAtGt(createdAtGt)
	return o
}

// SetCreatedAtGt adds the createdAtGt to the user alias list params
func (o *UserAliasListParams) SetCreatedAtGt(createdAtGt *string) {
	o.CreatedAtGt = createdAtGt
}

// WithCreatedAtGte adds the createdAtGte to the user alias list params
func (o *UserAliasListParams) WithCreatedAtGte(createdAtGte *string) *UserAliasListParams {
	o.SetCreatedAtGte(createdAtGte)
	return o
}

// SetCreatedAtGte adds the createdAtGte to the user alias list params
func (o *UserAliasListParams) SetCreatedAtGte(createdAtGte *string) {
	o.CreatedAtGte = createdAtGte
}

// WithCreatedAtLt adds the createdAtLt to the user alias list params
func (o *UserAliasListParams) WithCreatedAtLt(createdAtLt *string) *UserAliasListParams {
	o.SetCreatedAtLt(createdAtLt)
	return o
}

// SetCreatedAtLt adds the createdAtLt to the user alias list params
func (o *UserAliasListParams) SetCreatedAtLt(createdAtLt *string) {
	o.CreatedAtLt = createdAtLt
}

// WithCreatedAtLte adds the createdAtLte to the user alias list params
func (o *UserAliasListParams) WithCreatedAtLte(createdAtLte *string) *UserAliasListParams {
	o.SetCreatedAtLte(createdAtLte)
	return o
}

// SetCreatedAtLte adds the createdAtLte to the user alias list params
func (o *UserAliasListParams) SetCreatedAtLte(createdAtLte *string) {
	o.CreatedAtLte = createdAtLte
}

// WithEmail adds the email to the user alias list params
func (o *UserAliasListParams) WithEmail(email *string) *UserAliasListParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the user alias list params
func (o *UserAliasListParams) SetEmail(email *string) {
	o.Email = email
}

// WithEmailIcontains adds the emailIcontains to the user alias list params
func (o *UserAliasListParams) WithEmailIcontains(emailIcontains *string) *UserAliasListParams {
	o.SetEmailIcontains(emailIcontains)
	return o
}

// SetEmailIcontains adds the emailIcontains to the user alias list params
func (o *UserAliasListParams) SetEmailIcontains(emailIcontains *string) {
	o.EmailIcontains = emailIcontains
}

// WithEmailStartswith adds the emailStartswith to the user alias list params
func (o *UserAliasListParams) WithEmailStartswith(emailStartswith *string) *UserAliasListParams {
	o.SetEmailStartswith(emailStartswith)
	return o
}

// SetEmailStartswith adds the emailStartswith to the user alias list params
func (o *UserAliasListParams) SetEmailStartswith(emailStartswith *string) {
	o.EmailStartswith = emailStartswith
}

// WithFilters adds the filters to the user alias list params
func (o *UserAliasListParams) WithFilters(filters *string) *UserAliasListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the user alias list params
func (o *UserAliasListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithFirstActivityAt adds the firstActivityAt to the user alias list params
func (o *UserAliasListParams) WithFirstActivityAt(firstActivityAt *string) *UserAliasListParams {
	o.SetFirstActivityAt(firstActivityAt)
	return o
}

// SetFirstActivityAt adds the firstActivityAt to the user alias list params
func (o *UserAliasListParams) SetFirstActivityAt(firstActivityAt *string) {
	o.FirstActivityAt = firstActivityAt
}

// WithFirstActivityAtGt adds the firstActivityAtGt to the user alias list params
func (o *UserAliasListParams) WithFirstActivityAtGt(firstActivityAtGt *string) *UserAliasListParams {
	o.SetFirstActivityAtGt(firstActivityAtGt)
	return o
}

// SetFirstActivityAtGt adds the firstActivityAtGt to the user alias list params
func (o *UserAliasListParams) SetFirstActivityAtGt(firstActivityAtGt *string) {
	o.FirstActivityAtGt = firstActivityAtGt
}

// WithFirstActivityAtGte adds the firstActivityAtGte to the user alias list params
func (o *UserAliasListParams) WithFirstActivityAtGte(firstActivityAtGte *string) *UserAliasListParams {
	o.SetFirstActivityAtGte(firstActivityAtGte)
	return o
}

// SetFirstActivityAtGte adds the firstActivityAtGte to the user alias list params
func (o *UserAliasListParams) SetFirstActivityAtGte(firstActivityAtGte *string) {
	o.FirstActivityAtGte = firstActivityAtGte
}

// WithFirstActivityAtLt adds the firstActivityAtLt to the user alias list params
func (o *UserAliasListParams) WithFirstActivityAtLt(firstActivityAtLt *string) *UserAliasListParams {
	o.SetFirstActivityAtLt(firstActivityAtLt)
	return o
}

// SetFirstActivityAtLt adds the firstActivityAtLt to the user alias list params
func (o *UserAliasListParams) SetFirstActivityAtLt(firstActivityAtLt *string) {
	o.FirstActivityAtLt = firstActivityAtLt
}

// WithFirstActivityAtLte adds the firstActivityAtLte to the user alias list params
func (o *UserAliasListParams) WithFirstActivityAtLte(firstActivityAtLte *string) *UserAliasListParams {
	o.SetFirstActivityAtLte(firstActivityAtLte)
	return o
}

// SetFirstActivityAtLte adds the firstActivityAtLte to the user alias list params
func (o *UserAliasListParams) SetFirstActivityAtLte(firstActivityAtLte *string) {
	o.FirstActivityAtLte = firstActivityAtLte
}

// WithID adds the id to the user alias list params
func (o *UserAliasListParams) WithID(id *float64) *UserAliasListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user alias list params
func (o *UserAliasListParams) SetID(id *float64) {
	o.ID = id
}

// WithIDIn adds the iDIn to the user alias list params
func (o *UserAliasListParams) WithIDIn(iDIn *float64) *UserAliasListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the user alias list params
func (o *UserAliasListParams) SetIDIn(iDIn *float64) {
	o.IDIn = iDIn
}

// WithIntegration adds the integration to the user alias list params
func (o *UserAliasListParams) WithIntegration(integration *string) *UserAliasListParams {
	o.SetIntegration(integration)
	return o
}

// SetIntegration adds the integration to the user alias list params
func (o *UserAliasListParams) SetIntegration(integration *string) {
	o.Integration = integration
}

// WithLastActivityAt adds the lastActivityAt to the user alias list params
func (o *UserAliasListParams) WithLastActivityAt(lastActivityAt *string) *UserAliasListParams {
	o.SetLastActivityAt(lastActivityAt)
	return o
}

// SetLastActivityAt adds the lastActivityAt to the user alias list params
func (o *UserAliasListParams) SetLastActivityAt(lastActivityAt *string) {
	o.LastActivityAt = lastActivityAt
}

// WithLastActivityAtGt adds the lastActivityAtGt to the user alias list params
func (o *UserAliasListParams) WithLastActivityAtGt(lastActivityAtGt *string) *UserAliasListParams {
	o.SetLastActivityAtGt(lastActivityAtGt)
	return o
}

// SetLastActivityAtGt adds the lastActivityAtGt to the user alias list params
func (o *UserAliasListParams) SetLastActivityAtGt(lastActivityAtGt *string) {
	o.LastActivityAtGt = lastActivityAtGt
}

// WithLastActivityAtGte adds the lastActivityAtGte to the user alias list params
func (o *UserAliasListParams) WithLastActivityAtGte(lastActivityAtGte *string) *UserAliasListParams {
	o.SetLastActivityAtGte(lastActivityAtGte)
	return o
}

// SetLastActivityAtGte adds the lastActivityAtGte to the user alias list params
func (o *UserAliasListParams) SetLastActivityAtGte(lastActivityAtGte *string) {
	o.LastActivityAtGte = lastActivityAtGte
}

// WithLastActivityAtLt adds the lastActivityAtLt to the user alias list params
func (o *UserAliasListParams) WithLastActivityAtLt(lastActivityAtLt *string) *UserAliasListParams {
	o.SetLastActivityAtLt(lastActivityAtLt)
	return o
}

// SetLastActivityAtLt adds the lastActivityAtLt to the user alias list params
func (o *UserAliasListParams) SetLastActivityAtLt(lastActivityAtLt *string) {
	o.LastActivityAtLt = lastActivityAtLt
}

// WithLastActivityAtLte adds the lastActivityAtLte to the user alias list params
func (o *UserAliasListParams) WithLastActivityAtLte(lastActivityAtLte *string) *UserAliasListParams {
	o.SetLastActivityAtLte(lastActivityAtLte)
	return o
}

// SetLastActivityAtLte adds the lastActivityAtLte to the user alias list params
func (o *UserAliasListParams) SetLastActivityAtLte(lastActivityAtLte *string) {
	o.LastActivityAtLte = lastActivityAtLte
}

// WithLimit adds the limit to the user alias list params
func (o *UserAliasListParams) WithLimit(limit *int64) *UserAliasListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the user alias list params
func (o *UserAliasListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the user alias list params
func (o *UserAliasListParams) WithName(name *string) *UserAliasListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the user alias list params
func (o *UserAliasListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIcontains adds the nameIcontains to the user alias list params
func (o *UserAliasListParams) WithNameIcontains(nameIcontains *string) *UserAliasListParams {
	o.SetNameIcontains(nameIcontains)
	return o
}

// SetNameIcontains adds the nameIcontains to the user alias list params
func (o *UserAliasListParams) SetNameIcontains(nameIcontains *string) {
	o.NameIcontains = nameIcontains
}

// WithNameStartswith adds the nameStartswith to the user alias list params
func (o *UserAliasListParams) WithNameStartswith(nameStartswith *string) *UserAliasListParams {
	o.SetNameStartswith(nameStartswith)
	return o
}

// SetNameStartswith adds the nameStartswith to the user alias list params
func (o *UserAliasListParams) SetNameStartswith(nameStartswith *string) {
	o.NameStartswith = nameStartswith
}

// WithOffset adds the offset to the user alias list params
func (o *UserAliasListParams) WithOffset(offset *int64) *UserAliasListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the user alias list params
func (o *UserAliasListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the user alias list params
func (o *UserAliasListParams) WithOrdering(ordering *string) *UserAliasListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the user alias list params
func (o *UserAliasListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithSearch adds the search to the user alias list params
func (o *UserAliasListParams) WithSearch(search *string) *UserAliasListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the user alias list params
func (o *UserAliasListParams) SetSearch(search *string) {
	o.Search = search
}

// WithUsername adds the username to the user alias list params
func (o *UserAliasListParams) WithUsername(username *string) *UserAliasListParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the user alias list params
func (o *UserAliasListParams) SetUsername(username *string) {
	o.Username = username
}

// WithUsernameIcontains adds the usernameIcontains to the user alias list params
func (o *UserAliasListParams) WithUsernameIcontains(usernameIcontains *string) *UserAliasListParams {
	o.SetUsernameIcontains(usernameIcontains)
	return o
}

// SetUsernameIcontains adds the usernameIcontains to the user alias list params
func (o *UserAliasListParams) SetUsernameIcontains(usernameIcontains *string) {
	o.UsernameIcontains = usernameIcontains
}

// WithUsernameStartswith adds the usernameStartswith to the user alias list params
func (o *UserAliasListParams) WithUsernameStartswith(usernameStartswith *string) *UserAliasListParams {
	o.SetUsernameStartswith(usernameStartswith)
	return o
}

// SetUsernameStartswith adds the usernameStartswith to the user alias list params
func (o *UserAliasListParams) SetUsernameStartswith(usernameStartswith *string) {
	o.UsernameStartswith = usernameStartswith
}

// WriteToRequest writes these params to a swagger request
func (o *UserAliasListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApexUser != nil {

		// query param apex_user
		var qrApexUser string
		if o.ApexUser != nil {
			qrApexUser = *o.ApexUser
		}
		qApexUser := qrApexUser
		if qApexUser != "" {
			if err := r.SetQueryParam("apex_user", qApexUser); err != nil {
				return err
			}
		}

	}

	if o.ApexUserIn != nil {

		// query param apex_user__in
		var qrApexUserIn string
		if o.ApexUserIn != nil {
			qrApexUserIn = *o.ApexUserIn
		}
		qApexUserIn := qrApexUserIn
		if qApexUserIn != "" {
			if err := r.SetQueryParam("apex_user__in", qApexUserIn); err != nil {
				return err
			}
		}

	}

	if o.CreatedAt != nil {

		// query param created_at
		var qrCreatedAt string
		if o.CreatedAt != nil {
			qrCreatedAt = *o.CreatedAt
		}
		qCreatedAt := qrCreatedAt
		if qCreatedAt != "" {
			if err := r.SetQueryParam("created_at", qCreatedAt); err != nil {
				return err
			}
		}

	}

	if o.CreatedAtGt != nil {

		// query param created_at__gt
		var qrCreatedAtGt string
		if o.CreatedAtGt != nil {
			qrCreatedAtGt = *o.CreatedAtGt
		}
		qCreatedAtGt := qrCreatedAtGt
		if qCreatedAtGt != "" {
			if err := r.SetQueryParam("created_at__gt", qCreatedAtGt); err != nil {
				return err
			}
		}

	}

	if o.CreatedAtGte != nil {

		// query param created_at__gte
		var qrCreatedAtGte string
		if o.CreatedAtGte != nil {
			qrCreatedAtGte = *o.CreatedAtGte
		}
		qCreatedAtGte := qrCreatedAtGte
		if qCreatedAtGte != "" {
			if err := r.SetQueryParam("created_at__gte", qCreatedAtGte); err != nil {
				return err
			}
		}

	}

	if o.CreatedAtLt != nil {

		// query param created_at__lt
		var qrCreatedAtLt string
		if o.CreatedAtLt != nil {
			qrCreatedAtLt = *o.CreatedAtLt
		}
		qCreatedAtLt := qrCreatedAtLt
		if qCreatedAtLt != "" {
			if err := r.SetQueryParam("created_at__lt", qCreatedAtLt); err != nil {
				return err
			}
		}

	}

	if o.CreatedAtLte != nil {

		// query param created_at__lte
		var qrCreatedAtLte string
		if o.CreatedAtLte != nil {
			qrCreatedAtLte = *o.CreatedAtLte
		}
		qCreatedAtLte := qrCreatedAtLte
		if qCreatedAtLte != "" {
			if err := r.SetQueryParam("created_at__lte", qCreatedAtLte); err != nil {
				return err
			}
		}

	}

	if o.Email != nil {

		// query param email
		var qrEmail string
		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {
			if err := r.SetQueryParam("email", qEmail); err != nil {
				return err
			}
		}

	}

	if o.EmailIcontains != nil {

		// query param email__icontains
		var qrEmailIcontains string
		if o.EmailIcontains != nil {
			qrEmailIcontains = *o.EmailIcontains
		}
		qEmailIcontains := qrEmailIcontains
		if qEmailIcontains != "" {
			if err := r.SetQueryParam("email__icontains", qEmailIcontains); err != nil {
				return err
			}
		}

	}

	if o.EmailStartswith != nil {

		// query param email__startswith
		var qrEmailStartswith string
		if o.EmailStartswith != nil {
			qrEmailStartswith = *o.EmailStartswith
		}
		qEmailStartswith := qrEmailStartswith
		if qEmailStartswith != "" {
			if err := r.SetQueryParam("email__startswith", qEmailStartswith); err != nil {
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

	if o.FirstActivityAt != nil {

		// query param first_activity_at
		var qrFirstActivityAt string
		if o.FirstActivityAt != nil {
			qrFirstActivityAt = *o.FirstActivityAt
		}
		qFirstActivityAt := qrFirstActivityAt
		if qFirstActivityAt != "" {
			if err := r.SetQueryParam("first_activity_at", qFirstActivityAt); err != nil {
				return err
			}
		}

	}

	if o.FirstActivityAtGt != nil {

		// query param first_activity_at__gt
		var qrFirstActivityAtGt string
		if o.FirstActivityAtGt != nil {
			qrFirstActivityAtGt = *o.FirstActivityAtGt
		}
		qFirstActivityAtGt := qrFirstActivityAtGt
		if qFirstActivityAtGt != "" {
			if err := r.SetQueryParam("first_activity_at__gt", qFirstActivityAtGt); err != nil {
				return err
			}
		}

	}

	if o.FirstActivityAtGte != nil {

		// query param first_activity_at__gte
		var qrFirstActivityAtGte string
		if o.FirstActivityAtGte != nil {
			qrFirstActivityAtGte = *o.FirstActivityAtGte
		}
		qFirstActivityAtGte := qrFirstActivityAtGte
		if qFirstActivityAtGte != "" {
			if err := r.SetQueryParam("first_activity_at__gte", qFirstActivityAtGte); err != nil {
				return err
			}
		}

	}

	if o.FirstActivityAtLt != nil {

		// query param first_activity_at__lt
		var qrFirstActivityAtLt string
		if o.FirstActivityAtLt != nil {
			qrFirstActivityAtLt = *o.FirstActivityAtLt
		}
		qFirstActivityAtLt := qrFirstActivityAtLt
		if qFirstActivityAtLt != "" {
			if err := r.SetQueryParam("first_activity_at__lt", qFirstActivityAtLt); err != nil {
				return err
			}
		}

	}

	if o.FirstActivityAtLte != nil {

		// query param first_activity_at__lte
		var qrFirstActivityAtLte string
		if o.FirstActivityAtLte != nil {
			qrFirstActivityAtLte = *o.FirstActivityAtLte
		}
		qFirstActivityAtLte := qrFirstActivityAtLte
		if qFirstActivityAtLte != "" {
			if err := r.SetQueryParam("first_activity_at__lte", qFirstActivityAtLte); err != nil {
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

	if o.Integration != nil {

		// query param integration
		var qrIntegration string
		if o.Integration != nil {
			qrIntegration = *o.Integration
		}
		qIntegration := qrIntegration
		if qIntegration != "" {
			if err := r.SetQueryParam("integration", qIntegration); err != nil {
				return err
			}
		}

	}

	if o.LastActivityAt != nil {

		// query param last_activity_at
		var qrLastActivityAt string
		if o.LastActivityAt != nil {
			qrLastActivityAt = *o.LastActivityAt
		}
		qLastActivityAt := qrLastActivityAt
		if qLastActivityAt != "" {
			if err := r.SetQueryParam("last_activity_at", qLastActivityAt); err != nil {
				return err
			}
		}

	}

	if o.LastActivityAtGt != nil {

		// query param last_activity_at__gt
		var qrLastActivityAtGt string
		if o.LastActivityAtGt != nil {
			qrLastActivityAtGt = *o.LastActivityAtGt
		}
		qLastActivityAtGt := qrLastActivityAtGt
		if qLastActivityAtGt != "" {
			if err := r.SetQueryParam("last_activity_at__gt", qLastActivityAtGt); err != nil {
				return err
			}
		}

	}

	if o.LastActivityAtGte != nil {

		// query param last_activity_at__gte
		var qrLastActivityAtGte string
		if o.LastActivityAtGte != nil {
			qrLastActivityAtGte = *o.LastActivityAtGte
		}
		qLastActivityAtGte := qrLastActivityAtGte
		if qLastActivityAtGte != "" {
			if err := r.SetQueryParam("last_activity_at__gte", qLastActivityAtGte); err != nil {
				return err
			}
		}

	}

	if o.LastActivityAtLt != nil {

		// query param last_activity_at__lt
		var qrLastActivityAtLt string
		if o.LastActivityAtLt != nil {
			qrLastActivityAtLt = *o.LastActivityAtLt
		}
		qLastActivityAtLt := qrLastActivityAtLt
		if qLastActivityAtLt != "" {
			if err := r.SetQueryParam("last_activity_at__lt", qLastActivityAtLt); err != nil {
				return err
			}
		}

	}

	if o.LastActivityAtLte != nil {

		// query param last_activity_at__lte
		var qrLastActivityAtLte string
		if o.LastActivityAtLte != nil {
			qrLastActivityAtLte = *o.LastActivityAtLte
		}
		qLastActivityAtLte := qrLastActivityAtLte
		if qLastActivityAtLte != "" {
			if err := r.SetQueryParam("last_activity_at__lte", qLastActivityAtLte); err != nil {
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

	if o.Username != nil {

		// query param username
		var qrUsername string
		if o.Username != nil {
			qrUsername = *o.Username
		}
		qUsername := qrUsername
		if qUsername != "" {
			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
		}

	}

	if o.UsernameIcontains != nil {

		// query param username__icontains
		var qrUsernameIcontains string
		if o.UsernameIcontains != nil {
			qrUsernameIcontains = *o.UsernameIcontains
		}
		qUsernameIcontains := qrUsernameIcontains
		if qUsernameIcontains != "" {
			if err := r.SetQueryParam("username__icontains", qUsernameIcontains); err != nil {
				return err
			}
		}

	}

	if o.UsernameStartswith != nil {

		// query param username__startswith
		var qrUsernameStartswith string
		if o.UsernameStartswith != nil {
			qrUsernameStartswith = *o.UsernameStartswith
		}
		qUsernameStartswith := qrUsernameStartswith
		if qUsernameStartswith != "" {
			if err := r.SetQueryParam("username__startswith", qUsernameStartswith); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
