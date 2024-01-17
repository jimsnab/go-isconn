# Description
This is a test program for an `isConnected()` function that returns true if
a socket is connected, false if not.

It also includes a call to `Read()` with a 250ms read deadline to exercise
time-bound reads in combination with `isConnected()`.

Run it like this:

1. `go build`
2. Open a terminal for the server and run `./isconn srv`
3. Open a terminal for the client and run `./isconn cli`

Both ends will connect, and then loop, printing the result of `isConnected()`.
Stop one side or the other to test drops.

# Bug
It turns out the client call of `Recvmsg()` causes the next client `Read()`
to block. The `SetReadDeadline()` is not honored.