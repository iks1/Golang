//package io

type Reader interface{
	Read(p []byte) (n int, err error)
}

type Closer interface{
	Close() error
}

//combinations 
type ReadWriter interface{
	Reader 
	Writer
}

type ReadWriteCloser interface{
	Reader 
	Writer 
	Closer 
}


// also could have writtern as 

type ReadWriter interface{
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// or combinations of both 
type ReadWriter interface{
	Read(p []byte) (n int, err error)
	Writer 
}

// All three declarations have the same effect