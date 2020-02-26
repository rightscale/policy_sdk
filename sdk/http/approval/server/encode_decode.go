// Code generated by goa v3.0.6, DO NOT EDIT.
//
// Approval HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"

	approval "github.com/rightscale/policy_sdk/sdk/approval"
	approvalviews "github.com/rightscale/policy_sdk/sdk/approval/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeShowResponse returns an encoder for responses returned by the Approval
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*approvalviews.ApprovalRequest)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewShowResponseBody(res.Projected)
		case "extended":
			body = NewShowResponseBodyExtended(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the Approval show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID         uint
			approvalRequestID string
			view              *string
			apiVersion        string
			token             *string
			err               error

			params = mux.Vars(r)
		)
		{
			projectIDRaw := params["project_id"]
			v, err2 := strconv.ParseUint(projectIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("projectID", projectIDRaw, "unsigned integer"))
			}
			projectID = uint(v)
		}
		approvalRequestID = params["approval_request_id"]
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "extended") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "extended"}))
			}
		}
		apiVersion = r.Header.Get("Api-Version")
		if apiVersion == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Api-Version", "header"))
		}
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowPayload(projectID, approvalRequestID, view, apiVersion, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show Approval
// endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowNotFoundResponseBody(res)
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeIndexResponse returns an encoder for responses returned by the
// Approval index endpoint.
func EncodeIndexResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*approvalviews.ApprovalRequestList)
		w.Header().Set("goa-view", res.View)
		if res.Projected.NotModified != nil && *res.Projected.NotModified == "true" {
			w.Header().Set("Etag", *res.Projected.Etag)
			w.WriteHeader(http.StatusNotModified)
			return nil
		}
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewIndexOKResponseBody(res.Projected)
		case "extended":
			body = NewIndexOKResponseBodyExtended(res.Projected)
		}
		if res.Projected.Etag != nil {
			w.Header().Set("Etag", *res.Projected.Etag)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeIndexRequest returns a decoder for requests sent to the Approval index
// endpoint.
func DecodeIndexRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID   uint
			view        *string
			id          []string
			subjectKind *string
			subjectHref *string
			status      []string
			apiVersion  string
			etag        *string
			token       *string
			err         error

			params = mux.Vars(r)
		)
		{
			projectIDRaw := params["project_id"]
			v, err2 := strconv.ParseUint(projectIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("projectID", projectIDRaw, "unsigned integer"))
			}
			projectID = uint(v)
		}
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "extended") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "extended"}))
			}
		}
		id = r.URL.Query()["id"]
		subjectKindRaw := r.URL.Query().Get("subject_kind")
		if subjectKindRaw != "" {
			subjectKind = &subjectKindRaw
		}
		subjectHrefRaw := r.URL.Query().Get("subject_href")
		if subjectHrefRaw != "" {
			subjectHref = &subjectHrefRaw
		}
		status = r.URL.Query()["status"]
		for _, e := range status {
			if !(e == "pending" || e == "approved" || e == "denied") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("status[*]", e, []interface{}{"pending", "approved", "denied"}))
			}
		}
		apiVersion = r.Header.Get("Api-Version")
		if apiVersion == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Api-Version", "header"))
		}
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		etagRaw := r.Header.Get("If-None-Match")
		if etagRaw != "" {
			etag = &etagRaw
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewIndexPayload(projectID, view, id, subjectKind, subjectHref, status, apiVersion, etag, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeIndexError returns an encoder for errors returned by the index
// Approval endpoint.
func EncodeIndexError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeApproveResponse returns an encoder for responses returned by the
// Approval approve endpoint.
func EncodeApproveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeApproveRequest returns a decoder for requests sent to the Approval
// approve endpoint.
func DecodeApproveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body ApproveRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateApproveRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			projectID         uint
			approvalRequestID string
			apiVersion        string
			token             *string

			params = mux.Vars(r)
		)
		{
			projectIDRaw := params["project_id"]
			v, err2 := strconv.ParseUint(projectIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("projectID", projectIDRaw, "unsigned integer"))
			}
			projectID = uint(v)
		}
		approvalRequestID = params["approval_request_id"]
		apiVersion = r.Header.Get("Api-Version")
		if apiVersion == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Api-Version", "header"))
		}
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewApprovePayload(&body, projectID, approvalRequestID, apiVersion, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeApproveError returns an encoder for errors returned by the approve
// Approval endpoint.
func EncodeApproveError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewApproveNotFoundResponseBody(res)
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewApproveUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewApproveForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewApproveBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewApproveBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewApproveInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDenyResponse returns an encoder for responses returned by the Approval
// deny endpoint.
func EncodeDenyResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeDenyRequest returns a decoder for requests sent to the Approval deny
// endpoint.
func DecodeDenyRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body DenyRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			projectID         uint
			approvalRequestID string
			apiVersion        string
			token             *string

			params = mux.Vars(r)
		)
		{
			projectIDRaw := params["project_id"]
			v, err2 := strconv.ParseUint(projectIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("projectID", projectIDRaw, "unsigned integer"))
			}
			projectID = uint(v)
		}
		approvalRequestID = params["approval_request_id"]
		apiVersion = r.Header.Get("Api-Version")
		if apiVersion == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Api-Version", "header"))
		}
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewDenyPayload(&body, projectID, approvalRequestID, apiVersion, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDenyError returns an encoder for errors returned by the deny Approval
// endpoint.
func EncodeDenyError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDenyNotFoundResponseBody(res)
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDenyUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDenyForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDenyBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDenyBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDenyInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalApprovalviewsApprovalSubjectToApprovalSubject builds a value of type
// *ApprovalSubject from a value of type *approvalviews.ApprovalSubject.
func marshalApprovalviewsApprovalSubjectToApprovalSubject(v *approvalviews.ApprovalSubject) *ApprovalSubject {
	res := &ApprovalSubject{
		Kind: *v.Kind,
		Href: *v.Href,
	}

	return res
}

// marshalApprovalviewsParameterViewToParameterResponseBody builds a value of
// type *ParameterResponseBody from a value of type
// *approvalviews.ParameterView.
func marshalApprovalviewsParameterViewToParameterResponseBody(v *approvalviews.ParameterView) *ParameterResponseBody {
	if v == nil {
		return nil
	}
	res := &ParameterResponseBody{
		Name:                  *v.Name,
		Type:                  *v.Type,
		Label:                 *v.Label,
		Index:                 *v.Index,
		Category:              v.Category,
		Description:           v.Description,
		Default:               v.Default,
		MinLength:             v.MinLength,
		MaxLength:             v.MaxLength,
		MinValue:              v.MinValue,
		MaxValue:              v.MaxValue,
		ConstraintDescription: v.ConstraintDescription,
	}
	if v.NoEcho != nil {
		res.NoEcho = *v.NoEcho
	}
	if v.NoEcho == nil {
		res.NoEcho = false
	}
	if v.AllowedValues != nil {
		res.AllowedValues = make([]interface{}, len(v.AllowedValues))
		for i, val := range v.AllowedValues {
			res.AllowedValues[i] = val
		}
	}
	if v.AllowedPattern != nil {
		res.AllowedPattern = marshalApprovalviewsRegexpViewToRegexpResponseBody(v.AllowedPattern)
	}

	return res
}

// marshalApprovalviewsRegexpViewToRegexpResponseBody builds a value of type
// *RegexpResponseBody from a value of type *approvalviews.RegexpView.
func marshalApprovalviewsRegexpViewToRegexpResponseBody(v *approvalviews.RegexpView) *RegexpResponseBody {
	if v == nil {
		return nil
	}
	res := &RegexpResponseBody{
		Pattern: *v.Pattern,
		Options: v.Options,
	}

	return res
}

// marshalApprovalviewsConfigurationOptionViewToConfigurationOptionResponseBody
// builds a value of type *ConfigurationOptionResponseBody from a value of type
// *approvalviews.ConfigurationOptionView.
func marshalApprovalviewsConfigurationOptionViewToConfigurationOptionResponseBody(v *approvalviews.ConfigurationOptionView) *ConfigurationOptionResponseBody {
	if v == nil {
		return nil
	}
	res := &ConfigurationOptionResponseBody{
		Name:  *v.Name,
		Label: *v.Label,
		Type:  *v.Type,
		Value: v.Value,
	}

	return res
}

// marshalApprovalviewsUserViewToUserResponseBody builds a value of type
// *UserResponseBody from a value of type *approvalviews.UserView.
func marshalApprovalviewsUserViewToUserResponseBody(v *approvalviews.UserView) *UserResponseBody {
	if v == nil {
		return nil
	}
	res := &UserResponseBody{
		ID:    *v.ID,
		Email: *v.Email,
		Name:  *v.Name,
	}

	return res
}

// marshalApprovalviewsApprovalRequestViewToApprovalRequestResponseBody builds
// a value of type *ApprovalRequestResponseBody from a value of type
// *approvalviews.ApprovalRequestView.
func marshalApprovalviewsApprovalRequestViewToApprovalRequestResponseBody(v *approvalviews.ApprovalRequestView) *ApprovalRequestResponseBody {
	if v == nil {
		return nil
	}
	res := &ApprovalRequestResponseBody{
		ID:          *v.ID,
		ProjectID:   *v.ProjectID,
		Href:        *v.Href,
		Label:       v.Label,
		Description: v.Description,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		Status:      v.Status,
		Kind:        *v.Kind,
	}
	if v.Subject != nil {
		res.Subject = marshalApprovalviewsApprovalSubjectToApprovalSubject(v.Subject)
	}

	return res
}

// marshalApprovalviewsApprovalRequestViewToApprovalRequestResponseBodyExtended
// builds a value of type *ApprovalRequestResponseBodyExtended from a value of
// type *approvalviews.ApprovalRequestView.
func marshalApprovalviewsApprovalRequestViewToApprovalRequestResponseBodyExtended(v *approvalviews.ApprovalRequestView) *ApprovalRequestResponseBodyExtended {
	if v == nil {
		return nil
	}
	res := &ApprovalRequestResponseBodyExtended{
		ID:            *v.ID,
		ProjectID:     *v.ProjectID,
		Href:          *v.Href,
		Label:         v.Label,
		Description:   v.Description,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
		Status:        v.Status,
		ApprovedAt:    v.ApprovedAt,
		DenialComment: v.DenialComment,
		DeniedAt:      v.DeniedAt,
		Kind:          *v.Kind,
	}
	if v.Subject != nil {
		res.Subject = marshalApprovalviewsApprovalSubjectToApprovalSubject(v.Subject)
	}
	if v.Parameters != nil {
		res.Parameters = make(map[string]*ParameterResponseBody, len(v.Parameters))
		for key, val := range v.Parameters {
			tk := key
			res.Parameters[tk] = marshalApprovalviewsParameterViewToParameterResponseBody(val)
		}
	}
	if v.Options != nil {
		res.Options = make([]*ConfigurationOptionResponseBody, len(v.Options))
		for i, val := range v.Options {
			res.Options[i] = marshalApprovalviewsConfigurationOptionViewToConfigurationOptionResponseBody(val)
		}
	}
	if v.ApprovedBy != nil {
		res.ApprovedBy = marshalApprovalviewsUserViewToUserResponseBody(v.ApprovedBy)
	}
	if v.DeniedBy != nil {
		res.DeniedBy = marshalApprovalviewsUserViewToUserResponseBody(v.DeniedBy)
	}

	return res
}

// unmarshalConfigurationOptionCreateTypeRequestBodyToApprovalConfigurationOptionCreateType
// builds a value of type *approval.ConfigurationOptionCreateType from a value
// of type *ConfigurationOptionCreateTypeRequestBody.
func unmarshalConfigurationOptionCreateTypeRequestBodyToApprovalConfigurationOptionCreateType(v *ConfigurationOptionCreateTypeRequestBody) *approval.ConfigurationOptionCreateType {
	if v == nil {
		return nil
	}
	res := &approval.ConfigurationOptionCreateType{
		Name:  *v.Name,
		Value: v.Value,
	}

	return res
}
