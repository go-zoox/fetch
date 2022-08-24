package fetch

import "errors"

// HEAD is request method HEAD
const HEAD = "HEAD"

// GET is request method GET
const GET = "GET"

// POST is request method POST
const POST = "POST"

// PUT is request method PUT
const PUT = "PUT"

// DELETE is request method DELETE
const DELETE = "DELETE"

// PATCH is request method PATCH
const PATCH = "PATCH"

// METHODS is the list of supported methods
var METHODS = []string{
	HEAD,
	GET,
	POST,
	PUT,
	DELETE,
	PATCH,
}

// HeaderContentType is the content type header name
const HeaderContentType = "Content-Type"

// HeaderAccept is the accept header name
const HeaderAccept = "Accept"

// HeaderReferrer is the referrer header name
const HeaderReferrer = "Referer"

// HeaderUserAgent ...
const HeaderUserAgent = "User-Agent"

// HeaderAuthorization ...
const HeaderAuthorization = "Authorization"

// HeaderCacheControl ...
const HeaderCacheControl = "Cache-Control"

// HeaderAcceptEncoding ...
const HeaderAcceptEncoding = "Accept-Encoding"

// HeaderAcceptLanguage ...
const HeaderAcceptLanguage = "Accept-Language"

// HeaderCookie ...
const HeaderCookie = "Cookie"

// HeaderLocation ...
const HeaderLocation = "Location"

// HeaderContentLength ...
const HeaderContentLength = "Content-Length"

// HeaderContentEncoding ...
const HeaderContentEncoding = "Content-Encoding"

// HeaderTransferEncoding ...
const HeaderTransferEncoding = "Transfer-Encoding"

// HeaderContentLanguage ...
const HeaderContentLanguage = "Content-Language"

// HeaderSetCookie ...
const HeaderSetCookie = "Set-Cookie"

// HeaderXPoweredBy ...
const HeaderXPoweredBy = "X-Powered-By"

// HeaderXRequestID ...
const HeaderXRequestID = "X-Request-ID"

// HeaderXAcceptRanges ...
const HeaderXAcceptRanges = "X-Accept-Ranges"

// EnvDEBUG is the DEBUG env name
const EnvDEBUG = "GO_ZOOX_FETCH_DEBUG"

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

// ErrCookieEmptyKey is the error when the key is empty
var ErrCookieEmptyKey = errors.New("empty key")
