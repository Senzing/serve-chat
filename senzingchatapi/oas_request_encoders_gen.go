// Code generated by ogen, DO NOT EDIT.

package senzingchatapi

import (
	"bytes"
	"net/http"

	"github.com/go-faster/jx"

	ht "github.com/ogen-go/ogen/http"
)

func encodeAddPetRequest(
	req *Pet,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}
