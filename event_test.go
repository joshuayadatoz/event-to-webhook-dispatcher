package main

import (
	"testing"
	"time"
)

func TestValidateEventDate(t *testing.T) {
	testCases := []struct {
		name    string
		date    string
		wantErr bool
	}{
		{"ValidDate", time.Now().Format(time.RFC3339), false},
		{"InvalidDate", "2024-13-01T00:00:00Z", true},
		{"EmptyDate", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateEventDate(tc.date)
			if (err != nil) != tc.wantErr {
				t.Errorf("validateEventDate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestValidateEventData(t *testing.T) {
	validCovers := CoverData{"5", "350000"}
	invalidCovers := CoverData{"-1", "700000"}

	testCases := []struct {
		name    string
		data    EventData
		wantErr bool
	}{
		{"ValidData", EventData{"adjustment", validCovers}, false},
		{"InvalidType", EventData{"", validCovers}, true},
		{"InvalidCovers", EventData{"adjustment", invalidCovers}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateEventData(tc.data)
			if (err != nil) != tc.wantErr {
				t.Errorf("validateEventData() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestValidateCoverData(t *testing.T) {
	testCases := []struct {
		name    string
		cover   CoverData
		wantErr bool
	}{
		{"ValidCover", CoverData{"5", "700000"}, false},
		{"InvalidBikesCount", CoverData{"-5", "700000"}, true},
		{"InvalidBikeValue", CoverData{"5", "-700000"}, true},
		{"EmptyFields", CoverData{"", ""}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateCoverData(tc.cover)
			if (err != nil) != tc.wantErr {
				t.Errorf("validateCoverData() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
