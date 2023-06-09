// Code generated by ogen, DO NOT EDIT.

package senzingchatapi

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// EntityDetailsEntityDetailsGet implements entity_details_entity_details_get operation.
//
// Retrieve entity data based on the ID of a resolved identity.
//
// GET /entity_details
func (UnimplementedHandler) EntityDetailsEntityDetailsGet(ctx context.Context, params EntityDetailsEntityDetailsGetParams) (r EntityDetailsEntityDetailsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// EntityHowEntityHowGet implements entity_how_entity_how_get operation.
//
// Determines and details steps-by-step how records resolved to an ENTITY_ID.
//
// GET /entity_how
func (UnimplementedHandler) EntityHowEntityHowGet(ctx context.Context, params EntityHowEntityHowGetParams) (r EntityHowEntityHowGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// EntityReportEntityReportGet implements entity_report_entity_report_get operation.
//
// Return 10 entities with either matches, possible matches, or relationships.
//
// GET /entity_report
func (UnimplementedHandler) EntityReportEntityReportGet(ctx context.Context, params EntityReportEntityReportGetParams) (r EntityReportEntityReportGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// EntitySearchEntitySearchPost implements entity_search_entity_search_post operation.
//
// Retrieves entity data based on a user-specified set of entity attributes.
//
// POST /entity_search
func (UnimplementedHandler) EntitySearchEntitySearchPost(ctx context.Context, req *SearchAttributes) (r EntitySearchEntitySearchPostRes, _ error) {
	return r, ht.ErrNotImplemented
}
