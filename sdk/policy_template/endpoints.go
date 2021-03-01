// Code generated by goa v3.1.3, DO NOT EDIT.
//
// PolicyTemplate endpoints
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design -o .

package policytemplate

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "PolicyTemplate" service endpoints.
type Endpoints struct {
	Compile      goa.Endpoint
	Upload       goa.Endpoint
	Update       goa.Endpoint
	RetrieveData goa.Endpoint
	Show         goa.Endpoint
	Index        goa.Endpoint
	Delete       goa.Endpoint
}

// NewEndpoints wraps the methods of the "PolicyTemplate" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Compile:      NewCompileEndpoint(s, a.JWTAuth),
		Upload:       NewUploadEndpoint(s, a.JWTAuth),
		Update:       NewUpdateEndpoint(s, a.JWTAuth),
		RetrieveData: NewRetrieveDataEndpoint(s, a.JWTAuth),
		Show:         NewShowEndpoint(s, a.JWTAuth),
		Index:        NewIndexEndpoint(s, a.JWTAuth),
		Delete:       NewDeleteEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "PolicyTemplate" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Compile = m(e.Compile)
	e.Upload = m(e.Upload)
	e.Update = m(e.Update)
	e.RetrieveData = m(e.RetrieveData)
	e.Show = m(e.Show)
	e.Index = m(e.Index)
	e.Delete = m(e.Delete)
}

// NewCompileEndpoint returns an endpoint function that calls the method
// "compile" of service "PolicyTemplate".
func NewCompileEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CompilePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:policy_template:compile||common:project:own||common:org:own"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Compile(ctx, p)
	}
}

// NewUploadEndpoint returns an endpoint function that calls the method
// "upload" of service "PolicyTemplate".
func NewUploadEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UploadPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:policy_template:upload||common:project:own||common:org:own"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.Upload(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPolicyTemplate(res, view)
		return vres, nil
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "PolicyTemplate".
func NewUpdateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:policy_template:update||common:project:own||common:org:own"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.Update(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPolicyTemplate(res, view)
		return vres, nil
	}
}

// NewRetrieveDataEndpoint returns an endpoint function that calls the method
// "retrieve_data" of service "PolicyTemplate".
func NewRetrieveDataEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RetrieveDataPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:policy_template:retrieve_data||common:project:own||common:org:own"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.RetrieveData(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "PolicyTemplate".
func NewShowEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:policy_template:show||common:project:own||common:org:own"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPolicyTemplate(res, view)
		return vres, nil
	}
}

// NewIndexEndpoint returns an endpoint function that calls the method "index"
// of service "PolicyTemplate".
func NewIndexEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*IndexPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:policy_template:index||common:project:own||common:org:own"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.Index(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPolicyTemplateList(res, view)
		return vres, nil
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "PolicyTemplate".
func NewDeleteEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeletePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:policy_template:delete||common:project:own||common:org:own"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Delete(ctx, p)
	}
}
