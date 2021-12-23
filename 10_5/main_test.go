package main

import "testing"

func Test_find(t *testing.T) {
	type args struct {
		ss []string
		s  string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Happy path ball",
			args: args{
				ss: []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""},
				s:  "ball",
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "Happy path at",
			args: args{
				ss: []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""},
				s:  "at",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Happy path",
			args: args{
				ss: []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", "", ""},
				s:  "car",
			},
			want:    7,
			wantErr: false,
		},
		{
			name: "No word",
			args: args{
				ss: []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", "", ""},
				s:  "walls",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "repeated word", // maybe this is not a valid case?
			args: args{
				ss: []string{"walls", "walls"},
				s:  "walls",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "only one word",
			args: args{
				ss: []string{"cam"},
				s:  "cam",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "emtpy list",
			args: args{
				ss: []string{""},
				s:  "vargas",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := find(tt.args.ss, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}
