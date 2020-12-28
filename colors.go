package captcha

var (
	GIFEnd    = []byte("\x01\x11\x00;")
	colorSize = 13 + 48 + 10 + 1
)

type color []byte

var (
	Black = []string{
		"GIF89a", "\xc8\x00\x46\x00", "\x83", "\x00\x00",
		"\xff\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\x00\x00\x00",
		"\xff\xff\xff", ",", "\x00\x00\x00\x00", "\xc8\x00\x46\x00", "\x00", "\x04",
	}

	DeepOrange = []string{
		"GIF89a", "\xc8\x00\x46\x00", "\x83", "\x00\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xDD\x2C\x00",
		"\xff\xff\xff", ",", "\x00\x00\x00\x00", "\xc8\x00\x46\x00", "\x00", "\x04",
	}

	Blue = []string{
		"GIF89a", "\xc8\x00\x46\x00", "\x83", "\x00\x00",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\x29\x62\xFF",
		"\xff\xff\xff", ",", "\x00\x00\x00\x00", "\xc8\x00\x46\x00", "\x00", "\x04",
	}

	Pink = []string{
		"GIF89a", "\xc8\x00\x46\x00", "\x83", "\x00\x00",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xC5\x11\x62",
		"\xff\xff\xff", ",", "\x00\x00\x00\x00", "\xc8\x00\x46\x00", "\x00", "\x04",
	}

	DeepPurple = []string{
		"GIF89a", "\xc8\x00\x46\x00", "\x83", "\x00\x00",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\x62\x00\xEA",
		"\xff\xff\xff", ",", "\x00\x00\x00\x00", "\xc8\x00\x46\x00", "\x00", "\x04",
	}

	colors [][]byte
)

//把[]string统统转成[]byte
func loadColors() {

	c := [][]string{DeepPurple, Black, DeepOrange, Blue, Pink}

	clrs := make([][]byte, len(c))
	for _, color := range c {
		tmp := make([]byte, colorSize)

		for _, chars := range color {
			tmp = append(tmp, chars...)
		}

		clrs = append(clrs, tmp)
	}

	colors = clrs
}