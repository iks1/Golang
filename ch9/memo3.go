package memo 

type Memo struct {
	f Func 
	mu sync.Mutex 
	cache map[string]result
}

type Func func(key string)(interface{}, err)

type result struct{
	value interface{}
	err error 
}

func New(f Func) *Memo{
	return &Memo{f: f, cache: make(map[string]result)}
}

// concurrency-safe but might give a problem 
// two go routines might end up calling the function at the same time
// overwritting each others result
func (memo *Memo) Get(key string) (interface{}, error){
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	
	return res.value, res.err
}