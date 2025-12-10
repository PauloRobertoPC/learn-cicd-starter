package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t1 := http.Header{"Authorization": []string{""}}
	t2 := http.Header{"Authorization": []string{"fasfddasfs"}}
	t3 := http.Header{"Authorization": []string{"ApiKey"}}
	t4 := http.Header{"Authorization": []string{"ApiKey key-of-access"}}
	tests := map[string]struct {
		input       http.Header
		expectedS   string
		expectedErr bool
	}{
		"Empty":     {input: t1, expectedS: "", expectedErr: true},
		"No Method": {input: t2, expectedS: "", expectedErr: true},
		"No ApiKey": {input: t3, expectedS: "", expectedErr: true},
		"ApiKey":    {input: t4, expectedS: "key-of-accesss", expectedErr: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if err != nil && !tc.expectedErr {
				t.Fatalf("error was not expected\n")
			}
			if got != tc.expectedS {
				t.Fatalf("expected: %s, got: %s\n", tc.expectedS, got)
			}
		})
	}
}
