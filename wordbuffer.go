package wordvec

import (
    /* "fmt" */
    "os"
)

const RING_SIZE = 1000

// RingBuffer is a nice cached moving window data structure
type RingBuffer struct {
    Data [RING_SIZE]int
    Pk int
    WindowSize int
    Vocab *Vocab
    File *os.File
}

func (r RingBuffer) Current() int {
    return r.Data[r.Pk]
}

func (r RingBuffer) Context() []int {
    return append(r.Data[r.Pk-r.WindowSize:r.Pk], r.Data[r.Pk+1:r.Pk+r.WindowSize+1]...)
}

