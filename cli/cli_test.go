package cli

import (
	"testing"
)

func Test_integers_String(t *testing.T) {
	type fields struct {
		values []int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ToString values: 1 2 3",
			fields: fields{
				values: []int64{1, 2, 3},
			},
			want: "1 2 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &integers{
				values: tt.fields.values,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("[FAIL] integers.String() = %v, want %v", got, tt.want)
			} else {
				t.Logf("[PASS] integers.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_integers_Set(t *testing.T) {
	type fields struct {
		values []int64
	}
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "SetSrings Used Twice",
			fields: fields{
				values: []int64{1, 2, 3},
			},
			args: args{
				v: "1,2,3",
			},
			wantErr: true,
		},
		{
			name: "SetSrings() Used once",
			fields: fields{
				values: []int64{},
			},
			args: args{
				v: "1,2,3",
			},
			wantErr: false,
		},
		{
			name: "SetSrings() Used once, with negative int",
			fields: fields{
				values: []int64{},
			},
			args: args{
				v: "1,-2,3",
			},
			wantErr: true,
		},
		{
			name: "SetSrings() Used once, with invalid int",
			fields: fields{
				values: []int64{},
			},
			args: args{
				v: "1,adnf,3",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &integers{
				values: tt.fields.values,
			}
			if err := s.Set(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("[FAIL] integers.Set() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				t.Logf("[PASS] ntegers.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

/*
func TestCmds(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Cmds()
		})
	}
}
*/
func Test_isNotValidInt(t *testing.T) {
	tests := []struct {
		input string
		want  bool
		err   bool
	}{
		{
			input: "0",
			want:  false,
			err:   false,
		},
		{
			input: "a",
			want:  true,
			err:   true,
		},
		{
			input: "-342",
			want:  true,
			err:   true,
		},
		{
			input: "3.5",
			want:  true,
			err:   true,
		},
		{
			input: "-0.1",
			want:  true,
			err:   true,
		},
		{
			input: "3455245",
			want:  false,
			err:   false,
		},
		{
			input: "0.00001",
			want:  true,
			err:   true,
		},
	}
	for _, tt := range tests {
		got, err := isNotValidInt(tt.input)
		if got != tt.want {
			t.Errorf("[FAIL] Input is [%v]: got [%v], want [%v]", tt.input, got, tt.want)
		} else if (err != nil) != tt.err {
			t.Errorf("[FAIL] Input is [%v]: got err [%v], want err [%v]", tt.input, (err != nil), tt.err)
		} else if got != tt.want && (err != nil) != tt.err {
			t.Errorf("[FAIL] Input is [%v]: got [%v], want [%v] AND got err [%v], want err [%v]", tt.input, got, tt.want, (err != nil), tt.err)
		} else {
			t.Logf("[PASS] Input OK, Input is [%v]: got [%v], want [%v] AND got err [%v], want err [%v]", tt.input, got, tt.want, (err != nil), tt.err)
		}

	}
}
