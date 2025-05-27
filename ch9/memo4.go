package memo 

type entry struct{
	res result 
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	f Func 
	mu sync.Mutex 
	cache map[string]*entry
}

type Func func(key string)(interface{}, err)

type result struct{
	value interface{}
	err error 
}

func New(f Func) *Memo{
	return &Memo{f: f, cache: make(map[string]*result)}
}

// concurrency-safe
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