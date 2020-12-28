package captcha

var DefaultConfig = Config{
	Line:     true,
	FontSize: 6,
}

type Config struct {
	Line     bool
	FontSize int
}
