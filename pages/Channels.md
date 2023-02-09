- Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
- ```
  v, ok := <-ch
  ```
- `ok` is `false` if there are no more values to receive and the channel is closed.