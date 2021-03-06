package golog

import (
	"os"
	"testing"
)

func TestNewFileLogger(t *testing.T) {
	type args struct {
		filename string
		level    int
	}
	tests := []struct {
		name    string
		args    args
		want    *Logger
		wantErr bool
	}{
		{
			args: args{
				filename: "test.log",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFileLogger(tt.args.filename, tt.args.level, 0)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFileLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.Infoln("heheda")
			got.Println("not prefix")
			a := map[string]string{
				"ka": "va",
				"kb": "vb",
			}
			got.PrintJson(a)
			got.PrintJson(a, a, a)
		})
	}
}

func TestNewLogger(t *testing.T) {
	type args struct {
		level int
	}
	tests := []struct {
		name    string
		args    args
		want    *Logger
		wantOut string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLogger(os.Stdout, tt.args.level)
			got.Infoln("heheda")
		})
	}
}
