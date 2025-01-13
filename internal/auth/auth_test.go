package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey validapikey"},
			},
			want:    "validapikey",
			wantErr: nil,
		},
		{
			name: "No Authorization Header",
			headers: http.Header{
				"SomeHeader": []string{"value"},
			},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization Header",
			headers: http.Header{
				"Authorization": []string{"incorrect"},
			},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name: "Empty Authorization Header",
			headers: http.Header{
				"Authorization": []string{""},
			},
			want:    "",
			wantErr: errors.New("no authorization header included"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) || (err != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
