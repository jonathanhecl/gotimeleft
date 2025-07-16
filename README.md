# GoTimeLeft

[![Go](https://github.com/jonathanhecl/gotimeleft/actions/workflows/go.yml/badge.svg)](https://github.com/jonathanhecl/gotimeleft/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jonathanhecl/gotimeleft)](https://goreportcard.com/report/github.com/jonathanhecl/gotimeleft)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/jonathanhecl/gotimeleft.svg)](https://pkg.go.dev/github.com/jonathanhecl/gotimeleft)

A lightweight Go library for estimating time remaining for tasks and displaying progress bars in command-line applications.

## Features

- 🚀 Accurate time estimation using weighted moving averages
- 📊 Multiple progress visualization options
- ⚡ Lightweight and dependency-free
- 🛠️ Simple and intuitive API
- 📈 Handles progress tracking for tasks of any size
- 🎨 Customizable progress bar display

## Installation

```bash
go get github.com/jonathanhecl/gotimeleft
```

## Quick Start

```go
package main

import (
	"fmt"
	"time"
	"github.com/jonathanhecl/gotimeleft"
)

func main() {
	// Initialize with total number of items
	tl := gotimeleft.Init(100)

	// Simulate work
	for i := 0; i <= 100; i++ {
		time.Sleep(50 * time.Millisecond)
		
		// Update progress (either by step or value)
		tl.Step(1)
		// or: tl.Value(i)
		
		// Display progress
		fmt.Printf("\r%s %s %s",
			tl.GetProgressBar(30),
			tl.GetProgress(1),
			tl.GetTimeLeft().Round(time.Second),
		)
	}
}
```

## Usage Examples

### Basic Progress Tracking

```go
tl := gotimeleft.Init(1000) // Initialize with total items

// Update progress
tl.Step(10)  // Increment by 10
// or
tl.Value(100) // Set absolute value

// Get current progress
progress := tl.GetFloat64()  // 0.1 (10%)
```

### Displaying Progress

```go
// Get progress bar (30 characters wide)
progressBar := tl.GetProgressBar(30) // [=========>...................]

// Get percentage
percentage := tl.GetProgress(2) // "10.50%"

// Get values as string
values := tl.GetProgressValues() // "100/1000"
```

### Time Estimation

```go
// Get time left
timeLeft := tl.GetTimeLeft() // 1h30m45s

// Get time spent
timeSpent := tl.GetTimeSpent() // 45m12s

// Get operations per second
opsPerSec := tl.GetPerSecond() // 123.45
```

## Advanced Configuration

### Customizing Progress Bar

```go
// Get a progress bar with custom width (e.g., 50 characters)
bar := tl.GetProgressBar(50)

// The progress bar will look like:
// [======================>.................................] 45.0%
```

### Resetting Progress

```go
// Reset with new total
tl.Reset(200)
```

## Best Practices

1. **Initialize Early**: Create the TimeLeft instance before starting your task
2. **Update Frequently**: Call Step() or Value() regularly for accurate time estimation
3. **Handle Completion**: Check if progress reaches 100% to handle task completion
4. **Use Appropriate Precision**: Choose the right decimal places for your progress display

## Performance

GoTimeLeft is designed to be efficient with minimal overhead. The time estimation algorithm uses a weighted moving average to provide smooth and accurate predictions.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Example Output

```[========================>......................] 45.0% 12.5s
```

![Example Output](https://i.imgur.com/MhitUfV.png)

## Author

Jonathan Hecl

---

⭐ If you find this project useful, please consider giving it a star on GitHub!

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
