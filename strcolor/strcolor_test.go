package strcolor

import "testing"

func TestBlack(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print Black", args{"Black"}, "\u001b[30mBlack\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Black(tt.args.str); got != tt.want {
				t.Errorf("Black() = %v, want %v", got, tt.want)
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

func TestGreen(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print Green", args{"Green"}, "\u001b[32mGreen\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Green(tt.args.str); got != tt.want {
				t.Errorf("Green() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYellow(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print Yellow", args{"Yellow"}, "\u001b[33mYellow\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Yellow(tt.args.str); got != tt.want {
				t.Errorf("Yellow() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestMagenta(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print Magenta", args{"Magenta"}, "\u001b[35mMagenta\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Magenta(tt.args.str); got != tt.want {
				t.Errorf("Magenta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCyan(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print Cyan", args{"Cyan"}, "\u001b[36mCyan\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cyan(tt.args.str); got != tt.want {
				t.Errorf("Cyan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhite(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print White", args{"White"}, "\u001b[37mWhite\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := White(tt.args.str); got != tt.want {
				t.Errorf("White() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlackBright(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print BlackBright", args{"BlackBright"}, "\u001b[90mBlackBright\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlackBright(tt.args.str); got != tt.want {
				t.Errorf("BlackBright() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedBright(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print RedBright", args{"RedBright"}, "\u001b[91mRedBright\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RedBright(tt.args.str); got != tt.want {
				t.Errorf("RedBright() = %v, want %v", got, tt.want)
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
				t.Errorf("YellowBright() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlueBright(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print BlueBright", args{"BlueBright"}, "\u001b[94mBlueBright\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlueBright(tt.args.str); got != tt.want {
				t.Errorf("BlueBright() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagentaBright(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print MagentaBright", args{"MagentaBright"}, "\u001b[95mMagentaBright\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MagentaBright(tt.args.str); got != tt.want {
				t.Errorf("MagentaBright() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCyanBright(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print CyanBright", args{"CyanBright"}, "\u001b[96mCyanBright\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CyanBright(tt.args.str); got != tt.want {
				t.Errorf("CyanBright() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhiteBright(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Print WhiteBright", args{"WhiteBright"}, "\u001b[97mWhiteBright\u001b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhiteBright(tt.args.str); got != tt.want {
				t.Errorf("WhiteBright() = %v, want %v", got, tt.want)
			}
		})
	}
}
