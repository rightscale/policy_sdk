// Code generated by goa v3.1.3, DO NOT EDIT.
//
// IncidentAggregate client HTTP transport
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design -o .

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the IncidentAggregate service endpoint HTTP clients.
type Client struct {
	// Show Doer is the HTTP client used to make requests to the show endpoint.
	ShowDoer goahttp.Doer

	// ShowNonCatalog Doer is the HTTP client used to make requests to the
	// show_non_catalog endpoint.
	ShowNonCatalogDoer goahttp.Doer

	// Index Doer is the HTTP client used to make requests to the index endpoint.
	IndexDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the IncidentAggregate service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ShowDoer:            doer,
		ShowNonCatalogDoer:  doer,
		IndexDoer:           doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Show returns an endpoint that makes HTTP requests to the IncidentAggregate
// service show server.
func (c *Client) Show() goa.Endpoint {
	var (
		encodeRequest  = EncodeShowRequest(c.encoder)
		decodeResponse = DecodeShowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildShowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ShowDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("IncidentAggregate", "show", err)
		}
		return decodeResponse(resp)
	}
}

// ShowNonCatalog returns an endpoint that makes HTTP requests to the
// IncidentAggregate service show_non_catalog server.
func (c *Client) ShowNonCatalog() goa.Endpoint {
	var (
		encodeRequest  = EncodeShowNonCatalogRequest(c.encoder)
		decodeResponse = DecodeShowNonCatalogResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildShowNonCatalogRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ShowNonCatalogDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("IncidentAggregate", "show_non_catalog", err)
		}
		return decodeResponse(resp)
	}
}

// Index returns an endpoint that makes HTTP requests to the IncidentAggregate
// service index server.
func (c *Client) Index() goa.Endpoint {
	var (
		encodeRequest  = EncodeIndexRequest(c.encoder)
		decodeResponse = DecodeIndexResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildIndexRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.IndexDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("IncidentAggregate", "index", err)
		}
		return decodeResponse(resp)
	}
}
