package wordvec

const (
    MAX_SENTENCE_LENGTH = 1000
    SIG_TABLE_SIZE = 1000
    MAX_SIG = 6
    MAX_CODE_SIZE = 40
)

type real float32

// ModelType marks what kind of model we are going to train
type ModelType int

const (
    CBow ModelType = iota
    SkipGram
)

// SampleType denotes what kind of sampling we are going to do
type SampleType int

const (
    HeirarchicalSoftmax SampleType = iota
    NegativeSampling 
)

// Word2Vec is the struct representing a Word2Vec training network
type Word2Vec struct {
    Size int
    Syn0 [][]real
    Syn1 [][]real
    Alpha0 real
    SampleThreshold real
    Iterations int
    Threads int
    Model ModelType
    Sample SampleType
    Vocab *Vocab
    UnigramTable []real
}

func (w *Word2Vec) Train(id int) {

}
