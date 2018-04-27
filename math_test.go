package php

import (
	"testing"
)

func TestRound(t *testing.T) {
	type args struct {
		val       float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"test1",
			args{
				3.4,
				0,
			},
			3,
		},
		{
			"test2",
			args{
				3.5,
				0,
			},
			4,
		},
		{
			"test3",
			args{
				3.6,
				0,
			},
			4,
		},
		{
			"test4",
			args{
				1.95583,
				2,
			},
			1.96,
		},
		{
			"test5",
			args{
				1241757,
				-3,
			},
			1242000,
		},
		{
			"test6",
			args{
				5.045,
				2,
			},
			5.05,
		},
		{
			"test6",
			args{
				5.055,
				2,
			},
			5.06,
		},
		{
			"test7",
			args{
				1.55,
				1,
			},
			1.6,
		},
		{
			"test8",
			args{
				1.54,
				1,
			},
			1.5,
		},
		{
			"test9",
			args{
				-1.55,
				1,
			},
			-1.6,
		},
		{
			"test10",
			args{
				-1.54,
				1,
			},
			-1.5,
		},
		{
			"test11",
			args{
				0,
				1,
			},
			0,
		},
		{
			"test12",
			args{
				0.12345678901234567890123456789,
				15,
			},
			0.12345678901235,
		},
		{
			"test13",
			args{
				12345678901234567890,
				-15,
			},
			1.23457e+19,
		},
		{
			"test14",
			args{
				0,
				-1,
			},
			0,
		},
		{
			"test15",
			args{
				-1.55,
				-2,
			},
			0,
		},
		{
			"test16",
			args{
				-1.55,
				-1,
			},
			0,
		},
		{
			"test17",
			args{
				-1.55,
				0,
			},
			-2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.val, tt.args.precision); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}