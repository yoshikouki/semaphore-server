package api

import (
	"net/http"
	"testing"
)

func Test_lockIfNotExists(t *testing.T) {
	tests := []struct {
		name       string
		wantErr    bool
		wantStatus int
		params     LockIfNotExistsParams
	}{
		{
			name:       "OK",
			wantErr:    false,
			wantStatus: http.StatusOK,
			params: LockIfNotExistsParams{
				LockTarget: "org-repo-stage",
				User:       "testuser",
				TTL:        "10s",
			},
		},
		{
			name:       "Invalid TTL",
			wantErr:    false,
			wantStatus: http.StatusBadRequest,
			params: LockIfNotExistsParams{
				LockTarget: "org-repo-stage",
				User:       "testuser",
				TTL:        "InvalidTTL",
			},
		},
	}
	for _, tt := range tests {
		ctx, rec := dummyContext(t, http.MethodPost, "/semaphore/lock", tt.params)

		t.Run(tt.name, func(t *testing.T) {
			if err := lockIfNotExists(ctx); (err != nil) != tt.wantErr {
				t.Errorf("lockIfNotExists() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rec.Code != tt.wantStatus {
				t.Errorf("lockIfNotExistsstatus code does not match, expected %d, got %d.", tt.wantStatus, rec.Code)
			}
		})
	}
}
