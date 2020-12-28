package captcha

import (
	"crypto/rand"
	"strings"
)

const gifsize = 17646

var sw = [200]int{0, 4, 8, 12, 16, 20, 23, 27, 31, 35, 39, 43, 47, 50, 54, 58, 61, 65, 68, 71, 75, 78, 81, 84, 87, 90, 93, 96, 98, 101, 103, 105, 108, 110, 112, 114, 115, 117, 119, 120, 121, 122, 123, 124, 125, 126, 126, 127, 127, 127, 127, 127, 127, 127, 126, 126, 125, 124, 123, 122, 121, 120, 119, 117, 115, 114, 112, 110, 108, 105, 103, 101, 98, 96, 93, 90, 87, 84, 81, 78, 75, 71, 68, 65, 61, 58, 54, 50, 47, 43, 39, 35, 31, 27, 23, 20, 16, 12, 8, 4, 0, -4, -8, -12, -16, -20, -23, -27, -31, -35, -39, -43, -47, -50, -54, -58, -61, -65, -68, -71, -75, -78, -81, -84, -87, -90, -93, -96, -98, -101, -103, -105, -108, -110, -112, -114, -115, -117, -119, -120, -121, -122, -123, -124, -125, -126, -126, -127, -127, -127, -127, -127, -127, -127, -126, -126, -125, -124, -123, -122, -121, -120, -119, -117, -115, -114, -112, -110, -108, -105, -103, -101, -98, -96, -93, -90, -87, -84, -81, -78, -75, -71, -68, -65, -61, -58, -54, -50, -47, -43, -39, -35, -31, -27, -23, -20, -16, -12, -8, -4}
var letters = "abcdafahijklmnopqrstuvwxyz"

type Builder interface {
	Build() ([]byte, string, error)
}

type CaptchaBuilder struct {
	config Config

	gif     []byte
	matrix  []byte
	letters []byte
	swr     []byte // 200 len
	marks   []byte // 2 len
}

func NewBuilder(config Config) Builder {
	builder := &CaptchaBuilder{config: config}
	builder.reset()

	return builder
}

func (builder *CaptchaBuilder) Build() ([]byte, string, error) {
	defer builder.reset()

	if err := builder.makeCaptcha(); err != nil {
		return nil, "", err
	}

	builder.makegif()

	for i := range builder.letters {
		builder.letters[i] = letters[builder.letters[i]%25]
	}

	return builder.gif, string(builder.letters), nil
}

func (builder *CaptchaBuilder) reset() {
	//builder.pos = 30
	if builder.marks == nil {
		builder.marks = make([]byte, 2)
	}

	if builder.letters == nil {
		builder.letters = make([]byte, builder.config.FontSize)
	}

	if builder.swr == nil {
		builder.swr = make([]byte, 200)
	}

	if builder.matrix == nil {
		builder.matrix = make([]byte, 70*200)
		for i := range builder.matrix {
			builder.matrix[i] = 0xFF
		}
	}

	if builder.gif == nil {
		builder.gif = make([]byte, gifsize)
	}
}

func (builder *CaptchaBuilder) makeCaptcha() (err error) {
	if _, err = rand.Read(builder.letters); err != nil {
		return
	}

	if _, err = rand.Read(builder.marks); err != nil {
		return
	}

	if _, err = rand.Read(builder.swr); err != nil {
		return
	}

	if builder.config.Line {
		builder.drawLine()
	}

	pos := 30
	for x := 0; x < builder.config.FontSize; x++ {
		pos = builder.drawLetter(builder.letters[x]%25, pos)
	}

	return
}

func (builder *CaptchaBuilder) makegif() {
	color := []byte(strings.Join(DeepPurple, ""))
	copy(builder.gif, color)

	var (
		a, b, c, d byte
		mxNext     int
		gifNext    int = 13 + 48 + 10 + 1
	)

	for y := 0; y < 70; y++ {
		builder.gif[gifNext] += 250
		gifNext++
		for x := 0; x < 50; x++ {
			a = builder.matrix[mxNext] >> 4
			b = builder.matrix[mxNext+1] >> 4
			c = builder.matrix[mxNext+2] >> 4
			d = builder.matrix[mxNext+3] >> 4

			builder.gif[gifNext] = 16 | (a << 5)
			builder.gif[gifNext+1] = (a >> 3) | 64 | (b << 7)
			builder.gif[gifNext+2] = b >> 1
			builder.gif[gifNext+3] = 1 | (c << 1)
			builder.gif[gifNext+4] = 4 | (d << 3)

			mxNext += 4
			gifNext += 5
		}
	}

	builder.gif[gifsize-4] = '\x01'
	builder.gif[gifsize-3] = '\x11'
	builder.gif[gifsize-2] = '\x00'
	builder.gif[gifsize-1] = ';'
	return
}

func (builder *CaptchaBuilder) drawLine() {
	var (
		i   int
		sk1 = int(builder.marks[0])
	)

	for x := 0; x < 199; x++ {
		if sk1 >= 200 {
			sk1 = sk1 % 200
		}

		skew := int(sw[sk1] / 20)
		sk1 += (int(builder.swr[x]) & 3) + 1
		i = 200*(45+skew) + x

		builder.matrix[i] = 0
		builder.matrix[i+1] = 0
		builder.matrix[i+200] = 0
		builder.matrix[i+201] = 0
	}
}

func (builder *CaptchaBuilder) drawLetter(n byte, pos int) int {
	var (
		rpos = 200*16 + pos
		ipos = rpos
		sk1  = int(builder.marks[0]) + pos
		sk2  = int(builder.marks[1]) + pos
		mpos = pos
		row  int
		font = lt[n]
	)

	for _, pixel := range font {
		if pixel == -101 {
			break
		}

		if pixel < 0 {
			if pixel == -100 {
				rpos += 200
				ipos = rpos
				sk1 = int(builder.marks[0]) + pos
				row++
				continue
			}

			ipos += -pixel
			continue
		}

		skew := sw[sk1%200] / 16

		sk1 += int(builder.swr[(pos+ipos-rpos)%200]&0x1) + 1

		skewh := sw[sk2%200] / 70
		sk2 += int(builder.swr[row] & 0x1)

		if tmpPos := pos + ipos - rpos; mpos < tmpPos {
			mpos = tmpPos
		}

		if x := ipos + (skew * 200) + skewh; x < 70*200 {
			builder.matrix[x] = byte(pixel << 4)
		}

		ipos++

	}

	return mpos + 3
}
