/*
 *    Copyright 2021 Chen Quan
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package time

import (
	"reflect"
	"testing"
	"time"
)

func TestYear(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"1",
			time.Now().Year(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Year(); got != tt.want {
				t.Errorf("Year() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYearInLocal(t *testing.T) {
	type args struct {
		location *time.Location
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{time.Local},
			time.Now().In(time.Local).Year(),
		}, {
			"2",
			args{time.UTC},
			time.Now().In(time.UTC).Year(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YearInLocal(tt.args.location); got != tt.want {
				t.Errorf("YearInLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonth(t *testing.T) {
	tests := []struct {
		name string
		want time.Month
	}{
		{
			"1",
			time.Now().Month(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Month(); got != tt.want {
				t.Errorf("Month() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonthInLocal(t *testing.T) {
	type args struct {
		location *time.Location
	}
	tests := []struct {
		name string
		args args
		want time.Month
	}{
		{
			"1",
			args{time.UTC},
			time.Now().In(time.UTC).Month(),
		},
		{
			"2",
			args{time.Local},
			time.Now().In(time.Local).Month(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonthInLocal(tt.args.location); got != tt.want {
				t.Errorf("MonthInLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"1",
			time.Now().Day(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day(); got != tt.want {
				t.Errorf("Day() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDayInLocal(t *testing.T) {
	type args struct {
		location *time.Location
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{time.Local},
			time.Now().In(time.Local).Day(),
		}, {
			"2",
			args{time.UTC},
			time.Now().In(time.UTC).Day(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DayInLocal(tt.args.location); got != tt.want {
				t.Errorf("DayInLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYearDay(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"1",
			time.Now().YearDay(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YearDay(); got != tt.want {
				t.Errorf("YearDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYearDayInLocal(t *testing.T) {
	type args struct {
		location *time.Location
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{location: time.UTC},
			time.Now().In(time.UTC).YearDay(),
		}, {
			"2",
			args{location: time.Local},
			time.Now().In(time.Local).YearDay(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YearDayInLocal(tt.args.location); got != tt.want {
				t.Errorf("YearDayInLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDateFormat(t *testing.T) {
	type args struct {
		t *time.Time
	}
	date := time.Date(2021, 1, 17, 16, 0, 0, 0, time.Local)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{&date},
			"2021-01-17",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDateFormat(tt.args.t); got != tt.want {
				t.Errorf("ToDateFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDateFormatInLocal(t *testing.T) {
	type args struct {
		t        *time.Time
		location *time.Location
	}
	date1 := time.Date(2021, 1, 17, 16, 0, 0, 0, time.Local)
	date2 := time.Date(2021, 1, 17, 16, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				t:        &date1,
				location: time.Local,
			},
			"2021-01-17",
		}, {
			"2",
			args{
				t:        &date2,
				location: time.UTC,
			},
			"2021-01-17",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDateFormatInLocal(tt.args.t, tt.args.location); got != tt.want {
				t.Errorf("ToDateFormatInLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNowDateFormat(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"1",
			time.Now().Format(DateFormat),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NowDateFormat(); got != tt.want {
				t.Errorf("NowDateFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNowDateFormatInLocal(t *testing.T) {
	type args struct {
		location *time.Location
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{time.UTC},
			time.Now().In(time.UTC).Format("2006-01-02"),
		}, {
			"2",
			args{time.Local},
			time.Now().In(time.Local).Format("2006-01-02"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NowDateFormatInLocal(tt.args.location); got != tt.want {
				t.Errorf("NowDateFormatInLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDateTimeFormatInLocal(t *testing.T) {
	type args struct {
		t        *time.Time
		location *time.Location
	}
	now := time.Now()

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				t:        &now,
				location: time.Local,
			},
			time.Now().In(time.Local).Format("2006-01-02 15:04:05"),
		}, {
			"2",
			args{
				t:        &now,
				location: time.UTC,
			},
			time.Now().In(time.UTC).Format("2006-01-02 15:04:05"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDateTimeFormatInLocal(tt.args.t, tt.args.location); got != tt.want {
				t.Errorf("ToDateTimeFormatInLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNowDateTimeFormat(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"1",
			time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NowDateTimeFormat(); got != tt.want {
				t.Errorf("NowDateTimeFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNowDateTimeFormatInLocal(t *testing.T) {
	type args struct {
		location *time.Location
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1",
			args{time.UTC},
			time.Now().In(time.UTC).Format("2006-01-02 15:04:05"),
		}, {"2",
			args{time.Local},
			time.Now().In(time.Local).Format("2006-01-02 15:04:05"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NowDateTimeFormatInLocal(tt.args.location); got != tt.want {
				t.Errorf("NowDateTimeFormatInLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDateTimeFormat(t *testing.T) {
	type args struct {
		t *time.Time
	}
	date1 := time.Date(2021, 1, 17, 16, 0, 0, 0, time.Local)
	date2 := time.Date(2021, 1, 17, 16, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{
				&date1,
			},
			date1.Format("2006-01-02 15:04:05"),
		}, {
			"",
			args{
				&date2,
			},
			date2.Format("2006-01-02 15:04:05"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDateTimeFormat(tt.args.t); got != tt.want {
				t.Errorf("ToDateTimeFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeekday(t *testing.T) {
	tests := []struct {
		name string
		want time.Weekday
	}{
		{
			"1",
			time.Now().Weekday(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Weekday(); got != tt.want {
				t.Errorf("Weekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeekdayInLocal(t *testing.T) {
	type args struct {
		location *time.Location
	}
	tests := []struct {
		name string
		args args
		want time.Weekday
	}{
		{
			"1",
			args{time.Local},
			time.Now().In(time.Local).Weekday(),
		}, {
			"2",
			args{time.UTC},
			time.Now().In(time.UTC).Weekday(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeekdayInLocal(tt.args.location); got != tt.want {
				t.Errorf("WeekdayInLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDateTimeFormat(t *testing.T) {
	type args struct {
		timeFormat string
	}
	date1 := time.Date(2021, 1, 17, 16, 0, 0, 0, time.UTC)
	date2 := time.Date(2021, 1, 17, 17, 12, 0, 0, time.UTC)
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			"1",
			args{timeFormat: "2021-01-17 16:00:00"},
			&date1,
			false,
		}, {
			"2",
			args{timeFormat: "2021-01-17 17:12:00"},
			&date2,
			false,
		}, {
			"3",
			args{timeFormat: "2021-01-17 17:12:001"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDateTimeFormat(tt.args.timeFormat)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateTimeFormat() errs = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDateTimeFormat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDateTimeFormatInLocal(t *testing.T) {
	type args struct {
		timeFormat string
		location   *time.Location
	}
	date1 := time.Date(2021, 1, 17, 16, 0, 0, 0, time.Local)
	date2 := time.Date(2021, 1, 17, 17, 12, 0, 0, time.UTC)
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			"1",
			args{
				timeFormat: "2021-01-17 16:00:00",
				location:   time.Local,
			},
			&date1,
			false,
		}, {
			"2",
			args{
				timeFormat: "2021-01-17 17:12:00",
				location:   time.UTC,
			},
			&date2,
			false,
		}, {
			"3",
			args{
				timeFormat: "2021-01-17 17:12:001",
				location:   time.UTC,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDateTimeFormatInLocal(tt.args.timeFormat, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateTimeFormatInLocal() errs = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDateTimeFormatInLocal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDateFormat(t *testing.T) {
	type args struct {
		timeFormat string
	}
	date1 := time.Date(2021, 1, 17, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2021, 1, 17, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			"1",
			args{timeFormat: "2021-01-17"},
			&date1,
			false,
		}, {
			"2",
			args{timeFormat: "2021-01-17"},
			&date2,
			false,
		}, {
			"3",
			args{timeFormat: "2021-01-17 17:12:00"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDateFormat(tt.args.timeFormat)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateFormat() errs = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDateFormat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDateFormatInLocal(t *testing.T) {
	type args struct {
		timeFormat string
		location   *time.Location
	}

	date1 := time.Date(2021, 1, 17, 0, 0, 0, 0, time.Local)
	date2 := time.Date(2021, 1, 18, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			"1",
			args{
				timeFormat: "2021-01-17",
				location:   time.Local,
			},
			&date1,
			false,
		}, {
			"2",
			args{
				timeFormat: "2021-01-18",
				location:   time.UTC,
			},
			&date2,
			false,
		},
		{
			"3",
			args{
				timeFormat: "2021-01-181",
				location:   time.UTC,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDateFormatInLocal(tt.args.timeFormat, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateFormatInLocal() errs = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDateFormatInLocal() got = %v, want %v", got, tt.want)
			}
		})
	}
}
