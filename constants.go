package fetch

import "errors"

var HEAD = "HEAD"
var GET = "GET"
var POST = "POST"
var PUT = "PUT"
var DELETE = "DELETE"
var PATCH = "PATCH"

var METHODS = []string{
	HEAD,
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
var ErrInvalidContentType = errors.New("invalid content type")
var ErrorInvalidBody = errors.New("invalid body")
var ErrInvalidBodyMultipart = errors.New("invalid body multipart")
var ErrCannotCreateFormFile = errors.New("cannot create form file")
var ErrCannotCopyFile = errors.New("cannot copy file")
var ErrInvalidUrlFormEncodedBody = errors.New("invalid url form encoded body")
