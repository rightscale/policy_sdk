// Code generated by goa v3.1.3, DO NOT EDIT.
//
// PublishedTemplate views
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design -o .

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// PublishedTemplate is the viewed result type that is projected based on a
// view.
type PublishedTemplate struct {
	// Type to project
	Projected *PublishedTemplateView
	// View to render
	View string
}

// PublishedTemplateList is the viewed result type that is projected based on a
// view.
type PublishedTemplateList struct {
	// Type to project
	Projected *PublishedTemplateListView
	// View to render
	View string
}

// PublishedTemplateView is a type that runs validations on a projected type.
type PublishedTemplateView struct {
	// id identifies a published template by ID.
	ID *string
	// name is the unique name of the published template in the organization.
	Name *string
	// org_id is the ID of the organization that the published template belongs to.
	OrgID *uint
	// project_id is the ID of the project that the published template is published
	// from.
	ProjectID *uint
	// policy_template_id is the ID of the policy template from which the published
	// template originated.
	PolicyTemplateID *string
	// policy_template_url is the full URL to the policy template from which the
	// published template originated.
	PolicyTemplateURL *string
	// policy_template_fingerprint is fingerprint of the policy template. It is
	// used to determine if the policy template that this was published from is
	// outdated.
	PolicyTemplateFingerprint *string
	// rs_pt_ver is the policy engine version.
	RsPtVer *uint
	// short_description is the short description of the published template.
	ShortDescription *string
	// long_description is the long description of the published template. The
	// content can be markdown.
	LongDescription *string
	// doc_link is a HTTP URI to a page containing detailed documentation for the
	// policy.
	DocLink *string
	// info is an arbitrary set of key/value pairs that provide additional
	// information such as the policy author, support contact information, etc.
	Info map[string]string
	// default_frequency defines the interval the template will be run unless set
	// differently during application.
	DefaultFrequency *string
	// href is the self-referential href of the published template.
	Href *string
	// filename is the name of the file that was uploaded to create the policy
	// template.
	Filename *string
	// source is published template source code.
	Source *string
	// fingerprint is a SHA created during compilation. It is used to determine if
	// the policy template that this was published from is outdated.
	Fingerprint *string
	// category is the type categorization of the published template.
	Category *string
	// created_by is the RightScale user that created the published template.
	CreatedBy *UserView
	// created_at is the published template creation timestamp in RFC3339 format.
	CreatedAt *string
	// updated_by is the RightScale user that updated the published template.
	UpdatedBy *UserView
	// updated_at is the published template update timestamp in RFC3339 format.
	UpdatedAt *string
	// permissions is a list of permissions required to run the policy.
	Permissions map[string]*PermissionView
	// required_roles is a list of governance roles, derived from permissions,
	// required to run the policy.
	RequiredRoles []string
	// parameters is a list of parameters required to apply the policy.
	Parameters map[string]*ParameterView
	// severity defines the severity level of incidents raised from this published
	// template.
	Severity *string
	// built_in is a flag to indicate whether the published template is a
	// "built-in" RS-supplied template.
	BuiltIn *bool
	// hidden is a flag to indicate whether the published template is hidden.
	Hidden *bool
	// hidden_by is the RightScale user that marked the published template as
	// hidden.
	HiddenBy *UserView
	// hidden_at is the hidden at timestamp in RFC3339 format.
	HiddenAt *string
	// tenancy indicates whether this template can be run across multiple projects
	// or is restricted to a single project.
	Tenancy *string
	// credentials is a list of authorization for external APIs needed to run the
	// policy.
	Credentials map[string]*CredentialsView
	// kind is "gov#published_template".
	Kind *string
}

// UserView is a type that runs validations on a projected type.
type UserView struct {
	// ID of user
	ID *uint
	// email of user
	Email *string
	// name of user, usually of the form 'First Last'
	Name *string
}

// PermissionView is a type that runs validations on a projected type.
type PermissionView struct {
	// Name of a permission
	Name *string `json:"name"`
	// Label is used in the UI
	Label *string `json:"label"`
	// List of resource names the permission is applied to
	Resources []string `json:"resources"`
	// List of action names the permission is applied to
	Actions []string `json:"actions"`
}

// ParameterView is a type that runs validations on a projected type.
type ParameterView struct {
	// Name of the parameter
	Name *string `json:"name"`
	// Type of the parameter
	Type *string `json:"type"`
	// Label to show in the UI
	Label *string `json:"label"`
	// The index of this parameter in the list
	Index *uint `json:"index"`
	// The category used to group parameters
	Category *string `json:"category"`
	// Description of the parameter
	Description *string `json:"description"`
	// The default value for the parameter
	Default interface{} `json:"default"`
	// no_echo determines whether the value of the parameter should be hidden in
	// UIs and API responses.
	NoEcho *bool `json:"no_echo"`
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
	AllowedPattern *RegexpView `json:"allowed_pattern"`
	// The description used for constraints
	ConstraintDescription *string `json:"constraint_description"`
}

// RegexpView is a type that runs validations on a projected type.
type RegexpView struct {
	// Pattern is the regular expression pattern.
	Pattern *string `json:"pattern"`
	// Options are the regular expression options. Options i (case insensitve) and
	// m (match over newlines) supported.
	Options *string `json:"options"`
}

// CredentialsView is a type that runs validations on a projected type.
type CredentialsView struct {
	// Name in policy template source code
	Name *string `json:"name"`
	// Schemes of credentials service resource supported.
	Schemes []string `json:"schemes"`
	// Label for the auth reference
	Label *string `json:"label"`
	// Description of what types of permissions need to be provided by auth.
	Description *string `json:"description"`
	// Tags is an optional filter to only show credentials matching a certain tag.
	Tags []*CredentialsTagView `json:"tags"`
}

// CredentialsTagView is a type that runs validations on a projected type.
type CredentialsTagView struct {
	// Key is the tag key.
	Key *string `json:"key"`
	// Value is the tag value.
	Value *string `json:"value"`
}

// PublishedTemplateListView is a type that runs validations on a projected
// type.
type PublishedTemplateListView struct {
	// count is the number of published templates in the list.
	Count *uint
	// etag is an HTTP ETag for the published template list.
	Etag *string
	// items is the array of published templates.
	Items PublishedTemplateCollectionView
	// not_modified is a flag used internally that indicates how to encode the HTTP
	// response (i.e. 200 or 304).
	NotModified *string
	// kind is "gov#published_template_list".
	Kind *string
}

// PublishedTemplateCollectionView is a type that runs validations on a
// projected type.
type PublishedTemplateCollectionView []*PublishedTemplateView

var (
	// PublishedTemplateMap is a map of attribute names in result type
	// PublishedTemplate indexed by view name.
	PublishedTemplateMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"org_id",
			"project_id",
			"policy_template_id",
			"policy_template_url",
			"policy_template_fingerprint",
			"rs_pt_ver",
			"short_description",
			"doc_link",
			"info",
			"default_frequency",
			"href",
			"fingerprint",
			"category",
			"created_by",
			"created_at",
			"updated_by",
			"updated_at",
			"hidden_by",
			"hidden_at",
			"required_roles",
			"severity",
			"built_in",
			"hidden",
			"tenancy",
			"kind",
		},
		"extended": []string{
			"id",
			"name",
			"org_id",
			"project_id",
			"policy_template_id",
			"policy_template_url",
			"policy_template_fingerprint",
			"rs_pt_ver",
			"short_description",
			"long_description",
			"doc_link",
			"info",
			"default_frequency",
			"href",
			"fingerprint",
			"category",
			"created_by",
			"created_at",
			"updated_by",
			"updated_at",
			"hidden_by",
			"hidden_at",
			"permissions",
			"required_roles",
			"parameters",
			"severity",
			"built_in",
			"hidden",
			"tenancy",
			"credentials",
			"kind",
		},
		"source": []string{
			"id",
			"name",
			"href",
			"filename",
			"source",
			"fingerprint",
			"kind",
		},
		"link": []string{
			"id",
			"name",
			"href",
			"fingerprint",
			"updated_at",
			"updated_by",
			"built_in",
			"kind",
		},
	}
	// PublishedTemplateListMap is a map of attribute names in result type
	// PublishedTemplateList indexed by view name.
	PublishedTemplateListMap = map[string][]string{
		"default": []string{
			"count",
			"items",
			"etag",
			"not_modified",
			"kind",
		},
		"extended": []string{
			"count",
			"items",
			"etag",
			"not_modified",
			"kind",
		},
	}
	// PublishedTemplateCollectionMap is a map of attribute names in result type
	// PublishedTemplateCollection indexed by view name.
	PublishedTemplateCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"org_id",
			"project_id",
			"policy_template_id",
			"policy_template_url",
			"policy_template_fingerprint",
			"rs_pt_ver",
			"short_description",
			"doc_link",
			"info",
			"default_frequency",
			"href",
			"fingerprint",
			"category",
			"created_by",
			"created_at",
			"updated_by",
			"updated_at",
			"hidden_by",
			"hidden_at",
			"required_roles",
			"severity",
			"built_in",
			"hidden",
			"tenancy",
			"kind",
		},
		"extended": []string{
			"id",
			"name",
			"org_id",
			"project_id",
			"policy_template_id",
			"policy_template_url",
			"policy_template_fingerprint",
			"rs_pt_ver",
			"short_description",
			"long_description",
			"doc_link",
			"info",
			"default_frequency",
			"href",
			"fingerprint",
			"category",
			"created_by",
			"created_at",
			"updated_by",
			"updated_at",
			"hidden_by",
			"hidden_at",
			"permissions",
			"required_roles",
			"parameters",
			"severity",
			"built_in",
			"hidden",
			"tenancy",
			"credentials",
			"kind",
		},
		"source": []string{
			"id",
			"name",
			"href",
			"filename",
			"source",
			"fingerprint",
			"kind",
		},
		"link": []string{
			"id",
			"name",
			"href",
			"fingerprint",
			"updated_at",
			"updated_by",
			"built_in",
			"kind",
		},
	}
)

// ValidatePublishedTemplate runs the validations defined on the viewed result
// type PublishedTemplate.
func ValidatePublishedTemplate(result *PublishedTemplate) (err error) {
	switch result.View {
	case "default", "":
		err = ValidatePublishedTemplateView(result.Projected)
	case "extended":
		err = ValidatePublishedTemplateViewExtended(result.Projected)
	case "source":
		err = ValidatePublishedTemplateViewSource(result.Projected)
	case "link":
		err = ValidatePublishedTemplateViewLink(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "extended", "source", "link"})
	}
	return
}

// ValidatePublishedTemplateList runs the validations defined on the viewed
// result type PublishedTemplateList.
func ValidatePublishedTemplateList(result *PublishedTemplateList) (err error) {
	switch result.View {
	case "default", "":
		err = ValidatePublishedTemplateListView(result.Projected)
	case "extended":
		err = ValidatePublishedTemplateListViewExtended(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "extended"})
	}
	return
}

// ValidatePublishedTemplateView runs the validations defined on
// PublishedTemplateView using the "default" view.
func ValidatePublishedTemplateView(result *PublishedTemplateView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	if result.Fingerprint == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fingerprint", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.PolicyTemplateURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.policy_template_url", *result.PolicyTemplateURL, goa.FormatURI))
	}
	if result.DocLink != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.doc_link", *result.DocLink, goa.FormatURI))
	}
	if result.DefaultFrequency != nil {
		if !(*result.DefaultFrequency == "15 minutes" || *result.DefaultFrequency == "hourly" || *result.DefaultFrequency == "daily" || *result.DefaultFrequency == "weekly" || *result.DefaultFrequency == "monthly") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.default_frequency", *result.DefaultFrequency, []interface{}{"15 minutes", "hourly", "daily", "weekly", "monthly"}))
		}
	}
	if result.Href != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.href", *result.Href, "^/api/governance/orgs/[0-9]+/published_templates/[0-9a-f]+$"))
	}
	if result.CreatedBy != nil {
		if err2 := ValidateUserView(result.CreatedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.UpdatedBy != nil {
		if err2 := ValidateUserView(result.UpdatedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	if result.HiddenBy != nil {
		if err2 := ValidateUserView(result.HiddenBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.HiddenAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.hidden_at", *result.HiddenAt, goa.FormatDateTime))
	}
	if result.Severity != nil {
		if !(*result.Severity == "low" || *result.Severity == "medium" || *result.Severity == "high" || *result.Severity == "critical") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.severity", *result.Severity, []interface{}{"low", "medium", "high", "critical"}))
		}
	}
	if result.Tenancy != nil {
		if !(*result.Tenancy == "multi" || *result.Tenancy == "single") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.tenancy", *result.Tenancy, []interface{}{"multi", "single"}))
		}
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#published_template") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#published_template"}))
		}
	}
	return
}

// ValidatePublishedTemplateViewExtended runs the validations defined on
// PublishedTemplateView using the "extended" view.
func ValidatePublishedTemplateViewExtended(result *PublishedTemplateView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	if result.Fingerprint == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fingerprint", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.PolicyTemplateURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.policy_template_url", *result.PolicyTemplateURL, goa.FormatURI))
	}
	if result.DocLink != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.doc_link", *result.DocLink, goa.FormatURI))
	}
	if result.DefaultFrequency != nil {
		if !(*result.DefaultFrequency == "15 minutes" || *result.DefaultFrequency == "hourly" || *result.DefaultFrequency == "daily" || *result.DefaultFrequency == "weekly" || *result.DefaultFrequency == "monthly") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.default_frequency", *result.DefaultFrequency, []interface{}{"15 minutes", "hourly", "daily", "weekly", "monthly"}))
		}
	}
	if result.Href != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.href", *result.Href, "^/api/governance/orgs/[0-9]+/published_templates/[0-9a-f]+$"))
	}
	if result.CreatedBy != nil {
		if err2 := ValidateUserView(result.CreatedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.UpdatedBy != nil {
		if err2 := ValidateUserView(result.UpdatedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	if result.HiddenBy != nil {
		if err2 := ValidateUserView(result.HiddenBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.HiddenAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.hidden_at", *result.HiddenAt, goa.FormatDateTime))
	}
	for _, v := range result.Permissions {
		if v != nil {
			if err2 := ValidatePermissionView(v); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, v := range result.Parameters {
		if v != nil {
			if err2 := ValidateParameterView(v); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Severity != nil {
		if !(*result.Severity == "low" || *result.Severity == "medium" || *result.Severity == "high" || *result.Severity == "critical") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.severity", *result.Severity, []interface{}{"low", "medium", "high", "critical"}))
		}
	}
	if result.Tenancy != nil {
		if !(*result.Tenancy == "multi" || *result.Tenancy == "single") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.tenancy", *result.Tenancy, []interface{}{"multi", "single"}))
		}
	}
	for _, v := range result.Credentials {
		if v != nil {
			if err2 := ValidateCredentialsView(v); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#published_template") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#published_template"}))
		}
	}
	return
}

// ValidatePublishedTemplateViewSource runs the validations defined on
// PublishedTemplateView using the "source" view.
func ValidatePublishedTemplateViewSource(result *PublishedTemplateView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	if result.Fingerprint == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fingerprint", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.Href != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.href", *result.Href, "^/api/governance/orgs/[0-9]+/published_templates/[0-9a-f]+$"))
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#published_template") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#published_template"}))
		}
	}
	return
}

// ValidatePublishedTemplateViewLink runs the validations defined on
// PublishedTemplateView using the "link" view.
func ValidatePublishedTemplateViewLink(result *PublishedTemplateView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	if result.Fingerprint == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fingerprint", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.Href != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.href", *result.Href, "^/api/governance/orgs/[0-9]+/published_templates/[0-9a-f]+$"))
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	if result.UpdatedBy != nil {
		if err2 := ValidateUserView(result.UpdatedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#published_template") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#published_template"}))
		}
	}
	return
}

// ValidateUserView runs the validations defined on UserView.
func ValidateUserView(result *UserView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.email", *result.Email, goa.FormatEmail))
	}
	return
}

// ValidatePermissionView runs the validations defined on PermissionView.
func ValidatePermissionView(result *PermissionView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Resources == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("resources", "result"))
	}
	if result.Actions == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("actions", "result"))
	}
	return
}

// ValidateParameterView runs the validations defined on ParameterView.
func ValidateParameterView(result *ParameterView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Label == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("label", "result"))
	}
	if result.Index == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("index", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "string" || *result.Type == "list" || *result.Type == "number") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []interface{}{"string", "list", "number"}))
		}
	}
	if result.AllowedPattern != nil {
		if err2 := ValidateRegexpView(result.AllowedPattern); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateRegexpView runs the validations defined on RegexpView.
func ValidateRegexpView(result *RegexpView) (err error) {
	if result.Pattern == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pattern", "result"))
	}
	return
}

// ValidateCredentialsView runs the validations defined on CredentialsView.
func ValidateCredentialsView(result *CredentialsView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Schemes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("schemes", "result"))
	}
	if result.Label == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("label", "result"))
	}
	for _, e := range result.Tags {
		if e != nil {
			if err2 := ValidateCredentialsTagView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateCredentialsTagView runs the validations defined on
// CredentialsTagView.
func ValidateCredentialsTagView(result *CredentialsTagView) (err error) {
	if result.Key == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("key", "result"))
	}
	if result.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "result"))
	}
	return
}

// ValidatePublishedTemplateListView runs the validations defined on
// PublishedTemplateListView using the "default" view.
func ValidatePublishedTemplateListView(result *PublishedTemplateListView) (err error) {
	if result.Etag == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("etag", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.NotModified != nil {
		if !(*result.NotModified == "true" || *result.NotModified == "false") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.not_modified", *result.NotModified, []interface{}{"true", "false"}))
		}
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#published_template_list") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#published_template_list"}))
		}
	}
	if result.Items != nil {
		if err2 := ValidatePublishedTemplateCollectionView(result.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePublishedTemplateListViewExtended runs the validations defined on
// PublishedTemplateListView using the "extended" view.
func ValidatePublishedTemplateListViewExtended(result *PublishedTemplateListView) (err error) {
	if result.Etag == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("etag", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.NotModified != nil {
		if !(*result.NotModified == "true" || *result.NotModified == "false") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.not_modified", *result.NotModified, []interface{}{"true", "false"}))
		}
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#published_template_list") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#published_template_list"}))
		}
	}
	if result.Items != nil {
		if err2 := ValidatePublishedTemplateCollectionViewExtended(result.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePublishedTemplateCollectionView runs the validations defined on
// PublishedTemplateCollectionView using the "default" view.
func ValidatePublishedTemplateCollectionView(result PublishedTemplateCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidatePublishedTemplateView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePublishedTemplateCollectionViewExtended runs the validations defined
// on PublishedTemplateCollectionView using the "extended" view.
func ValidatePublishedTemplateCollectionViewExtended(result PublishedTemplateCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidatePublishedTemplateViewExtended(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePublishedTemplateCollectionViewSource runs the validations defined
// on PublishedTemplateCollectionView using the "source" view.
func ValidatePublishedTemplateCollectionViewSource(result PublishedTemplateCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidatePublishedTemplateViewSource(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePublishedTemplateCollectionViewLink runs the validations defined on
// PublishedTemplateCollectionView using the "link" view.
func ValidatePublishedTemplateCollectionViewLink(result PublishedTemplateCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidatePublishedTemplateViewLink(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
