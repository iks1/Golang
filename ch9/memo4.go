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
	e:= memo.cache[key]
	if e == nil {
		// this is the first request for this key 
		// this goroutine becomes responsible for computing
		// the value and broadcasting the ready condition
		e = &entry{ready: make(chan struct{})}
		memo.cache[key]=e 
		memo.mu.Unlock() 
		e.res.value, e.res.err = memo.f(key) 
		close(e.ready) // broadcast ready condition
	} else {
		//this is a repeat request for this key
		memo.mu.Unlock()
		<- e.ready // wait for ready condition
	}
	
	return e.res.value, e.res.err
}