# Channels

## Notes

#### https://go.dev/ref/spec#Channel_types

- A channel provides a mechanism for concurrently executing
  functions to communicate by sending and receiving
  values of a specified element type.

- A new, initialized channel value can be made using the built-in function make, which takes the channel type **and an optional capacity(size of the buffer in the channel)** as arguments

```
make(chan int, 100)
```

- If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are **ready**.
- A nil channel is never ready for communication<br>

- The multi-valued assignment form of the receive operator reports whether a received value was sent before the channel was closed.

```
x, ok := <-ch

```
