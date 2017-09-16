package merror

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

var httpErrorMessage = map[int]string{
	http.StatusBadRequest:          "400 bad request",
	http.StatusUnauthorized:        "401 unathorize",
	http.StatusForbidden:           "403 forbidden",
	http.StatusNotFound:            "404 page not found",
	http.StatusMethodNotAllowed:    "405 method not allowed",
	http.StatusRequestTimeout:      "408 request timeout",
	http.StatusInternalServerError: "500 internal server error",
	http.StatusServiceUnavailable:  "503 service unavailable",
}

func HTTPErrMessage(code int) string {
	if errm, ok := httpErrorMessage[code]; ok {
		return errm
	}
	return "internal server error"
}

//appError is an error that can handle by user, so it's important the user get meaningfull error Message
type appError interface {
	error
	//Message return the error message that send to user.
	Message() string
}

type AppErr struct {
	err error
	msg string
}

func (a *AppErr) Error() string {
	return a.err.Error()
}

func (a *AppErr) Message() string {
	return a.msg
}

//NewError golang have two pkg with name errors, use this to makesure we use github/pkg/errors
func NewError(msg string) error {
	return errors.New(msg)
}

//NewAppError ...
func NewAppError(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	err := errors.New(msg)
	return AppError(err, msg)
}

//AppError ...
func AppError(err error, message string) error {
	aerr := &AppErr{
		err: err,
		msg: message,
	}
	return errors.Wrap(aerr, message)
}

//AppErrorf ...
func AppErrorf(err error, format string, args ...interface{}) error {
	aerr := &AppErr{
		err: err,
		msg: fmt.Sprintf(format, args...),
	}
	return errors.Wrapf(aerr, format, args...)
}

//IsAppError ...
func IsAppError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := errors.Cause(err).(appError)
	return ok
}

func AppErrorGetMessage(err error) string {
	apperr, ok := errors.Cause(err).(appError)
	if !ok {
		return ""
	}
	return apperr.Message()
}

type httpError interface {
	error
	StatusCode() int
}

type httpErr struct {
	err  error
	code int
}

//HTTPError ....
func HTTPError(err error, statusCode int) error {
	herr := &httpErr{
		err:  err,
		code: statusCode,
	}
	msg := HTTPErrMessage(statusCode)
	return errors.Wrap(herr, msg)
}

func (h *httpErr) Error() string {
	return h.err.Error()
}

func (h *httpErr) StatusCode() int {
	return h.code
}

//IsHTTPError check is the cause of the erorr is implement httpError.
//IsHTTPError return false if the erorr is nil.
func IsHTTPError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := errors.Cause(err).(httpError)
	return ok
}

//IsSQLNoRows check if error is ql no rows
func IsSQLNoRows(err error) bool {
	if err == nil {
		return false
	}
	if err == sql.ErrNoRows {
		return true
	}
	cause := errors.Cause(err)
	return cause == sql.ErrNoRows
}

func PrintErr(err error) {
	fmt.Printf("%+v\n", err)
}
