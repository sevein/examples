// Code generated by goa v3.7.6, DO NOT EDIT.
//
// storage gRPC server types
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package server

import (
	"unicode/utf8"

	storagepb "goa.design/examples/cellar/gen/grpc/storage/pb"
	storage "goa.design/examples/cellar/gen/storage"
	storageviews "goa.design/examples/cellar/gen/storage/views"
	goa "goa.design/goa/v3/pkg"
)

// NewProtoStoredBottleCollection builds the gRPC response type from the result
// of the "list" endpoint of the "storage" service.
func NewProtoStoredBottleCollection(result storageviews.StoredBottleCollectionView) *storagepb.StoredBottleCollection {
	message := &storagepb.StoredBottleCollection{}
	message.Field = make([]*storagepb.StoredBottle, len(result))
	for i, val := range result {
		message.Field[i] = &storagepb.StoredBottle{}
		if val.ID != nil {
			message.Field[i].Id = *val.ID
		}
		if val.Name != nil {
			message.Field[i].Name = *val.Name
		}
		if val.Vintage != nil {
			message.Field[i].Vintage = *val.Vintage
		}
		if val.Description != nil {
			message.Field[i].Description = *val.Description
		}
		if val.Rating != nil {
			message.Field[i].Rating = *val.Rating
		}
		if val.Winery != nil {
			message.Field[i].Winery = svcStorageviewsWineryViewToStoragepbWinery(val.Winery)
		}
		if val.Composition != nil {
			message.Field[i].Composition = make([]*storagepb.Component, len(val.Composition))
			for j, val := range val.Composition {
				message.Field[i].Composition[j] = &storagepb.Component{}
				if val.Varietal != nil {
					message.Field[i].Composition[j].Varietal = *val.Varietal
				}
				if val.Percentage != nil {
					message.Field[i].Composition[j].Percentage = *val.Percentage
				}
			}
		}
	}
	return message
}

// NewShowPayload builds the payload of the "show" endpoint of the "storage"
// service from the gRPC request type.
func NewShowPayload(message *storagepb.ShowRequest, view *string) *storage.ShowPayload {
	v := &storage.ShowPayload{
		ID: message.Id,
	}
	v.View = view
	return v
}

// NewProtoShowResponse builds the gRPC response type from the result of the
// "show" endpoint of the "storage" service.
func NewProtoShowResponse(result *storageviews.StoredBottleView) *storagepb.ShowResponse {
	message := &storagepb.ShowResponse{}
	if result.ID != nil {
		message.Id = *result.ID
	}
	if result.Name != nil {
		message.Name = *result.Name
	}
	if result.Vintage != nil {
		message.Vintage = *result.Vintage
	}
	if result.Description != nil {
		message.Description = *result.Description
	}
	if result.Rating != nil {
		message.Rating = *result.Rating
	}
	if result.Winery != nil {
		message.Winery = svcStorageviewsWineryViewToStoragepbWinery(result.Winery)
	}
	if result.Composition != nil {
		message.Composition = make([]*storagepb.Component, len(result.Composition))
		for i, val := range result.Composition {
			message.Composition[i] = &storagepb.Component{}
			if val.Varietal != nil {
				message.Composition[i].Varietal = *val.Varietal
			}
			if val.Percentage != nil {
				message.Composition[i].Percentage = *val.Percentage
			}
		}
	}
	return message
}

// NewShowNotFoundError builds the gRPC error response type from the error of
// the "show" endpoint of the "storage" service.
func NewShowNotFoundError(er *storage.NotFound) *storagepb.ShowNotFoundError {
	message := &storagepb.ShowNotFoundError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// NewAddPayload builds the payload of the "add" endpoint of the "storage"
// service from the gRPC request type.
func NewAddPayload(message *storagepb.AddRequest) *storage.Bottle {
	v := &storage.Bottle{
		Name:    message.Name,
		Vintage: message.Vintage,
	}
	if message.Description != "" {
		v.Description = &message.Description
	}
	if message.Rating != 0 {
		v.Rating = &message.Rating
	}
	if message.Winery != nil {
		v.Winery = protobufStoragepbWineryToStorageWinery(message.Winery)
	}
	if message.Composition != nil {
		v.Composition = make([]*storage.Component, len(message.Composition))
		for i, val := range message.Composition {
			v.Composition[i] = &storage.Component{
				Varietal: val.Varietal,
			}
			if val.Percentage != 0 {
				v.Composition[i].Percentage = &val.Percentage
			}
		}
	}
	return v
}

// NewProtoAddResponse builds the gRPC response type from the result of the
// "add" endpoint of the "storage" service.
func NewProtoAddResponse(result string) *storagepb.AddResponse {
	message := &storagepb.AddResponse{}
	message.Field = result
	return message
}

// NewRemovePayload builds the payload of the "remove" endpoint of the
// "storage" service from the gRPC request type.
func NewRemovePayload(message *storagepb.RemoveRequest) *storage.RemovePayload {
	v := &storage.RemovePayload{
		ID: message.Id,
	}
	return v
}

// NewProtoRemoveResponse builds the gRPC response type from the result of the
// "remove" endpoint of the "storage" service.
func NewProtoRemoveResponse() *storagepb.RemoveResponse {
	message := &storagepb.RemoveResponse{}
	return message
}

// NewRatePayload builds the payload of the "rate" endpoint of the "storage"
// service from the gRPC request type.
func NewRatePayload(message *storagepb.RateRequest) map[uint32][]string {
	v := make(map[uint32][]string, len(message.Field))
	for key, val := range message.Field {
		tk := key
		tv := make([]string, len(val.Field))
		for i, val := range val.Field {
			tv[i] = val
		}
		v[tk] = tv
	}
	return v
}

// NewProtoRateResponse builds the gRPC response type from the result of the
// "rate" endpoint of the "storage" service.
func NewProtoRateResponse() *storagepb.RateResponse {
	message := &storagepb.RateResponse{}
	return message
}

// NewMultiAddPayload builds the payload of the "multi_add" endpoint of the
// "storage" service from the gRPC request type.
func NewMultiAddPayload(message *storagepb.MultiAddRequest) []*storage.Bottle {
	v := make([]*storage.Bottle, len(message.Field))
	for i, val := range message.Field {
		v[i] = &storage.Bottle{
			Name:    val.Name,
			Vintage: val.Vintage,
		}
		if val.Description != "" {
			v[i].Description = &val.Description
		}
		if val.Rating != 0 {
			v[i].Rating = &val.Rating
		}
		if val.Winery != nil {
			v[i].Winery = protobufStoragepbWineryToStorageWinery(val.Winery)
		}
		if val.Composition != nil {
			v[i].Composition = make([]*storage.Component, len(val.Composition))
			for j, val := range val.Composition {
				v[i].Composition[j] = &storage.Component{
					Varietal: val.Varietal,
				}
				if val.Percentage != 0 {
					v[i].Composition[j].Percentage = &val.Percentage
				}
			}
		}
	}
	return v
}

// NewProtoMultiAddResponse builds the gRPC response type from the result of
// the "multi_add" endpoint of the "storage" service.
func NewProtoMultiAddResponse(result []string) *storagepb.MultiAddResponse {
	message := &storagepb.MultiAddResponse{}
	message.Field = make([]string, len(result))
	for i, val := range result {
		message.Field[i] = val
	}
	return message
}

// NewMultiUpdatePayload builds the payload of the "multi_update" endpoint of
// the "storage" service from the gRPC request type.
func NewMultiUpdatePayload(message *storagepb.MultiUpdateRequest) *storage.MultiUpdatePayload {
	v := &storage.MultiUpdatePayload{}
	if message.Ids != nil {
		v.Ids = make([]string, len(message.Ids))
		for i, val := range message.Ids {
			v.Ids[i] = val
		}
	}
	if message.Bottles != nil {
		v.Bottles = make([]*storage.Bottle, len(message.Bottles))
		for i, val := range message.Bottles {
			v.Bottles[i] = &storage.Bottle{
				Name:    val.Name,
				Vintage: val.Vintage,
			}
			if val.Description != "" {
				v.Bottles[i].Description = &val.Description
			}
			if val.Rating != 0 {
				v.Bottles[i].Rating = &val.Rating
			}
			if val.Winery != nil {
				v.Bottles[i].Winery = protobufStoragepbWineryToStorageWinery(val.Winery)
			}
			if val.Composition != nil {
				v.Bottles[i].Composition = make([]*storage.Component, len(val.Composition))
				for j, val := range val.Composition {
					v.Bottles[i].Composition[j] = &storage.Component{
						Varietal: val.Varietal,
					}
					if val.Percentage != 0 {
						v.Bottles[i].Composition[j].Percentage = &val.Percentage
					}
				}
			}
		}
	}
	return v
}

// NewProtoMultiUpdateResponse builds the gRPC response type from the result of
// the "multi_update" endpoint of the "storage" service.
func NewProtoMultiUpdateResponse() *storagepb.MultiUpdateResponse {
	message := &storagepb.MultiUpdateResponse{}
	return message
}

// ValidateWinery runs the validations defined on Winery.
func ValidateWinery(message *storagepb.Winery) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.region", message.Region, "[a-zA-Z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("message.country", message.Country, "[a-zA-Z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("message.url", message.Url, "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	return
}

// ValidateComponent runs the validations defined on Component.
func ValidateComponent(message *storagepb.Component) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.varietal", message.Varietal, "[A-Za-z' ]+"))
	if utf8.RuneCountInString(message.Varietal) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.varietal", message.Varietal, utf8.RuneCountInString(message.Varietal), 100, false))
	}
	if message.Percentage < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.percentage", message.Percentage, 1, true))
	}
	if message.Percentage > 100 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.percentage", message.Percentage, 100, false))
	}
	return
}

// ValidateAddRequest runs the validations defined on AddRequest.
func ValidateAddRequest(message *storagepb.AddRequest) (err error) {
	if message.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "message"))
	}
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	if message.Winery != nil {
		if err2 := ValidateWinery(message.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if message.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 1900, true))
	}
	if message.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 2020, false))
	}
	for _, e := range message.Composition {
		if e != nil {
			if err2 := ValidateComponent(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if message.Description != "" {
		if utf8.RuneCountInString(message.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", message.Description, utf8.RuneCountInString(message.Description), 2000, false))
		}
	}
	if message.Rating != 0 {
		if message.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 1, true))
		}
	}
	if message.Rating != 0 {
		if message.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 5, false))
		}
	}
	return
}

// ValidateMultiAddRequest runs the validations defined on MultiAddRequest.
func ValidateMultiAddRequest(message *storagepb.MultiAddRequest) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateBottle(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateBottle runs the validations defined on Bottle.
func ValidateBottle(message *storagepb.Bottle) (err error) {
	if message.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "message"))
	}
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	if message.Winery != nil {
		if err2 := ValidateWinery(message.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if message.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 1900, true))
	}
	if message.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 2020, false))
	}
	for _, e := range message.Composition {
		if e != nil {
			if err2 := ValidateComponent(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if utf8.RuneCountInString(message.Description) > 2000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", message.Description, utf8.RuneCountInString(message.Description), 2000, false))
	}
	if message.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 1, true))
	}
	if message.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 5, false))
	}
	return
}

// ValidateMultiUpdateRequest runs the validations defined on
// MultiUpdateRequest.
func ValidateMultiUpdateRequest(message *storagepb.MultiUpdateRequest) (err error) {
	if message.Ids == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ids", "message"))
	}
	if message.Bottles == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bottles", "message"))
	}
	for _, e := range message.Bottles {
		if e != nil {
			if err2 := ValidateBottle(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// svcStorageviewsWineryViewToStoragepbWinery builds a value of type
// *storagepb.Winery from a value of type *storageviews.WineryView.
func svcStorageviewsWineryViewToStoragepbWinery(v *storageviews.WineryView) *storagepb.Winery {
	res := &storagepb.Winery{}
	if v.Name != nil {
		res.Name = *v.Name
	}
	if v.Region != nil {
		res.Region = *v.Region
	}
	if v.Country != nil {
		res.Country = *v.Country
	}
	if v.URL != nil {
		res.Url = *v.URL
	}

	return res
}

// protobufStoragepbWineryToStorageviewsWineryView builds a value of type
// *storageviews.WineryView from a value of type *storagepb.Winery.
func protobufStoragepbWineryToStorageviewsWineryView(v *storagepb.Winery) *storageviews.WineryView {
	res := &storageviews.WineryView{
		Name:    &v.Name,
		Region:  &v.Region,
		Country: &v.Country,
	}
	if v.Url != "" {
		res.URL = &v.Url
	}

	return res
}

// protobufStoragepbWineryToStorageWinery builds a value of type
// *storage.Winery from a value of type *storagepb.Winery.
func protobufStoragepbWineryToStorageWinery(v *storagepb.Winery) *storage.Winery {
	res := &storage.Winery{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
	}
	if v.Url != "" {
		res.URL = &v.Url
	}

	return res
}

// svcStorageWineryToStoragepbWinery builds a value of type *storagepb.Winery
// from a value of type *storage.Winery.
func svcStorageWineryToStoragepbWinery(v *storage.Winery) *storagepb.Winery {
	res := &storagepb.Winery{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
	}
	if v.URL != nil {
		res.Url = *v.URL
	}

	return res
}
