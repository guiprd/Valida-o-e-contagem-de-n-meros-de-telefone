package validation

import (
	"reflect"
	"testing"
)

func TestValidation(t *testing.T) {
	type args struct {
		phones []string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Right phone numbers with +",
			args: args{
				phones: []string{
					"+19992676547",
					"+01928745674",
				},
			},
			want: []bool{
				true,
				false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validation(tt.args.phones); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validation() = %v, want %v", got, tt.want)
			}
		})
	}
}
