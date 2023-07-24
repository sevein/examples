// Code generated by goa v3.12.2, DO NOT EDIT.
//
// tus HTTP server types
//
// Command:
// $ goa gen goa.design/examples/tus/design -o tus

package server

import (
	tus "goa.design/examples/tus/gen/tus"
)

// NewHeadPayload builds a tus service head endpoint payload.
func NewHeadPayload(id string, tusResumable string) *tus.HeadPayload {
	v := &tus.HeadPayload{}
	v.ID = id
	v.TusResumable = tusResumable

	return v
}

// NewPatchPayload builds a tus service patch endpoint payload.
func NewPatchPayload(id string, tusResumable string, uploadOffset int64, uploadChecksum *string) *tus.PatchPayload {
	v := &tus.PatchPayload{}
	v.ID = id
	v.TusResumable = tusResumable
	v.UploadOffset = uploadOffset
	v.UploadChecksum = uploadChecksum

	return v
}

// NewPostPayload builds a tus service post endpoint payload.
func NewPostPayload(tusResumable string, uploadLength *int64, uploadDeferLength *int, uploadChecksum *string, uploadMetadata *string, tusMaxSize *int64) *tus.PostPayload {
	v := &tus.PostPayload{}
	v.TusResumable = tusResumable
	v.UploadLength = uploadLength
	v.UploadDeferLength = uploadDeferLength
	v.UploadChecksum = uploadChecksum
	v.UploadMetadata = uploadMetadata
	v.TusMaxSize = tusMaxSize

	return v
}

// NewDeletePayload builds a tus service delete endpoint payload.
func NewDeletePayload(id string, tusResumable string) *tus.DeletePayload {
	v := &tus.DeletePayload{}
	v.ID = id
	v.TusResumable = tusResumable

	return v
}
