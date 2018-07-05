package core

import "errors"

var (
	// ErrCannotCreateSession is an error returned from any AWS session initialization that fails.
	ErrCannotCreateSession = errors.New("it was impossible to create an AWS session")

	// ErrCannotCreateSession is an error returned from any AWS session initialization that fails.
	ErrCannotRetrieveLambdas = errors.New("tt was impossible to retrieve AWS Lambda functions")
)
