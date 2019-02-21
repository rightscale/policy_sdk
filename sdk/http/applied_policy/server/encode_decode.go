// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// AppliedPolicy HTTP server encoders and decoders
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

	appliedpolicyviews "github.com/rightscale/policy_sdk/sdk/applied_policy/views"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// EncodeCreateResponse returns an encoder for responses returned by the
// AppliedPolicy create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*appliedpolicyviews.AppliedPolicy)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewCreateResponseBody(res.Projected)
		case "source":
			body = NewCreateResponseBodySource(res.Projected)
		case "link":
			body = NewCreateResponseBodyLink(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the AppliedPolicy
// create endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = body.Validate()
		if err != nil {
			return nil, err
		}

		var (
			projectID  uint
			apiVersion string
			token      *string

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
		payload := NewCreatePayload(&body, projectID, apiVersion, token)
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

// EncodeCreateError returns an encoder for errors returned by the create
// AppliedPolicy endpoint.
func EncodeCreateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unprocessable_entity":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateUnprocessableEntityResponseBody(res)
			w.Header().Set("goa-error", "unprocessable_entity")
			w.WriteHeader(http.StatusUnprocessableEntity)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewCreateInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteResponse returns an encoder for responses returned by the
// AppliedPolicy delete endpoint.
func EncodeDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteRequest returns a decoder for requests sent to the AppliedPolicy
// delete endpoint.
func DecodeDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID  uint
			policyID   string
			apiVersion string
			token      *string
			err        error

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
		policyID = params["policy_id"]
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
		payload := NewDeletePayload(projectID, policyID, apiVersion, token)
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

// EncodeDeleteError returns an encoder for errors returned by the delete
// AppliedPolicy endpoint.
func EncodeDeleteError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
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
			body := NewDeleteNotFoundResponseBody(res)
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewDeleteInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeShowResponse returns an encoder for responses returned by the
// AppliedPolicy show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*appliedpolicyviews.AppliedPolicy)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewShowResponseBody(res.Projected)
		case "source":
			body = NewShowResponseBodySource(res.Projected)
		case "link":
			body = NewShowResponseBodyLink(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the AppliedPolicy
// show endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID  uint
			policyID   string
			view       *string
			apiVersion string
			token      *string
			err        error

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
		policyID = params["policy_id"]
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "source") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "source"}))
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
		payload := NewShowPayload(projectID, policyID, view, apiVersion, token)
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

// EncodeShowError returns an encoder for errors returned by the show
// AppliedPolicy endpoint.
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

// EncodeShowStatusResponse returns an encoder for responses returned by the
// AppliedPolicy show_status endpoint.
func EncodeShowStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*appliedpolicyviews.AppliedPolicyStatus)
		enc := encoder(ctx, w)
		body := NewShowStatusResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowStatusRequest returns a decoder for requests sent to the
// AppliedPolicy show_status endpoint.
func DecodeShowStatusRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID  uint
			policyID   string
			apiVersion string
			token      *string
			err        error

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
		policyID = params["policy_id"]
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
		payload := NewShowStatusPayload(projectID, policyID, apiVersion, token)
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

// EncodeShowStatusError returns an encoder for errors returned by the
// show_status AppliedPolicy endpoint.
func EncodeShowStatusError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
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
			body := NewShowStatusNotFoundResponseBody(res)
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowStatusUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowStatusForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowStatusBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowStatusBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowStatusInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeShowLogResponse returns an encoder for responses returned by the
// AppliedPolicy show_log endpoint.
func EncodeShowLogResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*appliedpolicyviews.AppliedPolicyLog)
		enc := encoder(ctx, w)
		body := NewShowLogOKResponseBody(res.Projected)
		if res.Projected.Etag != nil {
			w.Header().Set("ETag", *res.Projected.Etag)
		}
		if res.Projected.LastModified != nil {
			w.Header().Set("Last-Modified", *res.Projected.LastModified)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowLogRequest returns a decoder for requests sent to the
// AppliedPolicy show_log endpoint.
func DecodeShowLogRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID  uint
			policyID   string
			apiVersion string
			etag       *string
			token      *string
			err        error

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
		policyID = params["policy_id"]
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
		payload := NewShowLogPayload(projectID, policyID, apiVersion, etag, token)
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

// EncodeShowLogError returns an encoder for errors returned by the show_log
// AppliedPolicy endpoint.
func EncodeShowLogError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
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
			body := NewShowLogNotFoundResponseBody(res)
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowLogUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowLogForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowLogBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowLogBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowLogInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeIndexResponse returns an encoder for responses returned by the
// AppliedPolicy index endpoint.
func EncodeIndexResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*appliedpolicyviews.AppliedPolicyList)
		w.Header().Set("goa-view", res.View)
		if res.Projected.NotModified != nil && *res.Projected.NotModified == "true" {
			w.Header().Set("ETag", *res.Projected.Etag)
			w.WriteHeader(http.StatusNotModified)
			return nil
		}
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewIndexOKResponseBody(res.Projected)
		case "link":
			body = NewIndexOKResponseBodyLink(res.Projected)
		}
		w.Header().Set("ETag", *res.Projected.Etag)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeIndexRequest returns a decoder for requests sent to the AppliedPolicy
// index endpoint.
func DecodeIndexRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID  uint
			view       *string
			name       []string
			apiVersion string
			etag       *string
			token      *string
			err        error

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
			if !(*view == "default") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default"}))
			}
		}
		name = r.URL.Query()["name"]
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
		payload := NewIndexPayload(projectID, view, name, apiVersion, etag, token)
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
// AppliedPolicy endpoint.
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

// EncodeEvaluateResponse returns an encoder for responses returned by the
// AppliedPolicy evaluate endpoint.
func EncodeEvaluateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeEvaluateRequest returns a decoder for requests sent to the
// AppliedPolicy evaluate endpoint.
func DecodeEvaluateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID  uint
			policyID   string
			apiVersion string
			token      *string
			err        error

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
		policyID = params["policy_id"]
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
		payload := NewEvaluatePayload(projectID, policyID, apiVersion, token)
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

// EncodeEvaluateError returns an encoder for errors returned by the evaluate
// AppliedPolicy endpoint.
func EncodeEvaluateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
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
			body := NewEvaluateNotFoundResponseBody(res)
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewEvaluateUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewEvaluateForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewEvaluateBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewEvaluateBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewEvaluateInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalProjectViewToProjectResponseBody builds a value of type
// *ProjectResponseBody from a value of type *appliedpolicyviews.ProjectView.
func marshalProjectViewToProjectResponseBody(v *appliedpolicyviews.ProjectView) *ProjectResponseBody {
	if v == nil {
		return nil
	}
	res := &ProjectResponseBody{
		ID:      *v.ID,
		Name:    *v.Name,
		OrgID:   *v.OrgID,
		OrgName: *v.OrgName,
	}

	return res
}

// marshalPolicyTemplateViewToPolicyTemplateResponseBodyLink builds a value of
// type *PolicyTemplateResponseBodyLink from a value of type
// *appliedpolicyviews.PolicyTemplateView.
func marshalPolicyTemplateViewToPolicyTemplateResponseBodyLink(v *appliedpolicyviews.PolicyTemplateView) *PolicyTemplateResponseBodyLink {
	if v == nil {
		return nil
	}
	res := &PolicyTemplateResponseBodyLink{
		ID:          *v.ID,
		Name:        *v.Name,
		Href:        *v.Href,
		Fingerprint: *v.Fingerprint,
		UpdatedAt:   v.UpdatedAt,
		Kind:        *v.Kind,
	}
	if v.UpdatedBy != nil {
		res.UpdatedBy = marshalUserViewToUserResponseBody(v.UpdatedBy)
	}

	return res
}

// marshalUserViewToUserResponseBody builds a value of type *UserResponseBody
// from a value of type *appliedpolicyviews.UserView.
func marshalUserViewToUserResponseBody(v *appliedpolicyviews.UserView) *UserResponseBody {
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

// marshalPublishedTemplateViewToPublishedTemplateResponseBodyLink builds a
// value of type *PublishedTemplateResponseBodyLink from a value of type
// *appliedpolicyviews.PublishedTemplateView.
func marshalPublishedTemplateViewToPublishedTemplateResponseBodyLink(v *appliedpolicyviews.PublishedTemplateView) *PublishedTemplateResponseBodyLink {
	if v == nil {
		return nil
	}
	res := &PublishedTemplateResponseBodyLink{
		ID:          *v.ID,
		Name:        *v.Name,
		Href:        *v.Href,
		Fingerprint: *v.Fingerprint,
		UpdatedAt:   v.UpdatedAt,
		BuiltIn:     v.BuiltIn,
		Kind:        *v.Kind,
	}
	if v.UpdatedBy != nil {
		res.UpdatedBy = marshalUserViewToUserResponseBody(v.UpdatedBy)
	}

	return res
}

// marshalPolicyTemplateViewToPolicyTemplateResponseBodySource builds a value
// of type *PolicyTemplateResponseBodySource from a value of type
// *appliedpolicyviews.PolicyTemplateView.
func marshalPolicyTemplateViewToPolicyTemplateResponseBodySource(v *appliedpolicyviews.PolicyTemplateView) *PolicyTemplateResponseBodySource {
	if v == nil {
		return nil
	}
	res := &PolicyTemplateResponseBodySource{
		ID:          *v.ID,
		Name:        *v.Name,
		Href:        *v.Href,
		Filename:    v.Filename,
		Source:      v.Source,
		Fingerprint: *v.Fingerprint,
		Kind:        *v.Kind,
	}

	return res
}

// marshalPublishedTemplateViewToPublishedTemplateResponseBodySource builds a
// value of type *PublishedTemplateResponseBodySource from a value of type
// *appliedpolicyviews.PublishedTemplateView.
func marshalPublishedTemplateViewToPublishedTemplateResponseBodySource(v *appliedpolicyviews.PublishedTemplateView) *PublishedTemplateResponseBodySource {
	if v == nil {
		return nil
	}
	res := &PublishedTemplateResponseBodySource{
		ID:          *v.ID,
		Name:        *v.Name,
		Href:        *v.Href,
		Filename:    v.Filename,
		Source:      v.Source,
		Fingerprint: *v.Fingerprint,
		Kind:        *v.Kind,
	}

	return res
}
