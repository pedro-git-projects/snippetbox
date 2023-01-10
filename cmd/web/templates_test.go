package main

import (
	"testing"
	"time"
)

func TestHumanData(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2023, 1, 9, 10, 25, 48, 0, time.UTC),
			want: "09 Jan 2023 at 10:25",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2023, 1, 9, 10, 25, 48, 0, time.FixedZone("CET", 1*60*60)),
			want: "09 Jan 2023 at 10:25",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hd := humanDate(test.tm)

			if hd != test.want {
				t.Errorf("got %q; want %q", hd, test.want)
			}
		})
	}
}
