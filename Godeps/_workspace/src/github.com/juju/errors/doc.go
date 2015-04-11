// Copyright 2013, 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

/*
The juju/errors provides an easy way to annotate errors without losing the
orginal error context.

The exported `New` and `Errorf` functions are designed to replace the
`errors.New` and `fmt.Errorf` functions respectively. The same underlying
error is there, but the package also records the location at which the error
was created.

A primary use case for this library is to add extra context any time an
error is returned from a function.

    if err := SomeFunc(); err != nil {
	    return err
	}

This instead becomes:

    if err := SomeFunc(); err != nil {
	    return errors.Trace(err)
	}

which just records the file and line number of the Trace call, or

    if err := SomeFunc(); err != nil {
	    return errors.Annotate(err, "more context")
	}

which also adds an annotation to the error.

When you want to check to see if an error is of a particular type, a helper
function is normally exported by the package that returned the error, like the
`os` package does.  The underlying cause of the error is available using the
`Cause` function.

	os.IsNotExist(errors.Cause(err))

The result of the `Error()` call on an annotated error is the annotations joined
with colons, then the result of the `Error()` method for the underlying error
that was the cause.

	err := errors.Errorf("original")
	err = errors.Annotatef(err, "context")
	err = errors.Annotatef(err, "more context")
	err.Error() -> "more context: context: original"

Obviously recording the file, line and functions is not very useful if you
cannot get them back out again.

	errors.ErrorStack(err)

will return something like:

	first error
	github.com/juju/errors/annotation_test.go:193:
	github.com/juju/errors/annotation_test.go:194: annotation
	github.com/juju/errors/annotation_test.go:195:
	github.com/juju/errors/annotation_test.go:196: more context
	github.com/juju/errors/annotation_test.go:197:

The first error was generated by an external system, so there was no location
associated. The second, fourth, and last lines were generated with Trace calls,
and the other two through Annotate.

Sometimes when responding to an error you want to return a more specific error
for the situation.

    if err := FindField(field); err != nil {
	    return errors.Wrap(err, errors.NotFoundf(field))
	}

This returns an error where the complete error stack is still available, and
`errors.Cause()` will return the `NotFound` error.

*/
package errors
