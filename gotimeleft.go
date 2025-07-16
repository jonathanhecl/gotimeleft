package gotimeleft

import (
	"math"
	"strconv"
	"strings"
	"time"
)

type (
	TimeLeft struct {
		totalValues         int
		initializationTime  time.Time
		speedPerMicrosecond float64
		lastValue           int
		lastStepTime        time.Time
		speedHistory        []float64
		maxHistorySize      int
	}
)

// Init creates a new TimeLeft instance
func Init(newTotal int) *TimeLeft {
	return &TimeLeft{
		totalValues:         newTotal,
		initializationTime:  time.Now(),
		speedPerMicrosecond: 0,
		lastValue:           0,
		lastStepTime:        time.Now(),
		speedHistory:        make([]float64, 0, 30), // Mantener historial de las últimas 30 mediciones
		maxHistorySize:      30,
	}
}

// Reset resets the progress
func (t *TimeLeft) Reset(newTotal int) *TimeLeft {
	t.initializationTime = time.Now()
	t.totalValues = newTotal
	t.speedPerMicrosecond = 0
	t.lastValue = 0
	t.lastStepTime = time.Now()
	t.speedHistory = make([]float64, 0, t.maxHistorySize)

	return t
}

// Step updates the progress with a new step
func (t *TimeLeft) Step(newStep int) *TimeLeft {
	if t.lastStepTime.IsZero() {
		t.lastStepTime = time.Now()
		t.lastValue = newStep
		return t
	}

	change := newStep

	if change > t.totalValues {
		change = t.totalValues - t.lastValue
		newStep = change
	}

	elapsedTime := float64(time.Since(t.lastStepTime).Microseconds())
	// Ensure minimum 1μs to prevent division by zero
	if elapsedTime < 1 {
		elapsedTime = 1
	}
	speedPerMicrosecond := float64(change) / elapsedTime

	if t.speedPerMicrosecond == 0 {
		t.speedPerMicrosecond = speedPerMicrosecond
	} else {
		t.speedPerMicrosecond = (t.speedPerMicrosecond + speedPerMicrosecond) / 2
	}
	t.lastValue = t.lastValue + newStep
	t.lastStepTime = time.Now()

	return t
}

// Value updates the progress with a new value
func (t *TimeLeft) Value(newValue int) *TimeLeft {
	if t.lastStepTime.IsZero() {
		t.lastStepTime = time.Now()
		t.lastValue = newValue
		return t
	}

	change := newValue - t.lastValue

	if change+newValue > t.totalValues {
		change = t.totalValues - t.lastValue
		newValue = t.totalValues
	}

	elapsedTime := float64(time.Since(t.lastStepTime).Microseconds())
	// Ensure minimum 1μs to prevent division by zero
	if elapsedTime < 1 {
		elapsedTime = 1
	}
	speedPerMicrosecond := float64(change) / elapsedTime

	if t.speedPerMicrosecond == 0 {
		t.speedPerMicrosecond = speedPerMicrosecond
	} else {
		t.speedPerMicrosecond = (t.speedPerMicrosecond + speedPerMicrosecond) / 2
	}

	t.lastValue = newValue
	t.lastStepTime = time.Now()

	return t
}

// GetValue returns the current value
func (t *TimeLeft) GetValue() int {
	return t.lastValue
}

// GetProgressValues returns the progress as a string (10/100)
func (t *TimeLeft) GetProgressValues() string {
	return strconv.Itoa(t.lastValue) + "/" + strconv.Itoa(t.totalValues)
}

// GetProgressBar returns a string representation of the progress bar
func (t *TimeLeft) GetProgressBar(fullBar int) string {
	if fullBar < 1 {
		fullBar = 30
	}
	percent := t.GetFloat64()
	bar := int(percent * float64(fullBar))

	if bar == 0 {
		return "[" + strings.Repeat(".", fullBar) + "]"
	} else if bar >= fullBar {
		return "[" + strings.Repeat("=", fullBar) + "]"
	} else {
		return "[" + strings.Repeat("=", bar-1) + ">" + strings.Repeat(".", fullBar-bar) + "]"
	}
}

// GetProgress returns the progress as a string (10.1% 15.5%)
func (t *TimeLeft) GetProgress(prec int) string { // 10.1% 15.5%
	return strconv.FormatFloat(float64(t.lastValue)/float64(t.totalValues)*100, 'f', prec, 64) + "%"
}

// GetFloat64 returns the progress as a float64 (0.0 to 1.0)
func (t *TimeLeft) GetFloat64() float64 {
	return float64(t.lastValue) / float64(t.totalValues)
}

// calculateAverageSpeed calculates the average speed considering the history
func (t *TimeLeft) calculateAverageSpeed(newSpeed float64) float64 {
	// Add the new speed to the history
	t.speedHistory = append(t.speedHistory, newSpeed)
	
	// Maintain the maximum history size
	if len(t.speedHistory) > t.maxHistorySize {
		t.speedHistory = t.speedHistory[1:]
	}

	// If there's not enough data, use a simple average
	if len(t.speedHistory) < 3 {
		sum := 0.0
		for _, s := range t.speedHistory {
			sum += s
		}
		return sum / float64(len(t.speedHistory))
	}

	// Calculate mean and standard deviation
	var sum, sumSq float64
	for _, s := range t.speedHistory {
		sum += s
		sumSq += s * s
	}
	mean := sum / float64(len(t.speedHistory))
	stdDev := math.Sqrt(sumSq/float64(len(t.speedHistory)) - mean*mean)

	// Filter outliers (outside 1.5 standard deviations)
	var filtered []float64
	lowerBound := mean - 1.5*stdDev
	upperBound := mean + 1.5*stdDev

	for _, s := range t.speedHistory {
		if s >= lowerBound && s <= upperBound {
			filtered = append(filtered, s)
		}
	}

	// If too few values remain after filtering, use the simple mean
	if len(filtered) < 3 {
		return mean
	}

	// Calculate weighted moving average (more weight to recent values)
	var weightedSum, weightSum float64
	for i, s := range filtered {
		// Exponential weight: more recent = higher weight
		weight := math.Exp(float64(i) / float64(len(filtered)))
		weightedSum += s * weight
		weightSum += weight
	}

	return weightedSum / weightSum
}

// GetTimeLeft returns the time left to complete the task
func (t *TimeLeft) GetTimeLeft() time.Duration {
	if t.speedPerMicrosecond <= 0 {
		// If speed is zero or negative, return a large duration instead of infinity
		return 24 * time.Hour // Default to 24 hours when unable to calculate
	}

	estimatedSpeed := t.calculateAverageSpeed(t.speedPerMicrosecond)
	if estimatedSpeed <= 0 {
		return 24 * time.Hour
	}

	return time.Duration(float64(t.totalValues-t.lastValue)/estimatedSpeed) * time.Microsecond
}

// GetTimeSpent returns the time elapsed since initialization
func (t *TimeLeft) GetTimeSpent() time.Duration {
	return time.Since(t.initializationTime)
}

// GetPerSecond returns the current speed in values per second
func (t *TimeLeft) GetPerSecond() float64 {
	return t.speedPerMicrosecond * 1000
}
