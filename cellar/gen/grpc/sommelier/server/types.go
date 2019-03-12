// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// sommelier gRPC server types
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package server

import (
	"unicode/utf8"

	sommelierpb "goa.design/examples/cellar/gen/grpc/sommelier/pb"
	sommelier "goa.design/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
	goa "goa.design/goa"
)

// NewPickPayload builds the payload of the "pick" endpoint of the "sommelier"
// service from the gRPC request type.
func NewPickPayload(message *sommelierpb.PickRequest) *sommelier.Criteria {
	v := &sommelier.Criteria{
		Name:   &message.Name,
		Winery: &message.Winery,
	}
	if message.Varietal != nil {
		v.Varietal = make([]string, len(message.Varietal))
		for i, val := range message.Varietal {
			v.Varietal[i] = val
		}
	}
	return v
}

// NewStoredBottleCollection builds the gRPC response type from the result of
// the "pick" endpoint of the "sommelier" service.
func NewStoredBottleCollection(result sommelierviews.StoredBottleCollectionView) *sommelierpb.StoredBottleCollection {
	message := &sommelierpb.StoredBottleCollection{}
	message.Field = make([]*sommelierpb.StoredBottle, len(result))
	for i, val := range result {
		message.Field[i] = &sommelierpb.StoredBottle{}
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
			message.Field[i].Winery = svcSommelierviewsWineryViewToSommelierpbWinery(val.Winery)
		}
		if val.Composition != nil {
			message.Field[i].Composition = make([]*sommelierpb.Component, len(val.Composition))
			for j, val := range val.Composition {
				message.Field[i].Composition[j] = &sommelierpb.Component{}
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

// ValidateStoredBottleCollection runs the validations defined on
// StoredBottleCollection.
func ValidateStoredBottleCollection(message *sommelierpb.StoredBottleCollection) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateStoredBottle(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStoredBottle runs the validations defined on StoredBottle.
func ValidateStoredBottle(message *sommelierpb.StoredBottle) (err error) {
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

// ValidateWinery runs the validations defined on Winery.
func ValidateWinery(message *sommelierpb.Winery) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.region", message.Region, "(?i)[a-z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("message.country", message.Country, "(?i)[a-z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("message.url", message.Url, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	return
}

// ValidateComponent runs the validations defined on Component.
func ValidateComponent(message *sommelierpb.Component) (err error) {
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

// svcSommelierviewsWineryViewToSommelierpbWinery builds a value of type
// *sommelierpb.Winery from a value of type *sommelierviews.WineryView.
func svcSommelierviewsWineryViewToSommelierpbWinery(v *sommelierviews.WineryView) *sommelierpb.Winery {
	if v == nil {
		return nil
	}
	res := &sommelierpb.Winery{}
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

// protobufSommelierpbWineryToSommelierviewsWineryView builds a value of type
// *sommelierviews.WineryView from a value of type *sommelierpb.Winery.
func protobufSommelierpbWineryToSommelierviewsWineryView(v *sommelierpb.Winery) *sommelierviews.WineryView {
	res := &sommelierviews.WineryView{
		Name:    &v.Name,
		Region:  &v.Region,
		Country: &v.Country,
		URL:     &v.Url,
	}

	return res
}
