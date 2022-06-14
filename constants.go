package fetch

import "errors"

// HEAD is request method HEAD
var HEAD = "HEAD"

// GET is request method GET
var GET = "GET"

// POST is request method POST
var POST = "POST"

// PUT is request method PUT
var PUT = "PUT"

// DELETE is request method DELETE
var DELETE = "DELETE"

// PATCH is request method PATCH
var PATCH = "PATCH"

// METHODS is the list of supported methods
var METHODS = []string{
	HEAD,
	GET,
	POST,
	PUT,
	DELETE,
	PATCH,
}

// ErrTooManyArguments is the error when the number of arguments is too many
var ErrTooManyArguments = errors.New("too many arguments")

// ErrInvalidMethod is the error when the method is invalid
var ErrInvalidMethod = errors.New("invalid method")

// ErrCannotCreateRequest is the error when the request cannot be created
var ErrCannotCreateRequest = errors.New("cannot create request")

// ErrCannotSendBodyWithGet is the error when the body cannot be sent with GET method
var ErrCannotSendBodyWithGet = errors.New("cannot send body with GET method")

// ErrInvalidJSONBody is the error when the body is not a valid JSON
var ErrInvalidJSONBody = errors.New("error marshalling body")

// ErrSendingRequest is the error when the request cannot be sent
var ErrSendingRequest = errors.New("error sending request")

// ErrReadingResponse is the error when the response cannot be read
var ErrReadingResponse = errors.New("error reading response")

// ErrInvalidContentType is the error when the content type is invalid
var ErrInvalidContentType = errors.New("invalid content type")

// ErrorInvalidBody is the error when the body is invalid
var ErrorInvalidBody = errors.New("invalid body")

// ErrInvalidBodyMultipart is the error when the body is invalid for multipart
var ErrInvalidBodyMultipart = errors.New("invalid body multipart")

// ErrCannotCreateFormFile is the error when the form file cannot be created
var ErrCannotCreateFormFile = errors.New("cannot create form file")

// ErrCannotCopyFile is the error when the file cannot be copied
var ErrCannotCopyFile = errors.New("cannot copy file")

// ErrInvalidURLFormEncodedBody is the error when the body is invalid for url form encoded
var ErrInvalidURLFormEncodedBody = errors.New("invalid url form encoded body")
