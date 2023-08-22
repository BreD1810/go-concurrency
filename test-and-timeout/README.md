Test and Timeout
===

One goroutine running a check over and over, will return a value on the `done` channel if the test is successful.

Second goroutine waiting on the result, or a timeout.

This could be useful in tests, for example, where you are waiting on a service that could have a delay for a result.

Note: Using a `chan struct{}` means that there is minimal memory usage. You can send `struct{}{}` on this channel, or close it.