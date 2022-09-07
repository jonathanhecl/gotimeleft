# GoTimeLeft

## How to install

> go get github.com/jonathanhecl/gotimeleft

## How to use

```
  timeleft := gotimeleft.Init(100) // Total 100, value 0
  
  ...

  timeleft.Reset(200) // Reset to total 200, value 0

  ...

  timeleft.Step(10) // value +10

  ...

  timeleft.Value(50) // value 50
  
  ...
  
  timeleft.GetProgressString() // => 55/100 string
  timeleft.GetProgress() // => 0.55 float64
  timeleft.GetTimeLeft() // => 0.5ms time.Duration
  timeleft.GetTimeSpent() // => 2s time.Duration
```
