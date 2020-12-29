package captcha

var DefaultConfig = Config{
	Line:     true,
	FontSize: 6,
	Colors:   Colors,
}

type Config struct {
	Line     bool
	FontSize int
	Colors   []Color
}
