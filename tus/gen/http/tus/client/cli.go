// Code generated by goa v3.13.2, DO NOT EDIT.
//
// tus HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/tus/design -o tus

package client

import (
	"fmt"
	"strconv"

	tus "goa.design/examples/tus/gen/tus"
	goa "goa.design/goa/v3/pkg"
)

// BuildHeadPayload builds the payload for the tus head endpoint from CLI flags.
func BuildHeadPayload(tusHeadID string, tusHeadTusResumable string) (*tus.HeadPayload, error) {
	var err error
	var id string
	{
		id = tusHeadID
		err = goa.MergeErrors(err, goa.ValidatePattern("id", id, "[0-9a-v]{20}"))
		if err != nil {
			return nil, err
		}
	}
	var tusResumable string
	{
		tusResumable = tusHeadTusResumable
		err = goa.MergeErrors(err, goa.ValidatePattern("tusResumable", tusResumable, "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"))
		if err != nil {
			return nil, err
		}
	}
	v := &tus.HeadPayload{}
	v.ID = id
	v.TusResumable = tusResumable

	return v, nil
}

// BuildPatchPayload builds the payload for the tus patch endpoint from CLI
// flags.
func BuildPatchPayload(tusPatchID string, tusPatchTusResumable string, tusPatchUploadOffset string, tusPatchUploadChecksum string) (*tus.PatchPayload, error) {
	var err error
	var id string
	{
		id = tusPatchID
		err = goa.MergeErrors(err, goa.ValidatePattern("id", id, "[0-9a-v]{20}"))
		if err != nil {
			return nil, err
		}
	}
	var tusResumable string
	{
		tusResumable = tusPatchTusResumable
		err = goa.MergeErrors(err, goa.ValidatePattern("tusResumable", tusResumable, "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"))
		if err != nil {
			return nil, err
		}
	}
	var uploadOffset int64
	{
		uploadOffset, err = strconv.ParseInt(tusPatchUploadOffset, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for uploadOffset, must be INT64")
		}
	}
	var uploadChecksum *string
	{
		if tusPatchUploadChecksum != "" {
			uploadChecksum = &tusPatchUploadChecksum
		}
	}
	v := &tus.PatchPayload{}
	v.ID = id
	v.TusResumable = tusResumable
	v.UploadOffset = uploadOffset
	v.UploadChecksum = uploadChecksum

	return v, nil
}

// BuildPostPayload builds the payload for the tus post endpoint from CLI flags.
func BuildPostPayload(tusPostTusResumable string, tusPostUploadLength string, tusPostUploadDeferLength string, tusPostUploadChecksum string, tusPostUploadMetadata string, tusPostTusMaxSize string) (*tus.PostPayload, error) {
	var err error
	var tusResumable string
	{
		tusResumable = tusPostTusResumable
		err = goa.MergeErrors(err, goa.ValidatePattern("tusResumable", tusResumable, "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"))
		if err != nil {
			return nil, err
		}
	}
	var uploadLength *int64
	{
		if tusPostUploadLength != "" {
			val, err := strconv.ParseInt(tusPostUploadLength, 10, 64)
			uploadLength = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for uploadLength, must be INT64")
			}
		}
	}
	var uploadDeferLength *int
	{
		if tusPostUploadDeferLength != "" {
			var v int64
			v, err = strconv.ParseInt(tusPostUploadDeferLength, 10, strconv.IntSize)
			val := int(v)
			uploadDeferLength = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for uploadDeferLength, must be INT")
			}
			if !(*uploadDeferLength == 1) {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("uploadDeferLength", *uploadDeferLength, []any{1}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var uploadChecksum *string
	{
		if tusPostUploadChecksum != "" {
			uploadChecksum = &tusPostUploadChecksum
		}
	}
	var uploadMetadata *string
	{
		if tusPostUploadMetadata != "" {
			uploadMetadata = &tusPostUploadMetadata
		}
	}
	var tusMaxSize *int64
	{
		if tusPostTusMaxSize != "" {
			val, err := strconv.ParseInt(tusPostTusMaxSize, 10, 64)
			tusMaxSize = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for tusMaxSize, must be INT64")
			}
		}
	}
	v := &tus.PostPayload{}
	v.TusResumable = tusResumable
	v.UploadLength = uploadLength
	v.UploadDeferLength = uploadDeferLength
	v.UploadChecksum = uploadChecksum
	v.UploadMetadata = uploadMetadata
	v.TusMaxSize = tusMaxSize

	return v, nil
}

// BuildDeletePayload builds the payload for the tus delete endpoint from CLI
// flags.
func BuildDeletePayload(tusDeleteID string, tusDeleteTusResumable string) (*tus.DeletePayload, error) {
	var err error
	var id string
	{
		id = tusDeleteID
		err = goa.MergeErrors(err, goa.ValidatePattern("id", id, "[0-9a-v]{20}"))
		if err != nil {
			return nil, err
		}
	}
	var tusResumable string
	{
		tusResumable = tusDeleteTusResumable
		err = goa.MergeErrors(err, goa.ValidatePattern("tusResumable", tusResumable, "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"))
		if err != nil {
			return nil, err
		}
	}
	v := &tus.DeletePayload{}
	v.ID = id
	v.TusResumable = tusResumable

	return v, nil
}
