package fetch

import "errors"

var GET = "GET"
var POST = "POST"
var PUT = "PUT"
var DELETE = "DELETE"
var PATCH = "PATCH"

var METHODS = []string{
	GET,
	POST,
	PUT,
	DELETE,
	PATCH,
}

var ErrTooManyArguments = errors.New("too many arguments")
var ErrInvalidMethod = errors.New("invalid method")
var ErrCannotCreateRequest = errors.New("cannot create request")
var ErrCannotSendBodyWithGet = errors.New("cannot send body with GET method")
var ErrInvalidJSONBody = errors.New("error marshalling body")
var ErrSendingRequest = errors.New("error sending request")
var ErrReadingResponse = errors.New("error reading response")
