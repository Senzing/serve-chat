// Code generated by ogen, DO NOT EDIT.

package senzingchatapi

import (
	"net/http"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// EntityDetailsEntityDetailsGetParams is parameters of entity_details_entity_details_get operation.
type EntityDetailsEntityDetailsGetParams struct {
	EntityID int
}

func unpackEntityDetailsEntityDetailsGetParams(packed middleware.Parameters) (params EntityDetailsEntityDetailsGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "entity_id",
			In:   "query",
		}
		params.EntityID = packed[key].(int)
	}
	return params
}

func decodeEntityDetailsEntityDetailsGetParams(args [0]string, argsEscaped bool, r *http.Request) (params EntityDetailsEntityDetailsGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: entity_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "entity_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.EntityID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "entity_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// EntityHowEntityHowGetParams is parameters of entity_how_entity_how_get operation.
type EntityHowEntityHowGetParams struct {
	EntityID int
}

func unpackEntityHowEntityHowGetParams(packed middleware.Parameters) (params EntityHowEntityHowGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "entity_id",
			In:   "query",
		}
		params.EntityID = packed[key].(int)
	}
	return params
}

func decodeEntityHowEntityHowGetParams(args [0]string, argsEscaped bool, r *http.Request) (params EntityHowEntityHowGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: entity_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "entity_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.EntityID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "entity_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// EntityReportEntityReportGetParams is parameters of entity_report_entity_report_get operation.
type EntityReportEntityReportGetParams struct {
	ExportFlags ExportFlags
}

func unpackEntityReportEntityReportGetParams(packed middleware.Parameters) (params EntityReportEntityReportGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "export_flags",
			In:   "query",
		}
		params.ExportFlags = packed[key].(ExportFlags)
	}
	return params
}

func decodeEntityReportEntityReportGetParams(args [0]string, argsEscaped bool, r *http.Request) (params EntityReportEntityReportGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: export_flags.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "export_flags",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.ExportFlags = ExportFlags(c)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.ExportFlags.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "export_flags",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
