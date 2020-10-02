// Code generated by goa v3.2.4, DO NOT EDIT.
//
// book HTTP server encoders and decoders
//
// Command:
// $ goa gen book/design

package server

import (
	book "book/gen/book"
	bookviews "book/gen/book/views"
	"context"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateResponse returns an encoder for responses returned by the book
// create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*bookviews.Book)
		enc := encoder(ctx, w)
		body := NewCreateResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the book create
// endpoint.
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
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateBook(&body)

		return payload, nil
	}
}

// EncodeListResponse returns an encoder for responses returned by the book
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]*book.Book)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalBookBookToBookResponse builds a value of type *BookResponse from a
// value of type *book.Book.
func marshalBookBookToBookResponse(v *book.Book) *BookResponse {
	res := &BookResponse{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		Price:       v.Price,
	}

	return res
}