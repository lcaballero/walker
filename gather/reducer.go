package gather

// Reducer listens on the inbound channel and collects Units of processed work.
type Reducer struct {
	inbound chan *Unit
	sums    *Reduction
}

// NewReducer allocates a reducer listening on the inbound channel.
func NewReducer(inbound chan *Unit) *Reducer {
	return &Reducer{
		inbound: inbound,
		sums: &Reduction{
			IndexInfo: &IndexInfo{
				ExtCount: make(map[string]int),
			},
			Units: make([]*Unit, 0),
		},
	}
}

// Start returns a finished signal channel which indicates to the reducer
// that ALL sub-jobs finished their tasks.
func (w *Reducer) Start() chan bool {
	sums := w.sums
	finishedReading := make(chan bool)
	complete := make(chan bool, 1)

	go func() {
		for {
			select {

			case <-complete:
				close(finishedReading)
				close(complete)
				return

			case wrapItUp := <-finishedReading:
				complete <- wrapItUp

			case p := <-w.inbound:
				sums.Units = append(sums.Units, p)
				n, ok := sums.ExtCount[p.Ext]
				if !ok {
					sums.ExtCount[p.Ext] = 1
				} else {
					sums.ExtCount[p.Ext] = n + 1
				}
				if p.IsSkipped {
					sums.SkippedReading++
				}
				if p.IsDir {
					sums.DirsFound++
				} else {
					sums.FilePathsCollected++
				}
			}
		}
	}()

	return finishedReading
}
