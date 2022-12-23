package gotimeleft

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
			name: "totalValues 100",
			args: args{
				total: 100,
			},
			want: &TimeLeft{
				totalValues:         100,
				initializationTime:  sameTime,
				speedPerMicrosecond: 0,
				lastValue:           0,
				lastStepTime:        sameTime,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.totalValues, got.totalValues)
				assert.Equal(t, expected.lastValue, got.lastValue)
			},
		},
		{
			name: "totalValues unset",
			args: args{},
			want: &TimeLeft{
				totalValues:         0,
				initializationTime:  sameTime,
				speedPerMicrosecond: 0,
				lastValue:           0,
				lastStepTime:        sameTime,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.totalValues, got.totalValues)
				assert.Equal(t, expected.lastValue, got.lastValue)
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

func TestTimeLeft_GetFloat64(t *testing.T) {

	type fields struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMicrosecond float64
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
			name: "totalValues 100, lastValue 50",
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
			name: "totalValues 100, lastValue 0",
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
			name: "totalValues 100, lastValue 100",
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
				totalValues:         tt.fields.Total,
				initializationTime:  tt.fields.InitializationTime,
				speedPerMicrosecond: tt.fields.SpeedPerMicrosecond,
				lastValue:           tt.fields.LastValue,
				lastStepTime:        tt.fields.LastStepTime,
			}

			got := t.GetFloat64()
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_GetProgress(t *testing.T) {

	type fields struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMicrosecond float64
		LastValue           int
		LastStepTime        time.Time
	}

	type args struct {
		precision int
	}

	tests := []struct {
		name    string
		args    args
		fields  fields
		want    string
		checker func(expected, got string)
	}{
		{
			name: "totalValues 100, lastValue 50",
			args: args{
				precision: 0,
			},
			fields: fields{
				Total:     100,
				LastValue: 50,
			},
			want: "50%",
			checker: func(expected, got string) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "totalValues 100, lastValue 0",
			args: args{
				precision: 1,
			},
			fields: fields{
				Total:     100,
				LastValue: 0,
			},
			want: "0.0%",
			checker: func(expected, got string) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "totalValues 100, lastValue 100",
			args: args{
				precision: 2,
			},
			fields: fields{
				Total:     100,
				LastValue: 100,
			},
			want: "100.00%",
			checker: func(expected, got string) {
				assert.Equal(t, expected, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &TimeLeft{
				totalValues:         tt.fields.Total,
				initializationTime:  tt.fields.InitializationTime,
				speedPerMicrosecond: tt.fields.SpeedPerMicrosecond,
				lastValue:           tt.fields.LastValue,
				lastStepTime:        tt.fields.LastStepTime,
			}

			got := t.GetProgress(tt.args.precision)
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_GetProgressBar(t *testing.T) {

	type fields struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMicrosecond float64
		LastValue           int
		LastStepTime        time.Time
	}

	type args struct {
		fullBar int
	}

	tests := []struct {
		name    string
		args    args
		fields  fields
		want    string
		checker func(expected, got string)
	}{
		{
			name: "totalValues 100, lastValue 50",
			args: args{
				fullBar: 30,
			},
			fields: fields{
				Total:     100,
				LastValue: 50,
			},
			want: "[==============>...............]",
			checker: func(expected, got string) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "totalValues 100, lastValue 0",
			args: args{
				fullBar: 30,
			},
			fields: fields{
				Total:     100,
				LastValue: 0,
			},
			want: "[..............................]",
			checker: func(expected, got string) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "totalValues 100, lastValue 100",
			args: args{
				fullBar: 0,
			},
			fields: fields{
				Total:     100,
				LastValue: 100,
			},
			want: "[==============================]",
			checker: func(expected, got string) {
				assert.Equal(t, expected, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &TimeLeft{
				totalValues:         tt.fields.Total,
				initializationTime:  tt.fields.InitializationTime,
				speedPerMicrosecond: tt.fields.SpeedPerMicrosecond,
				lastValue:           tt.fields.LastValue,
				lastStepTime:        tt.fields.LastStepTime,
			}

			got := t.GetProgressBar(tt.args.fullBar)
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_GetProgressValues(t *testing.T) {

	type fields struct {
		Total               int
		InitializationTime  time.Time
		SpeedPerMicrosecond float64
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
			name: "totalValues 100, lastValue 50",
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
			name: "totalValues 100, lastValue 0",
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
			name: "totalValues 100, lastValue 100",
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
				totalValues:         tt.fields.Total,
				initializationTime:  tt.fields.InitializationTime,
				speedPerMicrosecond: tt.fields.SpeedPerMicrosecond,
				lastValue:           tt.fields.LastValue,
				lastStepTime:        tt.fields.LastStepTime,
			}

			got := t.GetProgressValues()
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_GetTimeLeft(t *testing.T) {

	type fields struct {
		Total               int
		InitializationTime  time.Time
		speedPerMicrosecond float64
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
				speedPerMicrosecond: 0,
			},
			want: 0,
			checker: func(expected, got time.Duration) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "25µs time left",
			fields: fields{
				Total:               100,
				LastValue:           50,
				speedPerMicrosecond: 2,
			},
			want: 25 * time.Microsecond,
			checker: func(expected, got time.Duration) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "39.13ms time left",
			fields: fields{
				Total:               100,
				LastValue:           10,
				speedPerMicrosecond: 0.0023,
			},
			want: 39*time.Millisecond + 130*time.Microsecond,
			checker: func(expected, got time.Duration) {
				assert.Equal(t, expected, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &TimeLeft{
				totalValues:         tt.fields.Total,
				initializationTime:  tt.fields.InitializationTime,
				speedPerMicrosecond: tt.fields.speedPerMicrosecond,
				lastValue:           tt.fields.LastValue,
				lastStepTime:        tt.fields.LastStepTime,
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
		SpeedPerMicrosecond float64
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
				assert.Equal(t, expected, got.Round(time.Millisecond))
			},
		},
		{
			name: "1 hour ago",
			baseFields: fields{
				InitializationTime: sameTime.Add(-1 * time.Hour),
			},
			want: time.Hour,
			checker: func(expected, got time.Duration) {
				assert.Equal(t, expected, got.Round(time.Minute))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {

			t := &TimeLeft{
				totalValues:         tt.baseFields.Total,
				initializationTime:  tt.baseFields.InitializationTime,
				speedPerMicrosecond: tt.baseFields.SpeedPerMicrosecond,
				lastValue:           tt.baseFields.LastValue,
				lastStepTime:        tt.baseFields.LastStepTime,
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
		SpeedPerMicrosecond float64
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
			name: "totalValues 99",
			baseFields: fields{
				Total:     100,
				LastValue: 100,
			},
			args: args{
				total: 99,
			},
			want: &TimeLeft{
				totalValues:         99,
				initializationTime:  time.Now(),
				speedPerMicrosecond: 0,
				lastValue:           0,
				lastStepTime:        time.Now(),
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.totalValues, got.totalValues)
				assert.Equal(t, expected.lastValue, got.lastValue)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {

			t := &TimeLeft{
				totalValues:         tt.baseFields.Total,
				initializationTime:  tt.baseFields.InitializationTime,
				speedPerMicrosecond: tt.baseFields.SpeedPerMicrosecond,
				lastValue:           tt.baseFields.LastValue,
				lastStepTime:        tt.baseFields.LastStepTime,
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
		SpeedPerMicrosecond float64
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
				SpeedPerMicrosecond: 0.0,
				LastValue:           0,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newStep: 2,
			},
			want: &TimeLeft{
				totalValues:         100,
				speedPerMicrosecond: 2e-06,
				lastValue:           2,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.totalValues, got.totalValues)
				assert.Equal(t, expected.lastValue, got.lastValue)
				assert.Equal(t, expected.speedPerMicrosecond, got.speedPerMicrosecond)
			},
		},
		{
			name: "Second step",
			baseFields: fields{
				Total:               100,
				SpeedPerMicrosecond: 0.002,
				LastValue:           2,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newStep: 8,
			},
			want: &TimeLeft{
				totalValues:         100,
				speedPerMicrosecond: 0.0010040000000000001,
				lastValue:           10,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.totalValues, got.totalValues)
				assert.Equal(t, expected.lastValue, got.lastValue)
				assert.Equal(t, expected.speedPerMicrosecond, got.speedPerMicrosecond)
			},
		},
		{
			name: "Exceeding step",
			baseFields: fields{
				Total:               100,
				SpeedPerMicrosecond: 0.002,
				LastValue:           2,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newStep: 101,
			},
			want: &TimeLeft{
				totalValues:         100,
				speedPerMicrosecond: 0.001049,
				lastValue:           100,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.totalValues, got.totalValues)
				assert.Equal(t, expected.lastValue, got.lastValue)
				assert.Equal(t, expected.speedPerMicrosecond, got.speedPerMicrosecond)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &TimeLeft{
				totalValues:         tt.baseFields.Total,
				initializationTime:  tt.baseFields.InitializationTime,
				speedPerMicrosecond: tt.baseFields.SpeedPerMicrosecond,
				lastValue:           tt.baseFields.LastValue,
				lastStepTime:        tt.baseFields.LastStepTime,
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
		SpeedPerMicrosecond float64
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
				SpeedPerMicrosecond: 0.0,
				LastValue:           0,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newValue: 2,
			},
			want: &TimeLeft{
				totalValues:         100,
				speedPerMicrosecond: 0.002,
				lastValue:           2,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.totalValues, got.totalValues)
				assert.Equal(t, expected.lastValue, got.lastValue)
				assert.Equal(t, expected.speedPerMicrosecond, got.speedPerMicrosecond)
			},
		},
		{
			name: "Second value",
			baseFields: fields{
				Total:               100,
				SpeedPerMicrosecond: 0.002,
				LastValue:           2,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newValue: 10,
			},
			want: &TimeLeft{
				totalValues:         100,
				speedPerMicrosecond: 0.005,
				lastValue:           10,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.totalValues, got.totalValues)
				assert.Equal(t, expected.lastValue, got.lastValue)
				assert.Equal(t, expected.speedPerMicrosecond, got.speedPerMicrosecond)
			},
		},
		{
			name: "Exceeding value",
			baseFields: fields{
				Total:               100,
				SpeedPerMicrosecond: 0.002,
				LastValue:           2,
				InitializationTime:  sameTime.Add(-1 * time.Hour),
				LastStepTime:        sameTime.Add(-1 * time.Second),
			},
			args: args{
				newValue: 101,
			},
			want: &TimeLeft{
				totalValues:         100,
				speedPerMicrosecond: 0.05,
				lastValue:           100,
			},
			checker: func(expected, got *TimeLeft) {
				assert.Equal(t, expected.totalValues, got.totalValues)
				assert.Equal(t, expected.lastValue, got.lastValue)
				assert.Equal(t, expected.speedPerMicrosecond, got.speedPerMicrosecond)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &TimeLeft{
				totalValues:         tt.baseFields.Total,
				initializationTime:  tt.baseFields.InitializationTime,
				speedPerMicrosecond: tt.baseFields.SpeedPerMicrosecond,
				lastValue:           tt.baseFields.LastValue,
				lastStepTime:        tt.baseFields.LastStepTime,
			}

			got := t.Value(tt.args.newValue)
			tt.checker(tt.want, got)
		})
	}
}

func TestTimeLeft_GetPerSecond(t *testing.T) {

	type fields struct {
		totalValues         int
		initializationTime  time.Time
		speedPerMicrosecond float64
		lastValue           int
		lastStepTime        time.Time
	}

	tests := []struct {
		name       string
		baseFields fields
		want       float64
		checker    func(expected, got float64)
	}{
		{
			name: "First value",
			baseFields: fields{
				totalValues:         100,
				speedPerMicrosecond: 0.002,
				lastValue:           2,
			},
			want: 2.0,
			checker: func(expected, got float64) {
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "Second value",
			baseFields: fields{
				totalValues:         100,
				speedPerMicrosecond: 0.005,
				lastValue:           10,
			},
			want: 5.0,
			checker: func(expected, got float64) {
				assert.Equal(t, expected, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &TimeLeft{
				totalValues:         tt.baseFields.totalValues,
				initializationTime:  tt.baseFields.initializationTime,
				speedPerMicrosecond: tt.baseFields.speedPerMicrosecond,
				lastValue:           tt.baseFields.lastValue,
				lastStepTime:        tt.baseFields.lastStepTime,
			}

			got := t.GetPerSecond()
			tt.checker(tt.want, got)
		})
	}
}
