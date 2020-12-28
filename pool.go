package captcha

import "sync"

type Pool struct {
	rowPool sync.Pool
}

func NewPool(config Config) Builder {
	pool := &Pool{
		rowPool: sync.Pool{
			New: func() interface{} {
				return NewBuilder(config)
			},
		},
	}

	return pool
}

func (pool *Pool) Build() ([]byte, string, error) {
	builder := pool.rowPool.Get().(*CaptchaBuilder)
	defer pool.rowPool.Put(builder)

	return builder.Build()
}
