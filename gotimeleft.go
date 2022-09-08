package gotimeleft

import (
	"strconv"
	"time"
)

type (
	TimeLeft struct {
		totalValues         int
		initializationTime  time.Time
		speedPerMillisecond float64
		lastValue           int
		lastStepTime        time.Time
	}
)

func Init(newTotal int) *TimeLeft {
	return &TimeLeft{
		totalValues:         newTotal,
		initializationTime:  time.Now(),
		speedPerMillisecond: 0,
		lastValue:           0,
		lastStepTime:        time.Now(),
	}
}

func (t *TimeLeft) Reset(newTotal int) *TimeLeft {
	t.initializationTime = time.Now()
	t.totalValues = newTotal
	t.speedPerMillisecond = 0
	t.lastValue = 0
	t.lastStepTime = time.Now()

	return t
}

func (t *TimeLeft) Step(newStep int) *TimeLeft {
	change := newStep

	if change > t.totalValues {
		change = t.totalValues - t.lastValue
		newStep = change
	}
	elapsed := time.Since(t.lastStepTime)
	speedPerMillisecond := float64(change) / float64(elapsed.Milliseconds())

	if t.speedPerMillisecond == 0 {
		t.speedPerMillisecond = speedPerMillisecond
	} else {
		t.speedPerMillisecond = (t.speedPerMillisecond + speedPerMillisecond) / 2
	}
	t.lastValue = t.lastValue + newStep
	t.lastStepTime = time.Now()

	return t
}

func (t *TimeLeft) Value(newValue int) *TimeLeft {
	change := newValue - t.lastValue

	if change+newValue > t.totalValues {
		change = t.totalValues - t.lastValue
		newValue = t.totalValues
	}

	elapsed := time.Since(t.lastStepTime)
	speedPerMillisecond := float64(change) / float64(elapsed.Milliseconds())

	if t.speedPerMillisecond == 0 {
		t.speedPerMillisecond = speedPerMillisecond
	} else {
		t.speedPerMillisecond = (t.speedPerMillisecond + speedPerMillisecond) / 2
	}

	t.lastValue = newValue
	t.lastStepTime = time.Now()

	return t
}

func (t *TimeLeft) GetProgressValues() string {
	return strconv.Itoa(t.lastValue) + "/" + strconv.Itoa(t.totalValues)
}

func (t *TimeLeft) GetProgress(prec int) string { // 10.1% 15.5%
	return strconv.FormatFloat(float64(t.lastValue)/float64(t.totalValues)*100, 'f', prec, 64) + "%"
}

func (t *TimeLeft) GetFloat64() float64 {
	return float64(t.lastValue) / float64(t.totalValues)
}

func (t *TimeLeft) GetTimeLeft() time.Duration {
	return time.Duration(float64(t.totalValues-t.lastValue)/t.speedPerMillisecond) * time.Millisecond
}

func (t *TimeLeft) GetTimeSpent() time.Duration {
	return time.Since(t.initializationTime)
}

func (t *TimeLeft) GetPerSecond() float64 {
	return t.speedPerMillisecond * 1000
}
