// Code generated by goa v3.7.12, DO NOT EDIT.
//
// updown HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/upload_download/design -o upload_download

package client

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	updown "goa.design/examples/upload_download/gen/updown"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildUploadRequest instantiates a HTTP request object with method and path
// set to call the "updown" service "upload" endpoint
func (c *Client) BuildUploadRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		dir  string
		body io.Reader
	)
	{
		rd, ok := v.(*updown.UploadRequestData)
		if !ok {
			return nil, goahttp.ErrInvalidType("updown", "upload", "updown.UploadRequestData", v)
		}
		p := rd.Payload
		body = rd.Body
		dir = p.Dir
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UploadUpdownPath(dir)}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("updown", "upload", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUploadRequest returns an encoder for requests sent to the updown
// upload server.
func EncodeUploadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		data, ok := v.(*updown.UploadRequestData)
		if !ok {
			return goahttp.ErrInvalidType("updown", "upload", "*updown.UploadRequestData", v)
		}
		p := data.Payload
		{
			head := p.ContentType
			req.Header.Set("Content-Type", head)
		}
		return nil
	}
}

// DecodeUploadResponse returns a decoder for responses returned by the updown
// upload endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeUploadResponse may return the following errors:
//	- "invalid_media_type" (type *goa.ServiceError): http.StatusBadRequest
//	- "invalid_multipart_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeUploadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "invalid_media_type":
				var (
					body UploadInvalidMediaTypeResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("updown", "upload", err)
				}
				err = ValidateUploadInvalidMediaTypeResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("updown", "upload", err)
				}
				return nil, NewUploadInvalidMediaType(&body)
			case "invalid_multipart_request":
				var (
					body UploadInvalidMultipartRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("updown", "upload", err)
				}
				err = ValidateUploadInvalidMultipartRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("updown", "upload", err)
				}
				return nil, NewUploadInvalidMultipartRequest(&body)
			default:
				body, _ := ioutil.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("updown", "upload", resp.StatusCode, string(body))
			}
		case http.StatusInternalServerError:
			var (
				body UploadInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("updown", "upload", err)
			}
			err = ValidateUploadInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("updown", "upload", err)
			}
			return nil, NewUploadInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("updown", "upload", resp.StatusCode, string(body))
		}
	}
}

// // BuildUploadStreamPayload creates a streaming endpoint request payload from
// the method payload and the path to the file to be streamed
func BuildUploadStreamPayload(payload interface{}, fpath string) (*updown.UploadRequestData, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	return &updown.UploadRequestData{
		Payload: payload.(*updown.UploadPayload),
		Body:    f,
	}, nil
}

// BuildDownloadRequest instantiates a HTTP request object with method and path
// set to call the "updown" service "download" endpoint
func (c *Client) BuildDownloadRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		filename string
	)
	{
		p, ok := v.(string)
		if !ok {
			return nil, goahttp.ErrInvalidType("updown", "download", "string", v)
		}
		filename = p
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DownloadUpdownPath(filename)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("updown", "download", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDownloadResponse returns a decoder for responses returned by the
// updown download endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDownloadResponse may return the following errors:
//	- "invalid_file_path" (type *goa.ServiceError): http.StatusNotFound
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeDownloadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				length int64
				err    error
			)
			{
				lengthRaw := resp.Header.Get("Content-Length")
				if lengthRaw == "" {
					return nil, goahttp.ErrValidationError("updown", "download", goa.MissingFieldError("Content-Length", "header"))
				}
				v, err2 := strconv.ParseInt(lengthRaw, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("length", lengthRaw, "integer"))
				}
				length = v
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("updown", "download", err)
			}
			res := NewDownloadResultOK(length)
			return res, nil
		case http.StatusNotFound:
			var (
				body DownloadInvalidFilePathResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("updown", "download", err)
			}
			err = ValidateDownloadInvalidFilePathResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("updown", "download", err)
			}
			return nil, NewDownloadInvalidFilePath(&body)
		case http.StatusInternalServerError:
			var (
				body DownloadInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("updown", "download", err)
			}
			err = ValidateDownloadInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("updown", "download", err)
			}
			return nil, NewDownloadInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("updown", "download", resp.StatusCode, string(body))
		}
	}
}
