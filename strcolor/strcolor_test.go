package strcolor

import "testing"

func TestBlue(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print Blue", args{"Blue"}, "\u001b[34mBlue\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Blue(tt.args.str); got != tt.want {
				t.Errorf("Blue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreenBright(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print GreenBright", args{"GreenBright"}, "\u001b[92mGreenBright\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreenBright(tt.args.str); got != tt.want {
				t.Errorf("GreenBright() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRed(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print Red", args{"Red"}, "\u001b[31mRed\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Red(tt.args.str); got != tt.want {
				t.Errorf("Red() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYellowBright(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print YellowBright", args{"YellowBright"}, "\u001b[93mYellowBright\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YellowBright(tt.args.str); got != tt.want {
				t.Errorf("YelloBright() = %v, want %v", got, tt.want)
			}
		})
	}
}
