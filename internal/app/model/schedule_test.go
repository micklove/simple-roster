package model

import (
	"testing"
	"time"
)

func TestSchedule_GetShifts(t *testing.T) {
	shiftCount := 10
	shifts := createListOfShiftsWithDifferentNames(t, "some-roster-id", shiftCount)

	type fields struct {
		ID     int64
		Name   string
		Shifts Shifts
	}
	type args struct {
		startTime time.Time
		endTime   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test 1",
			fields: fields{
				10,
				"hello",
				shifts,
			},
			args: args{
				startTime: time.Now(),
				endTime:   time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schedule := &Schedule{
				ID:     tt.fields.ID,
				Name:   tt.fields.Name,
				Shifts: tt.fields.Shifts,
			}
			schedule.GetShifts(tt.args.startTime, tt.args.endTime)
		})
	}
}
