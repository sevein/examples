// Code generated by goa v3.13.2, DO NOT EDIT.
//
// sommelier HTTP server types
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package server

import (
	sommelier "goa.design/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
)

// PickRequestBody is the type of the "sommelier" service "pick" endpoint HTTP
// request body.
type PickRequestBody struct {
	// Name of bottle to pick
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Varietals in preference order
	Varietal []string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Winery of bottle to pick
	Winery *string `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
}

// StoredBottleResponseCollection is the type of the "sommelier" service "pick"
// endpoint HTTP response body.
type StoredBottleResponseCollection []*StoredBottleResponse

// StoredBottleResponse is used to define fields on response body types.
type StoredBottleResponse struct {
	// ID is the unique id of the bottle.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Winery that produces wine
	Winery *WineryResponseTiny `form:"winery" json:"winery" xml:"winery"`
	// Vintage of bottle
	Vintage uint32 `form:"vintage" json:"vintage" xml:"vintage"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentResponse `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// WineryResponseTiny is used to define fields on response body types.
type WineryResponseTiny struct {
	// Name of winery
	Name string `form:"name" json:"name" xml:"name"`
}

// ComponentResponse is used to define fields on response body types.
type ComponentResponse struct {
	// Grape varietal
	Varietal string `form:"varietal" json:"varietal" xml:"varietal"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// NewStoredBottleResponseCollection builds the HTTP response body from the
// result of the "pick" endpoint of the "sommelier" service.
func NewStoredBottleResponseCollection(res sommelierviews.StoredBottleCollectionView) StoredBottleResponseCollection {
	body := make([]*StoredBottleResponse, len(res))
	for i, val := range res {
		body[i] = marshalSommelierviewsStoredBottleViewToStoredBottleResponse(val)
	}
	return body
}

// NewPickCriteria builds a sommelier service pick endpoint payload.
func NewPickCriteria(body *PickRequestBody) *sommelier.Criteria {
	v := &sommelier.Criteria{
		Name:   body.Name,
		Winery: body.Winery,
	}
	if body.Varietal != nil {
		v.Varietal = make([]string, len(body.Varietal))
		for i, val := range body.Varietal {
			v.Varietal[i] = val
		}
	}

	return v
}
