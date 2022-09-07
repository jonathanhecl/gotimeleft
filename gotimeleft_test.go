package gotimeleft

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInit(t *testing.T) {

	var sameTime = time.Now()

	type args struct {
		total int
	}

	tests := []struct {
		name    string
		args    args
		want    *TimeLeft
		checker func(expected, got *TimeLeft)
	}{
		{
			name: "Total 100",
			args: args{
				total: 100,
			},
			want: &TimeLeft{
				Total:               100,
				InitializationTime:  sameTime,
				SpeedPerMillisecond: 0,
				LastValue:           0,
				LastStepTime:        sameTime,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "Total unset",
			args: args{},
			want: &TimeLeft{
				Total:               0,
				InitializationTime:  sameTime,
				SpeedPerMillisecond: 0,
				LastValue:           0,
				LastStepTime:        sameTime,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Init(tt.args.total)
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_GetProgress(t *testing.T) {

	type fields struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMillisecond float64
		LastValue           int
		LastStepTime        time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		want    float64
		checker func(expected, got float64)
	}{
		{
			name: "Total 100, LastValue 50",
			fields: fields{
				Total:     100,
				LastValue: 50,
			},
			want: 0.5,
			checker: func(expected, got float64) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "Total 100, LastValue 0",
			fields: fields{
				Total:     100,
				LastValue: 0,
			},
			want: 0,
			checker: func(expected, got float64) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "Total 100, LastValue 100",
			fields: fields{
				Total:     100,
				LastValue: 100,
			},
			want: 1,
			checker: func(expected, got float64) {
				assert.Equal(t, expected, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &TimeLeft{
				Total:               tt.fields.Total,
				InitializationTime:  tt.fields.InitializationTime,
				SpeedPerMillisecond: tt.fields.SpeedPerMillisecond,
				LastValue:           tt.fields.LastValue,
				LastStepTime:        tt.fields.LastStepTime,
			}

			got := t.GetProgress()
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_GetProgressString(t *testing.T) {

	type fields struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMillisecond float64
		LastValue           int
		LastStepTime        time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		want    string
		checker func(expected, got string)
	}{
		{
			name: "Total 100, LastValue 50",
			fields: fields{
				Total:     100,
				LastValue: 50,
			},
			want: "50/100",
			checker: func(expected, got string) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "Total 100, LastValue 0",
			fields: fields{
				Total:     100,
				LastValue: 0,
			},
			want: "0/100",
			checker: func(expected, got string) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "Total 100, LastValue 100",
			fields: fields{
				Total:     100,
				LastValue: 100,
			},
			want: "100/100",
			checker: func(expected, got string) {
				assert.Equal(t, expected, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {

			t := &TimeLeft{
				Total:               tt.fields.Total,
				InitializationTime:  tt.fields.InitializationTime,
				SpeedPerMillisecond: tt.fields.SpeedPerMillisecond,
				LastValue:           tt.fields.LastValue,
				LastStepTime:        tt.fields.LastStepTime,
			}

			got := t.GetProgressString()
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_GetTimeLeft(t *testing.T) {

	type fields struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMillisecond float64
		LastValue           int
		LastStepTime        time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		want    time.Duration
		checker func(expected, got time.Duration)
	}{
		{
			name: "No time left",
			fields: fields{
				Total:               100,
				LastValue:           0,
				SpeedPerMillisecond: 0,
			},
			want: 0,
			checker: func(expected, got time.Duration) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "25ms time left",
			fields: fields{
				Total:               100,
				LastValue:           50,
				SpeedPerMillisecond: 2,
			},
			want: 25 * time.Millisecond,
			checker: func(expected, got time.Duration) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "39.13s time left",
			fields: fields{
				Total:               100,
				LastValue:           10,
				SpeedPerMillisecond: 0.0023,
			},
			want: 39*time.Second + 130*time.Millisecond,
			checker: func(expected, got time.Duration) {
				assert.Equal(t, expected, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &TimeLeft{
				Total:               tt.fields.Total,
				InitializationTime:  tt.fields.InitializationTime,
				SpeedPerMillisecond: tt.fields.SpeedPerMillisecond,
				LastValue:           tt.fields.LastValue,
				LastStepTime:        tt.fields.LastStepTime,
			}

			got := t.GetTimeLeft()
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_GetTimeSpent(t *testing.T) {

	var sameTime = time.Now()

	type fields struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMillisecond float64
		LastValue           int
		LastStepTime        time.Time
	}

	tests := []struct {
		name       string
		baseFields fields
		want       time.Duration
		checker    func(expected, got time.Duration)
	}{
		{
			name: "Same time",
			baseFields: fields{
				InitializationTime: sameTime,
			},
			want: 0,
			checker: func(expected, got time.Duration) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "1 hour ago",
			baseFields: fields{
				InitializationTime: sameTime.Add(-1 * time.Hour),
			},
			want: time.Hour,
			checker: func(expected, got time.Duration) {
				assert.Equal(t, expected, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {

			t := &TimeLeft{
				Total:               tt.baseFields.Total,
				InitializationTime:  tt.baseFields.InitializationTime,
				SpeedPerMillisecond: tt.baseFields.SpeedPerMillisecond,
				LastValue:           tt.baseFields.LastValue,
				LastStepTime:        tt.baseFields.LastStepTime,
			}

			got := t.GetTimeSpent()
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_Reset(t *testing.T) {

	type fields struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMillisecond float64
		LastValue           int
		LastStepTime        time.Time
	}

	type args struct {
		total int
	}

	tests := []struct {
		name       string
		baseFields fields
		args       args
		want       *TimeLeft
		checker    func(expected, got *TimeLeft)
	}{
		{
			name: "Total 99",
			baseFields: fields{
				Total:     100,
				LastValue: 100,
			},
			args: args{
				total: 99,
			},
			want: &TimeLeft{
				Total:               99,
				InitializationTime:  time.Now(),
				SpeedPerMillisecond: 0,
				LastValue:           0,
				LastStepTime:        time.Now(),
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.Total, got.Total)
				assert.Equal(t, expected.LastValue, got.LastValue)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {

			t := &TimeLeft{
				Total:               tt.baseFields.Total,
				InitializationTime:  tt.baseFields.InitializationTime,
				SpeedPerMillisecond: tt.baseFields.SpeedPerMillisecond,
				LastValue:           tt.baseFields.LastValue,
				LastStepTime:        tt.baseFields.LastStepTime,
			}

			got := t.Reset(tt.args.total)
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_Step(t *testing.T) {

	var sameTime = time.Now()

	type fields struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMillisecond float64
		LastValue           int
		LastStepTime        time.Time
	}

	type args struct {
		newStep int
	}

	tests := []struct {
		name       string
		baseFields fields
		args       args
		want       *TimeLeft
		checker    func(expected, got *TimeLeft)
	}{
		{
			name: "First step",
			baseFields: fields{
				Total:               100,
				SpeedPerMillisecond: 0.0,
				LastValue:           0,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newStep: 2,
			},
			want: &TimeLeft{
				Total:               100,
				SpeedPerMillisecond: 0.002,
				LastValue:           2,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.Total, got.Total)
				assert.Equal(t, expected.LastValue, got.LastValue)
				assert.Equal(t, expected.SpeedPerMillisecond, got.SpeedPerMillisecond)
			},
		},
		{
			name: "Second step",
			baseFields: fields{
				Total:               100,
				SpeedPerMillisecond: 0.002,
				LastValue:           2,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newStep: 8,
			},
			want: &TimeLeft{
				Total:               100,
				SpeedPerMillisecond: 0.005,
				LastValue:           10,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.Total, got.Total)
				assert.Equal(t, expected.LastValue, got.LastValue)
				assert.Equal(t, expected.SpeedPerMillisecond, got.SpeedPerMillisecond)
			},
		},
		{
			name: "Exceeding step",
			baseFields: fields{
				Total:               100,
				SpeedPerMillisecond: 0.002,
				LastValue:           2,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newStep: 101,
			},
			want: &TimeLeft{
				Total:               100,
				SpeedPerMillisecond: 0.05,
				LastValue:           100,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.Total, got.Total)
				assert.Equal(t, expected.LastValue, got.LastValue)
				assert.Equal(t, expected.SpeedPerMillisecond, got.SpeedPerMillisecond)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &TimeLeft{
				Total:               tt.baseFields.Total,
				InitializationTime:  tt.baseFields.InitializationTime,
				SpeedPerMillisecond: tt.baseFields.SpeedPerMillisecond,
				LastValue:           tt.baseFields.LastValue,
				LastStepTime:        tt.baseFields.LastStepTime,
			}

			got := t.Step(tt.args.newStep)
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_Value(t *testing.T) {

	var sameTime = time.Now()

	type fields struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMillisecond float64
		LastValue           int
		LastStepTime        time.Time
	}

	type args struct {
		newValue int
	}

	tests := []struct {
		name       string
		baseFields fields
		args       args
		want       *TimeLeft
		checker    func(expected, got *TimeLeft)
	}{
		{
			name: "First value",
			baseFields: fields{
				Total:               100,
				SpeedPerMillisecond: 0.0,
				LastValue:           0,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newValue: 2,
			},
			want: &TimeLeft{
				Total:               100,
				SpeedPerMillisecond: 0.002,
				LastValue:           2,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.Total, got.Total)
				assert.Equal(t, expected.LastValue, got.LastValue)
				assert.Equal(t, expected.SpeedPerMillisecond, got.SpeedPerMillisecond)
			},
		},
		{
			name: "Second value",
			baseFields: fields{
				Total:               100,
				SpeedPerMillisecond: 0.002,
				LastValue:           2,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newValue: 10,
			},
			want: &TimeLeft{
				Total:               100,
				SpeedPerMillisecond: 0.005,
				LastValue:           10,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.Total, got.Total)
				assert.Equal(t, expected.LastValue, got.LastValue)
				assert.Equal(t, expected.SpeedPerMillisecond, got.SpeedPerMillisecond)
			},
		},
		{
			name: "Exceeding value",
			baseFields: fields{
				Total:               100,
				SpeedPerMillisecond: 0.002,
				LastValue:           2,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newValue: 101,
			},
			want: &TimeLeft{
				Total:               100,
				SpeedPerMillisecond: 0.0505,
				LastValue:           100,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.Total, got.Total)
				assert.Equal(t, expected.LastValue, got.LastValue)
				assert.Equal(t, expected.SpeedPerMillisecond, got.SpeedPerMillisecond)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &TimeLeft{
				Total:               tt.baseFields.Total,
				InitializationTime:  tt.baseFields.InitializationTime,
				SpeedPerMillisecond: tt.baseFields.SpeedPerMillisecond,
				LastValue:           tt.baseFields.LastValue,
				LastStepTime:        tt.baseFields.LastStepTime,
			}

			got := t.Value(tt.args.newValue)
			tt.checker(tt.want, got)
		})
	}
}
