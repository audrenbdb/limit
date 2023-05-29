# Throttle function

Discards any successive calls that are within a delay.

## Usage

```go 
limiter := limit.DelayLimiter{Delay: delay}
	
for i := 0; i < 100; i++ {
    limiter.Call(func() {
        fmt.Println("Hello world")
    })
}
// output: "Hello world" once.
```


