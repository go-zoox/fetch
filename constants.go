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

// // headers.ContentType is the content type header name
// const headers.ContentType = "Content-Type"

// // headers.Accept is the accept header name
// const headers.Accept = "Accept"

// // headers.Referrer is the referrer header name
// const headers.Referrer = "Referer"

// // headers.UserAgent ...
// const headers.UserAgent = "User-Agent"

// // headers.Authorization ...
// const headers.Authorization = "Authorization"

// // headers.CacheControl ...
// const headers.CacheControl = "Cache-Control"

// // headers.AcceptEncoding ...
// const headers.AcceptEncoding = "Accept-Encoding"

// // headers.AcceptLanguage ...
// const headers.AcceptLanguage = "Accept-Language"

// // headers.Cookie ...
// const headers.Cookie = "Cookie"

// // headers.Location ...
// const headers.Location = "Location"

// // headers.ContentLength ...
// const headers.ContentLength = "Content-Length"

// // headers.ContentEncoding ...
// const headers.ContentEncoding = "Content-Encoding"

// // headers.TransferEncoding ...
// const headers.TransferEncoding = "Transfer-Encoding"

// // headers.ContentLanguage ...
// const headers.ContentLanguage = "Content-Language"

// // headers.SetCookie ...
// const headers.SetCookie = "Set-Cookie"

// // headers.XPoweredBy ...
// const headers.XPoweredBy = "X-Powered-By"

// // headers.XRequestID ...
// const headers.XRequestID = "X-Request-ID"

// // headers.AcceptRanges ...
// const headers.AcceptRanges = "Accept-Ranges"

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
