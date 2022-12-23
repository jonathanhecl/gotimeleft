# GoTimeLeft
[![Go](https://github.com/jonathanhecl/gotimeleft/actions/workflows/go.yml/badge.svg)](https://github.com/jonathanhecl/gotimeleft/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jonathanhecl/gotimeleft)](https://goreportcard.com/report/github.com/jonathanhecl/gotimeleft)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

> It is used to know how much estimated time is left to finish a task.
> 
> Works like a progress bar too.

## Example: 

![Image](https://i.imgur.com/MhitUfV.png)
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
  timeleft.GetProgress(2) // => 55.33% string with 2 decimals
  timeleft.GetProgressBar(30) // [==============>...............] string with 30 chars
  timeleft.GetFloat64() // => 0.55 float64
  timeleft.GetPerSecond() // => 5.55 float64 per second
  timeleft.GetTimeLeft() // => 0.5ms time.Duration
  timeleft.GetTimeSpent() // => 2s time.Duration
```
