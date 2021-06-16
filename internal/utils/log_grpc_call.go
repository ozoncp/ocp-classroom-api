package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// TODO: add new func which takes funcName, req, res, err, kafkaErr and probably promErr

// LogGrpcCall logs gRPC calls
func LogGrpcCall(funcName string, req, res interface{}, err *error) {

	var logEvent *zerolog.Event

	if *err != nil {
		logEvent = log.Error()
	} else {
		logEvent = log.Debug()
	}

	logEvent.
		Err(*err).
		Interface("Request", req).
		Interface("Response", res).
		Msg(funcName)
}
