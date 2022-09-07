# GoTimeLeft
[![Go Report Card](https://goreportcard.com/badge/github.com/jonathanhecl/gotimeleft)](https://goreportcard.com/report/github.com/jonathanhecl/gotimeleft)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

    With this package you can get the time left, with a knowed total.

## Example: 
![Image](https://i.imgur.com/mKIGzX5.png)
![Image](https://i.imgur.com/2vVI9qM.png)

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
  
  timeleft.GetProgressValues() // => 55/100 string
  timeleft.GetProgress() // => 55% string
  timeleft.GetFloat64() // => 0.55 float64
  timeleft.GetTimeLeft() // => 0.5ms time.Duration
  timeleft.GetTimeSpent() // => 2s time.Duration
```
