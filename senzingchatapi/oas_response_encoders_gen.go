// Code generated by ogen, DO NOT EDIT.

package senzingchatapi

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeEntityDetailsEntityDetailsGetResponse(response EntityDetailsEntityDetailsGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *EntityDetailsEntityDetailsGetOK:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *HTTPValidationError:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		span.SetStatus(codes.Error, http.StatusText(422))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeEntityHowEntityHowGetResponse(response EntityHowEntityHowGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *EntityHowEntityHowGetOK:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *HTTPValidationError:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		span.SetStatus(codes.Error, http.StatusText(422))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeEntityReportEntityReportGetResponse(response EntityReportEntityReportGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *EntityReportEntityReportGetOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *HTTPValidationError:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		span.SetStatus(codes.Error, http.StatusText(422))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeEntitySearchEntitySearchPostResponse(response EntitySearchEntitySearchPostRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *EntitySearchEntitySearchPostOK:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *HTTPValidationError:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		span.SetStatus(codes.Error, http.StatusText(422))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
