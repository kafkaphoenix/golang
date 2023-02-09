- Fibonnaci with closure:
- ```
  func fibonacci() func() int {
    a, b := 1, 0
    return func() int {
        a, b = b, a+b
        return a
    }
  }
  ```
- Fibonacci with channel:
- ```
  func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := range c {
        c <- x
        x, y = y, x+y+i
    }
    close(c)
  }
  ```
- Fibonnaci with defer:
- ```
  func fibonacci() func() int {
    f1:= 0
    f2:= 1
    return func() int {
        defer func() { f1, f2= f2, f1+f2}()
        return f1
    }
  }
  ```