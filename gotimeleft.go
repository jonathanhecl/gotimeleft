package gotimeleft

import (
	"strconv"
	"time"
)

type (
	TimeLeft struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMillisecond float64
		LastValue           int
		LastStepTime        time.Time
	}
)

func Init(total int) *TimeLeft {
	return &TimeLeft{
		Total:               total,
		InitializationTime:  time.Now(),
		SpeedPerMillisecond: 0,
		LastValue:           0,
		LastStepTime:        time.Now(),
	}
}

func (t *TimeLeft) Reset(total int) *TimeLeft {
	t.InitializationTime = time.Now()
	t.Total = total
	t.SpeedPerMillisecond = 0
	t.LastValue = 0
	t.LastStepTime = time.Now()

	return t
}

func (t *TimeLeft) Step(newStep int) *TimeLeft {
	change := newStep

	if change > t.Total {
		change = t.Total - t.LastValue
		newStep = change
	}
	elapsed := time.Since(t.LastStepTime)
	speedPerMillisecond := float64(change) / float64(elapsed.Milliseconds())

	if t.SpeedPerMillisecond == 0 {
		t.SpeedPerMillisecond = speedPerMillisecond
	} else {
		t.SpeedPerMillisecond = (t.SpeedPerMillisecond + speedPerMillisecond) / 2
	}
	t.LastValue = t.LastValue + newStep
	t.LastStepTime = time.Now()

	return t
}

func (t *TimeLeft) Value(newValue int) *TimeLeft {
	change := newValue - t.LastValue

	if change+newValue > t.Total {
		change = t.Total - t.LastValue
		newValue = t.Total
	}

	elapsed := time.Since(t.LastStepTime)
	speedPerMillisecond := float64(change) / float64(elapsed.Milliseconds())

	if t.SpeedPerMillisecond == 0 {
		t.SpeedPerMillisecond = speedPerMillisecond
	} else {
		t.SpeedPerMillisecond = (t.SpeedPerMillisecond + speedPerMillisecond) / 2
	}

	t.LastValue = newValue
	t.LastStepTime = time.Now()

	return t
}

func (t *TimeLeft) GetProgressString() string {
	return strconv.Itoa(t.LastValue) + "/" + strconv.Itoa(t.Total)
}

func (t *TimeLeft) GetProgress() float64 {
	return float64(t.LastValue) / float64(t.Total)
}

func (t *TimeLeft) GetTimeLeft() time.Duration {
	return time.Duration(float64(t.Total-t.LastValue)/t.SpeedPerMillisecond) * time.Millisecond
}

func (t *TimeLeft) GetTimeSpent() time.Duration {
	return time.Since(t.InitializationTime)
}
