package main

import "testing"

var unknown = `{
		"id": 1,
		"name": "bob",
		"addr": {
			"street": "Lazy Lane",
			"city": "Exit",
			"zip": "99999"
		},
		"extra": 21.1
	}`

func TestContains(t *testing.T) {
	var known = []string{
		`{"id": 1}`,
		`{"extra": 21.1}`,
		`{"name": "bob"}`,
		`{"addr": {"street": "Lazy Lane", "city": "Exit"}}`,
	}

	for _, k := range known {
		if err := CheckData(k, []byte(unknown)); err != nil {
			t.Errorf("invalid: %s (%s)\n", k, err)
		}
	}
}

func TestNotContains(t *testing.T) {
	var known = []string{
		`{"id": 2}`,
		`{"pid": 1}`,
		`{"name": "bobby"}`,
		`{"first": "bob"}`,
		`{"addr": {"street": "Lazy Lane", "city": "Alpha"}}`,
		// dup the above with "funk" and "extra" to up coverage
	}

	for _, k := range known {
		if err := CheckData(k, []byte(unknown)); err == nil {
			t.Errorf("false positive: %s\n", k)
		} else {
			t.Log(err)
		}
	}
}

func TestCheckDataResults(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{`{"id": 1}`, false},        // expecting error
		{`{"pid": 1}`, true},        // expecting error
		{`{"name": "bobby"}`, true}, // expecting error
		{`{"first": "bob"}`, true},  // expecting error
		{`{"addr": {"street": "Lazy Lane", "city": "Alpha"}}`, true},

		{`{"funk":"y"}`, true},  // expecting nil
		{`{"extra":123}`, true}, // expecting nil
	}

	for _, tc := range tests {
		err := CheckData(tc.input, []byte(unknown))

		if tc.expectError && err == nil {
			t.Errorf("expected error but got nil for input: %s", tc.input)
		} else if !tc.expectError && err != nil {
			t.Errorf("expected nil but got error: %v for input: %s", err, tc.input)
		}

		// log actual result
		t.Logf("input=%s, result=%v", tc.input, err)
	}
}
