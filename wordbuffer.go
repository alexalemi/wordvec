package wordvec

import "bufio"

const BUFFER_SIZE = 1000

// ContextBuffer is a nice cached moving window data structure
type ContextBuffer struct {
	Data       [BUFFER_SIZE]int
	Pk         int
	WindowSize int
	Vocab      *Vocab
	Scanner    *bufio.Scanner
}

// Fill fills up the rest of the buffer
func (r *ContextBuffer) Fill(start int) error {
	length := len(r.Data)
	for i := start; i < length; i++ {

		ok := r.Scanner.Scan()
		if !ok {
			return r.Scanner.Err()
		}
		pk, ok := r.Vocab.Pk(r.Scanner.Text())

		r.Data[i] = r.Vocab.Pk(r.Scanner.Text())
	}
	return nil
}

// Incr increments the ContextBuffer, copying and refilling if necessary
func (r *ContextBuffer) Incr() (err error) {
	r.Pk++
	length := len(r.Data)
	if r.Pk+r.WindowSize > len(length) {
		// Need to copy and fill
		end := r.Pk + r.WindowSize - length
		copy(r.Data[:r.WindowSize+1+end], r.Data[r.Pk-r.WindowSize:r.Pk+end+1])
		err = r.Fill(r.Pk + end + 1)
	}
	return err
}

// Current returns the current integer
func (r ContextBuffer) Current() int {
	return r.Data[r.Pk]
}

// Context returns a slice of the given context of the given size
func (r ContextBuffer) Context(size int) []int {
	return append(r.Data[r.Pk-size:r.Pk], r.Data[r.Pk+1:r.Pk+size+1]...)
}
