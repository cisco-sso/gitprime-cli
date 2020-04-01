// Code generated by go-swagger; DO NOT EDIT.

package team_membership

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

// NewTeamMembershipListParams creates a new TeamMembershipListParams object
// with the default values initialized.
func NewTeamMembershipListParams() *TeamMembershipListParams {
	var ()
	return &TeamMembershipListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTeamMembershipListParamsWithTimeout creates a new TeamMembershipListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTeamMembershipListParamsWithTimeout(timeout time.Duration) *TeamMembershipListParams {
	var ()
	return &TeamMembershipListParams{

		timeout: timeout,
	}
}

// NewTeamMembershipListParamsWithContext creates a new TeamMembershipListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTeamMembershipListParamsWithContext(ctx context.Context) *TeamMembershipListParams {
	var ()
	return &TeamMembershipListParams{

		Context: ctx,
	}
}

// NewTeamMembershipListParamsWithHTTPClient creates a new TeamMembershipListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTeamMembershipListParamsWithHTTPClient(client *http.Client) *TeamMembershipListParams {
	var ()
	return &TeamMembershipListParams{
		HTTPClient: client,
	}
}

/*TeamMembershipListParams contains all the parameters to send to the API endpoint
for the team membership list operation typically these are written to a http.Request
*/
type TeamMembershipListParams struct {

	/*ApexUser
	  Filter-traversable object

	*/
	ApexUser *string
	/*ApexUserNameIcontains*/
	ApexUserNameIcontains *string
	/*ApexUserID*/
	ApexUserID *string
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
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*MembershipType*/
	MembershipType *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Ordering
	  Valid ordering fields are: `apex_user_id`, `apex_user__hidden_from_reports`, `apex_user__name`, `depth`, `id`, `membership_type`, `team_id`

	*/
	Ordering *string
	/*Search
	  A search term.

	*/
	Search *string
	/*Team
	  Filter-traversable object

	*/
	Team *string
	/*TeamAncestorID
	  Match team members who are a member of the team or its descendant teams

	*/
	TeamAncestorID *float64
	/*TeamID*/
	TeamID *string
	/*TeamIDIn
	  Multiple values may be separated by commas.

	*/
	TeamIDIn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the team membership list params
func (o *TeamMembershipListParams) WithTimeout(timeout time.Duration) *TeamMembershipListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the team membership list params
func (o *TeamMembershipListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the team membership list params
func (o *TeamMembershipListParams) WithContext(ctx context.Context) *TeamMembershipListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the team membership list params
func (o *TeamMembershipListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the team membership list params
func (o *TeamMembershipListParams) WithHTTPClient(client *http.Client) *TeamMembershipListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the team membership list params
func (o *TeamMembershipListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApexUser adds the apexUser to the team membership list params
func (o *TeamMembershipListParams) WithApexUser(apexUser *string) *TeamMembershipListParams {
	o.SetApexUser(apexUser)
	return o
}

// SetApexUser adds the apexUser to the team membership list params
func (o *TeamMembershipListParams) SetApexUser(apexUser *string) {
	o.ApexUser = apexUser
}

// WithApexUserNameIcontains adds the apexUserNameIcontains to the team membership list params
func (o *TeamMembershipListParams) WithApexUserNameIcontains(apexUserNameIcontains *string) *TeamMembershipListParams {
	o.SetApexUserNameIcontains(apexUserNameIcontains)
	return o
}

// SetApexUserNameIcontains adds the apexUserNameIcontains to the team membership list params
func (o *TeamMembershipListParams) SetApexUserNameIcontains(apexUserNameIcontains *string) {
	o.ApexUserNameIcontains = apexUserNameIcontains
}

// WithApexUserID adds the apexUserID to the team membership list params
func (o *TeamMembershipListParams) WithApexUserID(apexUserID *string) *TeamMembershipListParams {
	o.SetApexUserID(apexUserID)
	return o
}

// SetApexUserID adds the apexUserId to the team membership list params
func (o *TeamMembershipListParams) SetApexUserID(apexUserID *string) {
	o.ApexUserID = apexUserID
}

// WithFilters adds the filters to the team membership list params
func (o *TeamMembershipListParams) WithFilters(filters *string) *TeamMembershipListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the team membership list params
func (o *TeamMembershipListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithID adds the id to the team membership list params
func (o *TeamMembershipListParams) WithID(id *float64) *TeamMembershipListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the team membership list params
func (o *TeamMembershipListParams) SetID(id *float64) {
	o.ID = id
}

// WithIDIn adds the iDIn to the team membership list params
func (o *TeamMembershipListParams) WithIDIn(iDIn *float64) *TeamMembershipListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the team membership list params
func (o *TeamMembershipListParams) SetIDIn(iDIn *float64) {
	o.IDIn = iDIn
}

// WithLimit adds the limit to the team membership list params
func (o *TeamMembershipListParams) WithLimit(limit *int64) *TeamMembershipListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the team membership list params
func (o *TeamMembershipListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMembershipType adds the membershipType to the team membership list params
func (o *TeamMembershipListParams) WithMembershipType(membershipType *string) *TeamMembershipListParams {
	o.SetMembershipType(membershipType)
	return o
}

// SetMembershipType adds the membershipType to the team membership list params
func (o *TeamMembershipListParams) SetMembershipType(membershipType *string) {
	o.MembershipType = membershipType
}

// WithOffset adds the offset to the team membership list params
func (o *TeamMembershipListParams) WithOffset(offset *int64) *TeamMembershipListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the team membership list params
func (o *TeamMembershipListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the team membership list params
func (o *TeamMembershipListParams) WithOrdering(ordering *string) *TeamMembershipListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the team membership list params
func (o *TeamMembershipListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithSearch adds the search to the team membership list params
func (o *TeamMembershipListParams) WithSearch(search *string) *TeamMembershipListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the team membership list params
func (o *TeamMembershipListParams) SetSearch(search *string) {
	o.Search = search
}

// WithTeam adds the team to the team membership list params
func (o *TeamMembershipListParams) WithTeam(team *string) *TeamMembershipListParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the team membership list params
func (o *TeamMembershipListParams) SetTeam(team *string) {
	o.Team = team
}

// WithTeamAncestorID adds the teamAncestorID to the team membership list params
func (o *TeamMembershipListParams) WithTeamAncestorID(teamAncestorID *float64) *TeamMembershipListParams {
	o.SetTeamAncestorID(teamAncestorID)
	return o
}

// SetTeamAncestorID adds the teamAncestorId to the team membership list params
func (o *TeamMembershipListParams) SetTeamAncestorID(teamAncestorID *float64) {
	o.TeamAncestorID = teamAncestorID
}

// WithTeamID adds the teamID to the team membership list params
func (o *TeamMembershipListParams) WithTeamID(teamID *string) *TeamMembershipListParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the team membership list params
func (o *TeamMembershipListParams) SetTeamID(teamID *string) {
	o.TeamID = teamID
}

// WithTeamIDIn adds the teamIDIn to the team membership list params
func (o *TeamMembershipListParams) WithTeamIDIn(teamIDIn *string) *TeamMembershipListParams {
	o.SetTeamIDIn(teamIDIn)
	return o
}

// SetTeamIDIn adds the teamIdIn to the team membership list params
func (o *TeamMembershipListParams) SetTeamIDIn(teamIDIn *string) {
	o.TeamIDIn = teamIDIn
}

// WriteToRequest writes these params to a swagger request
func (o *TeamMembershipListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ApexUserNameIcontains != nil {

		// query param apex_user__name__icontains
		var qrApexUserNameIcontains string
		if o.ApexUserNameIcontains != nil {
			qrApexUserNameIcontains = *o.ApexUserNameIcontains
		}
		qApexUserNameIcontains := qrApexUserNameIcontains
		if qApexUserNameIcontains != "" {
			if err := r.SetQueryParam("apex_user__name__icontains", qApexUserNameIcontains); err != nil {
				return err
			}
		}

	}

	if o.ApexUserID != nil {

		// query param apex_user_id
		var qrApexUserID string
		if o.ApexUserID != nil {
			qrApexUserID = *o.ApexUserID
		}
		qApexUserID := qrApexUserID
		if qApexUserID != "" {
			if err := r.SetQueryParam("apex_user_id", qApexUserID); err != nil {
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

	if o.MembershipType != nil {

		// query param membership_type
		var qrMembershipType string
		if o.MembershipType != nil {
			qrMembershipType = *o.MembershipType
		}
		qMembershipType := qrMembershipType
		if qMembershipType != "" {
			if err := r.SetQueryParam("membership_type", qMembershipType); err != nil {
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

	if o.Team != nil {

		// query param team
		var qrTeam string
		if o.Team != nil {
			qrTeam = *o.Team
		}
		qTeam := qrTeam
		if qTeam != "" {
			if err := r.SetQueryParam("team", qTeam); err != nil {
				return err
			}
		}

	}

	if o.TeamAncestorID != nil {

		// query param team_ancestor_id
		var qrTeamAncestorID float64
		if o.TeamAncestorID != nil {
			qrTeamAncestorID = *o.TeamAncestorID
		}
		qTeamAncestorID := swag.FormatFloat64(qrTeamAncestorID)
		if qTeamAncestorID != "" {
			if err := r.SetQueryParam("team_ancestor_id", qTeamAncestorID); err != nil {
				return err
			}
		}

	}

	if o.TeamID != nil {

		// query param team_id
		var qrTeamID string
		if o.TeamID != nil {
			qrTeamID = *o.TeamID
		}
		qTeamID := qrTeamID
		if qTeamID != "" {
			if err := r.SetQueryParam("team_id", qTeamID); err != nil {
				return err
			}
		}

	}

	if o.TeamIDIn != nil {

		// query param team_id__in
		var qrTeamIDIn string
		if o.TeamIDIn != nil {
			qrTeamIDIn = *o.TeamIDIn
		}
		qTeamIDIn := qrTeamIDIn
		if qTeamIDIn != "" {
			if err := r.SetQueryParam("team_id__in", qTeamIDIn); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
