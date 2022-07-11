package memo

type Memo struct {
	f     Func
	cache map[string]result
}
type Func func(url string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}
func (memo *Memo) get(name string) (interface{}, error) {
	res, ok := memo.cache[name]
	if !ok {
		res.value, res.err = memo.f(name)
		memo.cache[name] = res
	}
	return res.value, res.err
}
