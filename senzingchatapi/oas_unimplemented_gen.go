// Code generated by ogen, DO NOT EDIT.

package senzingchatapi

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AddPet implements addPet operation.
//
// Add a new pet to the store.
//
// POST /pet
func (UnimplementedHandler) AddPet(ctx context.Context, req *Pet) (r *Pet, _ error) {
	return r, ht.ErrNotImplemented
}

// DeletePet implements deletePet operation.
//
// Deletes a pet.
//
// DELETE /pet/{petId}
func (UnimplementedHandler) DeletePet(ctx context.Context, params DeletePetParams) error {
	return ht.ErrNotImplemented
}

// GetPetById implements getPetById operation.
//
// Returns a single pet.
//
// GET /pet/{petId}
func (UnimplementedHandler) GetPetById(ctx context.Context, params GetPetByIdParams) (r GetPetByIdRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdatePet implements updatePet operation.
//
// Updates a pet in the store.
//
// POST /pet/{petId}
func (UnimplementedHandler) UpdatePet(ctx context.Context, params UpdatePetParams) error {
	return ht.ErrNotImplemented
}
