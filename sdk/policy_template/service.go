// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// PolicyTemplate service
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package policytemplate

import (
	"context"

	policytemplateviews "github.com/rightscale/right_pt/sdk/policy_template/views"
	"goa.design/goa"
	"goa.design/goa/security"
)

// Service is the PolicyTemplate service interface.
type Service interface {
	// Compile compiles a policy template for a project. This is only to be used
	// for checking the syntax of a policy template; the results are not stored.
	Compile(context.Context, *CompilePayload) (err error)
	// Upload uploads a policy template for a project, first compiling it. On
	// failure, an array of syntax errors will be returned.
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "extended"
	//	- "source"
	//	- "link"
	Upload(context.Context, *UploadPayload) (res *PolicyTemplate, view string, err error)
	// Update updates a policy template in place for a project, by replacing it.
	// Any existing applied policies using the template will not be updated; they
	// must be deleted and recreated again.
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "extended"
	//	- "source"
	//	- "link"
	Update(context.Context, *UpdatePayload) (res *PolicyTemplate, view string, err error)
	// Retrieve Data retrieves the data sources specified in a give policy template.
	RetrieveData(context.Context, *RetrieveDataPayload) (res []*Data, err error)
	// Show retrieves the details of a policy template.
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "extended"
	//	- "source"
	//	- "link"
	Show(context.Context, *ShowPayload) (res *PolicyTemplate, view string, err error)
	// IndexPolicyTemplates retrieves the list of policy templates in a project.
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "extended"
	Index(context.Context, *IndexPayload) (res *PolicyTemplateList, view string, err error)
	// Delete deletes a policy template from a project. Deleting a policy template
	// will not delete any applied policies created from the template, they must be
	// stopped explicitly.
	Delete(context.Context, *DeletePayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "PolicyTemplate"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [7]string{"compile", "upload", "update", "retrieve_data", "show", "index", "delete"}

// CompilePayload is the payload type of the PolicyTemplate service compile
// method.
type CompilePayload struct {
	// project_id identifies a project by ID.
	ProjectID uint
	// filename is the name of the policy template source code file.
	Filename string
	// source is the policy template source code.
	Source string
	// JWT token used to perform authorization
	Token *string
	// API Version, must be specified using this header
	APIVersion string
}

// UploadPayload is the payload type of the PolicyTemplate service upload
// method.
type UploadPayload struct {
	// project_id identifies a project by ID.
	ProjectID uint
	// filename is the name of the policy template source code file.
	Filename string
	// source is the policy template source code.
	Source string
	// JWT token used to perform authorization
	Token *string
	// API Version, must be specified using this header
	APIVersion string
}

// PolicyTemplate is the result type of the PolicyTemplate service upload
// method.
type PolicyTemplate struct {
	// id identifies a policy template by ID.
	ID string
	// name is the unique name of the policy template in the project.
	Name string
	// project_id is the ID of the project that the policy template applies to.
	ProjectID *uint
	// rs_pt_ver is the policy engine version.
	RsPtVer *uint
	// short_description is the short description of the policy template.
	ShortDescription *string
	// long_description is the long description of the policy template. The content
	// can be markdown.
	LongDescription *string
	// doc_link is an HTTP URI to a page containing detailed documentation for the
	// policy.
	DocLink *string
	// info is an arbitrary set of key/value pairs that provide additional
	// information such as the policy author, support contact information, etc.
	Info map[string]string
	// href is the href of the policy template.
	Href string
	// filename is the name of the file that was uploaded to create the policy
	// template.
	Filename *string
	// source is the policy template source code.
	Source *string
	// fingerprint is a SHA created during compilation. It is used to determine if
	// the template is outdated.
	Fingerprint string
	// category is the type categorization of the policy template.
	Category *string
	// created_by is the RightScale user that created the policy template.
	CreatedBy *User
	// created_at is the policy template creation timestamp in RFC3339 format.
	CreatedAt *string
	// updated_by is the RightScale user that updated the policy template.
	UpdatedBy *User
	// updated_at is the last update timestamp in RFC3339 format.
	UpdatedAt *string
	// permissions is a list of permissions required to run the policy.
	Permissions map[string]*Permission
	// required_roles is a list of governance roles, derived from permissions,
	// required to run the policy.
	RequiredRoles []string
	// parameters is a list of parameters required to apply the policy.
	Parameters map[string]*Parameter
	// severity defines the severity level of incidents raised from this policy
	// template.
	Severity *string
	// kind is "gov#policy_template".
	Kind string
}

// UpdatePayload is the payload type of the PolicyTemplate service update
// method.
type UpdatePayload struct {
	// project_id identifies a project by ID.
	ProjectID uint
	// template_id identifies a policy template by ID.
	TemplateID string
	// filename is the name of the policy template source code file.
	Filename string
	// source is the policy template source code.
	Source string
	// JWT token used to perform authorization
	Token *string
	// API Version, must be specified using this header
	APIVersion string
}

// RetrieveDataPayload is the payload type of the PolicyTemplate service
// retrieve_data method.
type RetrieveDataPayload struct {
	// project_id identifies a project by ID.
	ProjectID uint
	// template_id identifies a policy template by ID.
	TemplateID string
	// options lists the configuration options used to parameterize the policy.
	Options []*ConfigurationOptionCreateType
	// names is a filter to only retrieve datasources or resources that match the
	// given names
	Names []string
	// JWT token used to perform authorization
	Token *string
	// API Version, must be specified using this header
	APIVersion string
}

// ShowPayload is the payload type of the PolicyTemplate service show method.
type ShowPayload struct {
	// project_id identifies a project by ID.
	ProjectID uint
	// template_id identifies a policy template by ID.
	TemplateID string
	// View used to render policy template
	View *string
	// JWT token used to perform authorization
	Token *string
	// API Version, must be specified using this header
	APIVersion string
}

// IndexPayload is the payload type of the PolicyTemplate service index method.
type IndexPayload struct {
	// project_id identifies a project by ID.
	ProjectID uint
	// etag is an HTTP ETag. It is typically the previous ETag value retrieved by
	// client if any. Service returns an empty response with HTTP status code 304
	// Not Modified if value matches current value server side.
	Etag *string
	// View used to render policy templates
	View *string
	// JWT token used to perform authorization
	Token *string
	// API Version, must be specified using this header
	APIVersion string
}

// PolicyTemplateList is the result type of the PolicyTemplate service index
// method.
type PolicyTemplateList struct {
	// count is the number of policy templates in the list.
	Count *uint
	// etag is an HTTP ETag for the policy template list.
	Etag string
	// items is the array of policy templates.
	Items PolicyTemplateCollection
	// not_modified is a flag used internally that indicates how to encode the HTTP
	// response (i.e. 200 or 304).
	NotModified *string
	// kind is "gov#policy_template_list".
	Kind string
}

// DeletePayload is the payload type of the PolicyTemplate service delete
// method.
type DeletePayload struct {
	// project_id identifies a project by ID.
	ProjectID uint
	// template_id identifies a policy template by ID.
	TemplateID string
	// JWT token used to perform authorization
	Token *string
	// API Version, must be specified using this header
	APIVersion string
}

// User represents a registered RightScale user.
type User struct {
	// ID of user
	ID uint
	// email of user
	Email string
	// name of user, usually of the form 'First Last'
	Name string
}

// Permission defines a role required in RightScale to perform actions on
// resources
type Permission struct {
	// Name of a permission
	Name string `json:"name"`
	// Label is used in the UI
	Label *string `json:"label"`
	// List of resource names the permission is applied to
	Resources []string `json:"resources"`
	// List of action names the permission is applied to
	Actions []string `json:"actions"`
}

// Parameter defines a parameter given as input to a Policy
type Parameter struct {
	// Name of the parameter
	Name string `json:"name"`
	// Type of the parameter
	Type string `json:"type"`
	// Label to show in the UI
	Label string `json:"label"`
	// The index of this parameter in the list
	Index uint `json:"index"`
	// The category used to group parameters
	Category *string `json:"category"`
	// Description of the parameter
	Description *string `json:"description"`
	// The default value for the parameter
	Default interface{} `json:"default"`
	// no_echo determines whether the value of the parameter should be hidden in
	// UIs and API responses.
	NoEcho bool `json:"no_echo"`
	// List of values allowed for this parameter
	AllowedValues []interface{} `json:"allowed_values"`
	// The minimum length of a string parameter
	MinLength *uint `json:"min_length"`
	// The maximum length of a string parameter
	MaxLength *uint `json:"max_length"`
	// The minimum value of a number parmameter
	MinValue *float64 `json:"min_value"`
	// The maximum value of a number parameter
	MaxValue *float64 `json:"max_value"`
	// The regular expression pattern used to validate a string parameter
	AllowedPattern *Regexp `json:"allowed_pattern"`
	// The description used for constraints
	ConstraintDescription *string `json:"constraint_description"`
}

// Regular expression
type Regexp struct {
	// Pattern is the regular expression pattern.
	Pattern string `json:"pattern"`
	// Options are the regular expression options. Options i (case insensitve) and
	// m (match over newlines) supported.
	Options *string `json:"options"`
}

// ConfigurationOptionCreateType is the payload for creating a single parameter
// value used to configure an applied policy.
type ConfigurationOptionCreateType struct {
	// name of option
	Name string
	// value of option
	Value interface{}
}

// Data contains retrieved datasource or resource data.
type Data struct {
	// name is the unique name of the datasource.
	Name string
	// type is is either Resource or Datasource
	Type string
	// is the extracted data
	Data interface{}
}

type PolicyTemplateCollection []*PolicyTemplate

// Errors occurred during compilation.
type CompilationErrors struct {
	// id is a unique identifier for this particular occurrence of the problem.
	ID string
	// errors is a JSON Array of compilation errors.
	Errors []*CompilationError
	// name of error
	Name string
}

// CompilationError is a single compilation error.
type CompilationError struct {
	// origin includes the template name and line number which generated the error.
	Origin string
	// problem is the kind of error.
	Problem string
	// summary includes detailed error information.
	Summary string
	// resolutions indicates how to resolve the error.
	Resolution string
}

// A template with this name already exists.
type ConflictError struct {
	// a unique identifier for this particular occurrence of the problem.
	ID string
	// name of error.
	Name     string
	Location string
	// updatable indicates whether a 'create' request error can be resolved by
	// making an 'update' request instead. It is used to indicate whether a
	// resource is built-in and whether the user has permission to modify it.
	Updatable bool
}

// Error returns an error description.
func (e *CompilationErrors) Error() string {
	return "Errors occurred during compilation."
}

// ErrorName returns "CompilationErrors".
func (e *CompilationErrors) ErrorName() string {
	return e.Name
}

// Error returns an error description.
func (e *CompilationError) Error() string {
	return "CompilationError is a single compilation error."
}

// ErrorName returns "CompilationError".
func (e *CompilationError) ErrorName() string {
	return "CompilationError"
}

// Error returns an error description.
func (e *ConflictError) Error() string {
	return "A template with this name already exists."
}

// ErrorName returns "ConflictError".
func (e *ConflictError) ErrorName() string {
	return e.Name
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeForbidden builds a goa.ServiceError from an error.
func MakeForbidden(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "forbidden",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad_request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadGateway builds a goa.ServiceError from an error.
func MakeBadGateway(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad_gateway",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal_error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeUnprocessableEntity builds a goa.ServiceError from an error.
func MakeUnprocessableEntity(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unprocessable_entity",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewPolicyTemplate initializes result type PolicyTemplate from viewed result
// type PolicyTemplate.
func NewPolicyTemplate(vres *policytemplateviews.PolicyTemplate) *PolicyTemplate {
	var res *PolicyTemplate
	switch vres.View {
	case "default", "":
		res = newPolicyTemplate(vres.Projected)
	case "extended":
		res = newPolicyTemplateExtended(vres.Projected)
	case "source":
		res = newPolicyTemplateSource(vres.Projected)
	case "link":
		res = newPolicyTemplateLink(vres.Projected)
	}
	return res
}

// NewViewedPolicyTemplate initializes viewed result type PolicyTemplate from
// result type PolicyTemplate using the given view.
func NewViewedPolicyTemplate(res *PolicyTemplate, view string) *policytemplateviews.PolicyTemplate {
	var vres *policytemplateviews.PolicyTemplate
	switch view {
	case "default", "":
		p := newPolicyTemplateView(res)
		vres = &policytemplateviews.PolicyTemplate{p, "default"}
	case "extended":
		p := newPolicyTemplateViewExtended(res)
		vres = &policytemplateviews.PolicyTemplate{p, "extended"}
	case "source":
		p := newPolicyTemplateViewSource(res)
		vres = &policytemplateviews.PolicyTemplate{p, "source"}
	case "link":
		p := newPolicyTemplateViewLink(res)
		vres = &policytemplateviews.PolicyTemplate{p, "link"}
	}
	return vres
}

// NewPolicyTemplateList initializes result type PolicyTemplateList from viewed
// result type PolicyTemplateList.
func NewPolicyTemplateList(vres *policytemplateviews.PolicyTemplateList) *PolicyTemplateList {
	var res *PolicyTemplateList
	switch vres.View {
	case "default", "":
		res = newPolicyTemplateList(vres.Projected)
	case "extended":
		res = newPolicyTemplateListExtended(vres.Projected)
	}
	return res
}

// NewViewedPolicyTemplateList initializes viewed result type
// PolicyTemplateList from result type PolicyTemplateList using the given view.
func NewViewedPolicyTemplateList(res *PolicyTemplateList, view string) *policytemplateviews.PolicyTemplateList {
	var vres *policytemplateviews.PolicyTemplateList
	switch view {
	case "default", "":
		p := newPolicyTemplateListView(res)
		vres = &policytemplateviews.PolicyTemplateList{p, "default"}
	case "extended":
		p := newPolicyTemplateListViewExtended(res)
		vres = &policytemplateviews.PolicyTemplateList{p, "extended"}
	}
	return vres
}

// newPolicyTemplate converts projected type PolicyTemplate to service type
// PolicyTemplate.
func newPolicyTemplate(vres *policytemplateviews.PolicyTemplateView) *PolicyTemplate {
	res := &PolicyTemplate{
		ProjectID:        vres.ProjectID,
		RsPtVer:          vres.RsPtVer,
		ShortDescription: vres.ShortDescription,
		DocLink:          vres.DocLink,
		Category:         vres.Category,
		CreatedAt:        vres.CreatedAt,
		UpdatedAt:        vres.UpdatedAt,
		Severity:         vres.Severity,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Href != nil {
		res.Href = *vres.Href
	}
	if vres.Fingerprint != nil {
		res.Fingerprint = *vres.Fingerprint
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	if vres.Info != nil {
		res.Info = make(map[string]string, len(vres.Info))
		for key, val := range vres.Info {
			tk := key
			tv := val
			res.Info[tk] = tv
		}
	}
	if vres.CreatedBy != nil {
		res.CreatedBy = unmarshalUserViewToUser(vres.CreatedBy)
	}
	if vres.UpdatedBy != nil {
		res.UpdatedBy = unmarshalUserViewToUser(vres.UpdatedBy)
	}
	if vres.RequiredRoles != nil {
		res.RequiredRoles = make([]string, len(vres.RequiredRoles))
		for i, val := range vres.RequiredRoles {
			res.RequiredRoles[i] = val
		}
	}
	return res
}

// newPolicyTemplateExtended converts projected type PolicyTemplate to service
// type PolicyTemplate.
func newPolicyTemplateExtended(vres *policytemplateviews.PolicyTemplateView) *PolicyTemplate {
	res := &PolicyTemplate{
		ProjectID:        vres.ProjectID,
		RsPtVer:          vres.RsPtVer,
		ShortDescription: vres.ShortDescription,
		LongDescription:  vres.LongDescription,
		DocLink:          vres.DocLink,
		Category:         vres.Category,
		CreatedAt:        vres.CreatedAt,
		UpdatedAt:        vres.UpdatedAt,
		Severity:         vres.Severity,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Href != nil {
		res.Href = *vres.Href
	}
	if vres.Fingerprint != nil {
		res.Fingerprint = *vres.Fingerprint
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	if vres.Info != nil {
		res.Info = make(map[string]string, len(vres.Info))
		for key, val := range vres.Info {
			tk := key
			tv := val
			res.Info[tk] = tv
		}
	}
	if vres.CreatedBy != nil {
		res.CreatedBy = unmarshalUserViewToUser(vres.CreatedBy)
	}
	if vres.UpdatedBy != nil {
		res.UpdatedBy = unmarshalUserViewToUser(vres.UpdatedBy)
	}
	if vres.Permissions != nil {
		res.Permissions = make(map[string]*Permission, len(vres.Permissions))
		for key, val := range vres.Permissions {
			tk := key
			tv := &Permission{
				Name:  *val.Name,
				Label: val.Label,
			}
			tv.Resources = make([]string, len(val.Resources))
			for i, val := range val.Resources {
				tv.Resources[i] = val
			}
			tv.Actions = make([]string, len(val.Actions))
			for i, val := range val.Actions {
				tv.Actions[i] = val
			}
			res.Permissions[tk] = tv
		}
	}
	if vres.RequiredRoles != nil {
		res.RequiredRoles = make([]string, len(vres.RequiredRoles))
		for i, val := range vres.RequiredRoles {
			res.RequiredRoles[i] = val
		}
	}
	if vres.Parameters != nil {
		res.Parameters = make(map[string]*Parameter, len(vres.Parameters))
		for key, val := range vres.Parameters {
			tk := key
			tv := &Parameter{
				Name:                  *val.Name,
				Type:                  *val.Type,
				Label:                 *val.Label,
				Index:                 *val.Index,
				Category:              val.Category,
				Description:           val.Description,
				MinLength:             val.MinLength,
				MaxLength:             val.MaxLength,
				MinValue:              val.MinValue,
				MaxValue:              val.MaxValue,
				ConstraintDescription: val.ConstraintDescription,
			}
			if val.Default != nil {
				tv.Default = *val.Default
			}
			if val.NoEcho != nil {
				tv.NoEcho = *val.NoEcho
			}
			if val.NoEcho == nil {
				tv.NoEcho = false
			}
			if val.AllowedValues != nil {
				tv.AllowedValues = make([]interface{}, len(val.AllowedValues))
				for i, val := range val.AllowedValues {
					tv.AllowedValues[i] = val
				}
			}
			if val.AllowedPattern != nil {
				tv.AllowedPattern = unmarshalRegexpViewToRegexp(val.AllowedPattern)
			}
			res.Parameters[tk] = tv
		}
	}
	return res
}

// newPolicyTemplateSource converts projected type PolicyTemplate to service
// type PolicyTemplate.
func newPolicyTemplateSource(vres *policytemplateviews.PolicyTemplateView) *PolicyTemplate {
	res := &PolicyTemplate{
		Filename: vres.Filename,
		Source:   vres.Source,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Href != nil {
		res.Href = *vres.Href
	}
	if vres.Fingerprint != nil {
		res.Fingerprint = *vres.Fingerprint
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	return res
}

// newPolicyTemplateLink converts projected type PolicyTemplate to service type
// PolicyTemplate.
func newPolicyTemplateLink(vres *policytemplateviews.PolicyTemplateView) *PolicyTemplate {
	res := &PolicyTemplate{
		UpdatedAt: vres.UpdatedAt,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Href != nil {
		res.Href = *vres.Href
	}
	if vres.Fingerprint != nil {
		res.Fingerprint = *vres.Fingerprint
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	if vres.UpdatedBy != nil {
		res.UpdatedBy = unmarshalUserViewToUser(vres.UpdatedBy)
	}
	return res
}

// newPolicyTemplateView projects result type PolicyTemplate into projected
// type PolicyTemplateView using the "default" view.
func newPolicyTemplateView(res *PolicyTemplate) *policytemplateviews.PolicyTemplateView {
	vres := &policytemplateviews.PolicyTemplateView{
		ID:               &res.ID,
		Name:             &res.Name,
		ProjectID:        res.ProjectID,
		RsPtVer:          res.RsPtVer,
		ShortDescription: res.ShortDescription,
		DocLink:          res.DocLink,
		Href:             &res.Href,
		Fingerprint:      &res.Fingerprint,
		Category:         res.Category,
		CreatedAt:        res.CreatedAt,
		UpdatedAt:        res.UpdatedAt,
		Severity:         res.Severity,
		Kind:             &res.Kind,
	}
	if res.Info != nil {
		vres.Info = make(map[string]string, len(res.Info))
		for key, val := range res.Info {
			tk := key
			tv := val
			vres.Info[tk] = tv
		}
	}
	if res.CreatedBy != nil {
		vres.CreatedBy = marshalUserToUserView(res.CreatedBy)
	}
	if res.UpdatedBy != nil {
		vres.UpdatedBy = marshalUserToUserView(res.UpdatedBy)
	}
	if res.RequiredRoles != nil {
		vres.RequiredRoles = make([]string, len(res.RequiredRoles))
		for i, val := range res.RequiredRoles {
			vres.RequiredRoles[i] = val
		}
	}
	return vres
}

// newPolicyTemplateViewExtended projects result type PolicyTemplate into
// projected type PolicyTemplateView using the "extended" view.
func newPolicyTemplateViewExtended(res *PolicyTemplate) *policytemplateviews.PolicyTemplateView {
	vres := &policytemplateviews.PolicyTemplateView{
		ID:               &res.ID,
		Name:             &res.Name,
		ProjectID:        res.ProjectID,
		RsPtVer:          res.RsPtVer,
		ShortDescription: res.ShortDescription,
		LongDescription:  res.LongDescription,
		DocLink:          res.DocLink,
		Href:             &res.Href,
		Fingerprint:      &res.Fingerprint,
		Category:         res.Category,
		CreatedAt:        res.CreatedAt,
		UpdatedAt:        res.UpdatedAt,
		Severity:         res.Severity,
		Kind:             &res.Kind,
	}
	if res.Info != nil {
		vres.Info = make(map[string]string, len(res.Info))
		for key, val := range res.Info {
			tk := key
			tv := val
			vres.Info[tk] = tv
		}
	}
	if res.CreatedBy != nil {
		vres.CreatedBy = marshalUserToUserView(res.CreatedBy)
	}
	if res.UpdatedBy != nil {
		vres.UpdatedBy = marshalUserToUserView(res.UpdatedBy)
	}
	if res.Permissions != nil {
		vres.Permissions = make(map[string]*policytemplateviews.PermissionView, len(res.Permissions))
		for key, val := range res.Permissions {
			tk := key
			tv := &policytemplateviews.PermissionView{
				Name:  &val.Name,
				Label: val.Label,
			}
			if val.Resources != nil {
				tv.Resources = make([]string, len(val.Resources))
				for i, val := range val.Resources {
					tv.Resources[i] = val
				}
			}
			if val.Actions != nil {
				tv.Actions = make([]string, len(val.Actions))
				for i, val := range val.Actions {
					tv.Actions[i] = val
				}
			}
			vres.Permissions[tk] = tv
		}
	}
	if res.RequiredRoles != nil {
		vres.RequiredRoles = make([]string, len(res.RequiredRoles))
		for i, val := range res.RequiredRoles {
			vres.RequiredRoles[i] = val
		}
	}
	if res.Parameters != nil {
		vres.Parameters = make(map[string]*policytemplateviews.ParameterView, len(res.Parameters))
		for key, val := range res.Parameters {
			tk := key
			tv := &policytemplateviews.ParameterView{
				Name:                  &val.Name,
				Type:                  &val.Type,
				Label:                 &val.Label,
				Index:                 &val.Index,
				Category:              val.Category,
				Description:           val.Description,
				Default:               &val.Default,
				NoEcho:                &val.NoEcho,
				MinLength:             val.MinLength,
				MaxLength:             val.MaxLength,
				MinValue:              val.MinValue,
				MaxValue:              val.MaxValue,
				ConstraintDescription: val.ConstraintDescription,
			}
			if val.AllowedValues != nil {
				tv.AllowedValues = make([]interface{}, len(val.AllowedValues))
				for i, val := range val.AllowedValues {
					tv.AllowedValues[i] = val
				}
			}
			if val.AllowedPattern != nil {
				tv.AllowedPattern = marshalRegexpToRegexpView(val.AllowedPattern)
			}
			vres.Parameters[tk] = tv
		}
	}
	return vres
}

// newPolicyTemplateViewSource projects result type PolicyTemplate into
// projected type PolicyTemplateView using the "source" view.
func newPolicyTemplateViewSource(res *PolicyTemplate) *policytemplateviews.PolicyTemplateView {
	vres := &policytemplateviews.PolicyTemplateView{
		ID:          &res.ID,
		Name:        &res.Name,
		Href:        &res.Href,
		Filename:    res.Filename,
		Source:      res.Source,
		Fingerprint: &res.Fingerprint,
		Kind:        &res.Kind,
	}
	return vres
}

// newPolicyTemplateViewLink projects result type PolicyTemplate into projected
// type PolicyTemplateView using the "link" view.
func newPolicyTemplateViewLink(res *PolicyTemplate) *policytemplateviews.PolicyTemplateView {
	vres := &policytemplateviews.PolicyTemplateView{
		ID:          &res.ID,
		Name:        &res.Name,
		Href:        &res.Href,
		Fingerprint: &res.Fingerprint,
		UpdatedAt:   res.UpdatedAt,
		Kind:        &res.Kind,
	}
	if res.UpdatedBy != nil {
		vres.UpdatedBy = marshalUserToUserView(res.UpdatedBy)
	}
	return vres
}

// newPolicyTemplateList converts projected type PolicyTemplateList to service
// type PolicyTemplateList.
func newPolicyTemplateList(vres *policytemplateviews.PolicyTemplateListView) *PolicyTemplateList {
	res := &PolicyTemplateList{
		Count:       vres.Count,
		NotModified: vres.NotModified,
	}
	if vres.Etag != nil {
		res.Etag = *vres.Etag
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	if vres.Items != nil {
		res.Items = newPolicyTemplateCollection(vres.Items)
	}
	return res
}

// newPolicyTemplateListExtended converts projected type PolicyTemplateList to
// service type PolicyTemplateList.
func newPolicyTemplateListExtended(vres *policytemplateviews.PolicyTemplateListView) *PolicyTemplateList {
	res := &PolicyTemplateList{
		Count:       vres.Count,
		NotModified: vres.NotModified,
	}
	if vres.Etag != nil {
		res.Etag = *vres.Etag
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	if vres.Items != nil {
		res.Items = newPolicyTemplateCollectionExtended(vres.Items)
	}
	return res
}

// newPolicyTemplateListView projects result type PolicyTemplateList into
// projected type PolicyTemplateListView using the "default" view.
func newPolicyTemplateListView(res *PolicyTemplateList) *policytemplateviews.PolicyTemplateListView {
	vres := &policytemplateviews.PolicyTemplateListView{
		Count:       res.Count,
		Etag:        &res.Etag,
		NotModified: res.NotModified,
		Kind:        &res.Kind,
	}
	if res.Items != nil {
		vres.Items = newPolicyTemplateCollectionView(res.Items)
	}
	return vres
}

// newPolicyTemplateListViewExtended projects result type PolicyTemplateList
// into projected type PolicyTemplateListView using the "extended" view.
func newPolicyTemplateListViewExtended(res *PolicyTemplateList) *policytemplateviews.PolicyTemplateListView {
	vres := &policytemplateviews.PolicyTemplateListView{
		Count:       res.Count,
		Etag:        &res.Etag,
		NotModified: res.NotModified,
		Kind:        &res.Kind,
	}
	if res.Items != nil {
		vres.Items = newPolicyTemplateCollectionViewExtended(res.Items)
	}
	return vres
}

// newPolicyTemplateCollection converts projected type PolicyTemplateCollection
// to service type PolicyTemplateCollection.
func newPolicyTemplateCollection(vres policytemplateviews.PolicyTemplateCollectionView) PolicyTemplateCollection {
	res := make(PolicyTemplateCollection, len(vres))
	for i, n := range vres {
		res[i] = newPolicyTemplate(n)
	}
	return res
}

// newPolicyTemplateCollectionExtended converts projected type
// PolicyTemplateCollection to service type PolicyTemplateCollection.
func newPolicyTemplateCollectionExtended(vres policytemplateviews.PolicyTemplateCollectionView) PolicyTemplateCollection {
	res := make(PolicyTemplateCollection, len(vres))
	for i, n := range vres {
		res[i] = newPolicyTemplateExtended(n)
	}
	return res
}

// newPolicyTemplateCollectionSource converts projected type
// PolicyTemplateCollection to service type PolicyTemplateCollection.
func newPolicyTemplateCollectionSource(vres policytemplateviews.PolicyTemplateCollectionView) PolicyTemplateCollection {
	res := make(PolicyTemplateCollection, len(vres))
	for i, n := range vres {
		res[i] = newPolicyTemplateSource(n)
	}
	return res
}

// newPolicyTemplateCollectionLink converts projected type
// PolicyTemplateCollection to service type PolicyTemplateCollection.
func newPolicyTemplateCollectionLink(vres policytemplateviews.PolicyTemplateCollectionView) PolicyTemplateCollection {
	res := make(PolicyTemplateCollection, len(vres))
	for i, n := range vres {
		res[i] = newPolicyTemplateLink(n)
	}
	return res
}

// newPolicyTemplateCollectionView projects result type
// PolicyTemplateCollection into projected type PolicyTemplateCollectionView
// using the "default" view.
func newPolicyTemplateCollectionView(res PolicyTemplateCollection) policytemplateviews.PolicyTemplateCollectionView {
	vres := make(policytemplateviews.PolicyTemplateCollectionView, len(res))
	for i, n := range res {
		vres[i] = newPolicyTemplateView(n)
	}
	return vres
}

// newPolicyTemplateCollectionViewExtended projects result type
// PolicyTemplateCollection into projected type PolicyTemplateCollectionView
// using the "extended" view.
func newPolicyTemplateCollectionViewExtended(res PolicyTemplateCollection) policytemplateviews.PolicyTemplateCollectionView {
	vres := make(policytemplateviews.PolicyTemplateCollectionView, len(res))
	for i, n := range res {
		vres[i] = newPolicyTemplateViewExtended(n)
	}
	return vres
}

// newPolicyTemplateCollectionViewSource projects result type
// PolicyTemplateCollection into projected type PolicyTemplateCollectionView
// using the "source" view.
func newPolicyTemplateCollectionViewSource(res PolicyTemplateCollection) policytemplateviews.PolicyTemplateCollectionView {
	vres := make(policytemplateviews.PolicyTemplateCollectionView, len(res))
	for i, n := range res {
		vres[i] = newPolicyTemplateViewSource(n)
	}
	return vres
}

// newPolicyTemplateCollectionViewLink projects result type
// PolicyTemplateCollection into projected type PolicyTemplateCollectionView
// using the "link" view.
func newPolicyTemplateCollectionViewLink(res PolicyTemplateCollection) policytemplateviews.PolicyTemplateCollectionView {
	vres := make(policytemplateviews.PolicyTemplateCollectionView, len(res))
	for i, n := range res {
		vres[i] = newPolicyTemplateViewLink(n)
	}
	return vres
}

// unmarshalUserViewToUser builds a value of type *User from a value of type
// *policytemplateviews.UserView.
func unmarshalUserViewToUser(v *policytemplateviews.UserView) *User {
	if v == nil {
		return nil
	}
	res := &User{
		ID:    *v.ID,
		Email: *v.Email,
		Name:  *v.Name,
	}

	return res
}

// unmarshalRegexpViewToRegexp builds a value of type *Regexp from a value of
// type *policytemplateviews.RegexpView.
func unmarshalRegexpViewToRegexp(v *policytemplateviews.RegexpView) *Regexp {
	if v == nil {
		return nil
	}
	res := &Regexp{
		Pattern: *v.Pattern,
		Options: v.Options,
	}

	return res
}

// marshalUserToUserView builds a value of type *policytemplateviews.UserView
// from a value of type *User.
func marshalUserToUserView(v *User) *policytemplateviews.UserView {
	if v == nil {
		return nil
	}
	res := &policytemplateviews.UserView{
		ID:    &v.ID,
		Email: &v.Email,
		Name:  &v.Name,
	}

	return res
}

// marshalRegexpToRegexpView builds a value of type
// *policytemplateviews.RegexpView from a value of type *Regexp.
func marshalRegexpToRegexpView(v *Regexp) *policytemplateviews.RegexpView {
	if v == nil {
		return nil
	}
	res := &policytemplateviews.RegexpView{
		Pattern: &v.Pattern,
		Options: v.Options,
	}

	return res
}
