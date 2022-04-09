// Code generated by goa v3.7.1, DO NOT EDIT.
//
// tus HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/tus/design -o tus

package server

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	tus "goa.design/examples/tus/gen/tus"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeHeadResponse returns an encoder for responses returned by the tus head
// endpoint.
func EncodeHeadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*tus.HeadResult)
		w.Header().Set("Tus-Resumable", res.TusResumable)
		{
			val := res.UploadOffset
			uploadOffsets := strconv.FormatInt(val, 10)
			w.Header().Set("Upload-Offset", uploadOffsets)
		}
		if res.UploadLength != nil {
			val := res.UploadLength
			uploadLengths := strconv.FormatInt(*val, 10)
			w.Header().Set("Upload-Length", uploadLengths)
		}
		if res.UploadDeferLength != nil {
			val := res.UploadDeferLength
			uploadDeferLengths := strconv.Itoa(*val)
			w.Header().Set("Upload-Defer-Length", uploadDeferLengths)
		}
		if res.UploadMetadata != nil {
			w.Header().Set("Upload-Metadata", *res.UploadMetadata)
		}
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeHeadRequest returns a decoder for requests sent to the tus head
// endpoint.
func DecodeHeadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id           string
			tusResumable string
			err          error

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidatePattern("id", id, "[0-9a-v]{20}"))
		tusResumable = r.Header.Get("Tus-Resumable")
		if tusResumable == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Tus-Resumable", "header"))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("tusResumable", tusResumable, "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"))
		if err != nil {
			return nil, err
		}
		payload := NewHeadPayload(id, tusResumable)

		return payload, nil
	}
}

// EncodeHeadError returns an encoder for errors returned by the head tus
// endpoint.
func EncodeHeadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "NotFound":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return nil
		case "Gone":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusGone)
			return nil
		case "InvalidTusResumable":
			var res *tus.ErrInvalidTUSResumable
			errors.As(v, &res)
			w.Header().Set("Tus-Version", res.TusVersion)
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusPreconditionFailed)
			return nil
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodePatchResponse returns an encoder for responses returned by the tus
// patch endpoint.
func EncodePatchResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*tus.PatchResult)
		w.Header().Set("Tus-Resumable", res.TusResumable)
		{
			val := res.UploadOffset
			uploadOffsets := strconv.FormatInt(val, 10)
			w.Header().Set("Upload-Offset", uploadOffsets)
		}
		if res.UploadExpires != nil {
			w.Header().Set("Upload-Expires", *res.UploadExpires)
		}
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodePatchRequest returns a decoder for requests sent to the tus patch
// endpoint.
func DecodePatchRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id             string
			tusResumable   string
			uploadOffset   int64
			uploadChecksum *string
			err            error

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidatePattern("id", id, "[0-9a-v]{20}"))
		tusResumable = r.Header.Get("Tus-Resumable")
		if tusResumable == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Tus-Resumable", "header"))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("tusResumable", tusResumable, "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"))
		{
			uploadOffsetRaw := r.Header.Get("Upload-Offset")
			if uploadOffsetRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("Upload-Offset", "header"))
			}
			v, err2 := strconv.ParseInt(uploadOffsetRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("uploadOffset", uploadOffsetRaw, "integer"))
			}
			uploadOffset = v
		}
		uploadChecksumRaw := r.Header.Get("Upload-Checksum")
		if uploadChecksumRaw != "" {
			uploadChecksum = &uploadChecksumRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewPatchPayload(id, tusResumable, uploadOffset, uploadChecksum)

		return payload, nil
	}
}

// EncodePatchError returns an encoder for errors returned by the patch tus
// endpoint.
func EncodePatchError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "InvalidContentType":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return nil
		case "InvalidOffset":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusConflict)
			return nil
		case "NotFound":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return nil
		case "Gone":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusGone)
			return nil
		case "InvalidChecksumAlgorithm":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return nil
		case "ChecksumMismatch":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(460)
			return nil
		case "Internal":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return nil
		case "InvalidTusResumable":
			var res *tus.ErrInvalidTUSResumable
			errors.As(v, &res)
			w.Header().Set("Tus-Version", res.TusVersion)
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusPreconditionFailed)
			return nil
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeOptionsResponse returns an encoder for responses returned by the tus
// options endpoint.
func EncodeOptionsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*tus.OptionsResult)
		w.Header().Set("Tus-Resumable", res.TusResumable)
		w.Header().Set("Tus-Version", res.TusVersion)
		w.Header().Set("Tus-Extension", res.TusExtension)
		if res.TusMaxSize != nil {
			val := res.TusMaxSize
			tusMaxSizes := strconv.FormatInt(*val, 10)
			w.Header().Set("Tus-Max-Size", tusMaxSizes)
		}
		w.Header().Set("Tus-Checksum-Algorithm", res.TusChecksumAlgorithm)
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// EncodeOptionsError returns an encoder for errors returned by the options tus
// endpoint.
func EncodeOptionsError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "InvalidTusResumable":
			var res *tus.ErrInvalidTUSResumable
			errors.As(v, &res)
			w.Header().Set("Tus-Version", res.TusVersion)
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusPreconditionFailed)
			return nil
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodePostResponse returns an encoder for responses returned by the tus post
// endpoint.
func EncodePostResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*tus.PostResult)
		w.Header().Set("Location", res.Location)
		w.Header().Set("Tus-Resumable", res.TusResumable)
		{
			val := res.UploadOffset
			uploadOffsets := strconv.FormatInt(val, 10)
			w.Header().Set("Upload-Offset", uploadOffsets)
		}
		if res.UploadExpires != nil {
			w.Header().Set("Upload-Expires", *res.UploadExpires)
		}
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodePostRequest returns a decoder for requests sent to the tus post
// endpoint.
func DecodePostRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			tusResumable      string
			uploadLength      *int64
			uploadDeferLength *int
			uploadChecksum    *string
			uploadMetadata    *string
			tusMaxSize        *int64
			err               error
		)
		tusResumable = r.Header.Get("Tus-Resumable")
		if tusResumable == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Tus-Resumable", "header"))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("tusResumable", tusResumable, "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"))
		{
			uploadLengthRaw := r.Header.Get("Upload-Length")
			if uploadLengthRaw != "" {
				v, err2 := strconv.ParseInt(uploadLengthRaw, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("uploadLength", uploadLengthRaw, "integer"))
				}
				uploadLength = &v
			}
		}
		{
			uploadDeferLengthRaw := r.Header.Get("Upload-Defer-Length")
			if uploadDeferLengthRaw != "" {
				v, err2 := strconv.ParseInt(uploadDeferLengthRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("uploadDeferLength", uploadDeferLengthRaw, "integer"))
				}
				pv := int(v)
				uploadDeferLength = &pv
			}
		}
		if uploadDeferLength != nil {
			if !(*uploadDeferLength == 1) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("uploadDeferLength", *uploadDeferLength, []interface{}{1}))
			}
		}
		uploadChecksumRaw := r.Header.Get("Upload-Checksum")
		if uploadChecksumRaw != "" {
			uploadChecksum = &uploadChecksumRaw
		}
		uploadMetadataRaw := r.Header.Get("Upload-Metadata")
		if uploadMetadataRaw != "" {
			uploadMetadata = &uploadMetadataRaw
		}
		{
			tusMaxSizeRaw := r.Header.Get("Tus-Max-Size")
			if tusMaxSizeRaw != "" {
				v, err2 := strconv.ParseInt(tusMaxSizeRaw, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("tusMaxSize", tusMaxSizeRaw, "integer"))
				}
				tusMaxSize = &v
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewPostPayload(tusResumable, uploadLength, uploadDeferLength, uploadChecksum, uploadMetadata, tusMaxSize)

		return payload, nil
	}
}

// EncodePostError returns an encoder for errors returned by the post tus
// endpoint.
func EncodePostError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "MissingHeader":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return nil
		case "InvalidDeferLength":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return nil
		case "InvalidChecksumAlgorithm":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return nil
		case "MaximumSizeExceeded":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			return nil
		case "ChecksumMismatch":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(460)
			return nil
		case "InvalidTusResumable":
			var res *tus.ErrInvalidTUSResumable
			errors.As(v, &res)
			w.Header().Set("Tus-Version", res.TusVersion)
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusPreconditionFailed)
			return nil
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteResponse returns an encoder for responses returned by the tus
// delete endpoint.
func EncodeDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*tus.DeleteResult)
		w.Header().Set("Tus-Resumable", res.TusResumable)
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteRequest returns a decoder for requests sent to the tus delete
// endpoint.
func DecodeDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id           string
			tusResumable string
			err          error

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidatePattern("id", id, "[0-9a-v]{20}"))
		tusResumable = r.Header.Get("Tus-Resumable")
		if tusResumable == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Tus-Resumable", "header"))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("tusResumable", tusResumable, "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"))
		if err != nil {
			return nil, err
		}
		payload := NewDeletePayload(id, tusResumable)

		return payload, nil
	}
}

// EncodeDeleteError returns an encoder for errors returned by the delete tus
// endpoint.
func EncodeDeleteError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "NotFound":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return nil
		case "Gone":
			var res *goa.ServiceError
			errors.As(v, &res)
			w.Header().Set("Goa-Attribute-Name", res.Name)
			w.Header().Set("Goa-Attribute-Id", res.ID)
			w.Header().Set("Goa-Attribute-Message", res.Message)
			{
				val := res.Temporary
				temporarys := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Temporary", temporarys)
			}
			{
				val := res.Timeout
				timeouts := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Timeout", timeouts)
			}
			{
				val := res.Fault
				faults := strconv.FormatBool(val)
				w.Header().Set("Goa-Attribute-Fault", faults)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusGone)
			return nil
		case "InvalidTusResumable":
			var res *tus.ErrInvalidTUSResumable
			errors.As(v, &res)
			w.Header().Set("Tus-Version", res.TusVersion)
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusPreconditionFailed)
			return nil
		default:
			return encodeError(ctx, w, v)
		}
	}
}
