package model

import (
	"context"
	"github.com/yoshikouki/semaphore-server/server"
	"testing"
)

func TestModel_Unlock(t *testing.T) {
	rdb, _ := server.NewRedis("localhost", 6379, "", 0)
	m, _ := NewModel(rdb)

	type args struct {
		ctx    context.Context
		target string
		user   string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   string
		wantErr bool
	}{
		{
			name: "OK",
			args: args{

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := m.Unlock(tt.args.ctx, tt.args.target, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Unlock() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Unlock() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
