package gather

// Reducer listens on the inbound channel and collects Units of processed work.
type Reducer struct {
	inbound chan *Unit
	sums    *Reduction
	closer  chan bool
}

// NewReducer allocates a reducer listening on the inbound channel.
func NewReducer(inbound chan *Unit) *Reducer {
	return &Reducer{
		inbound: inbound,
		closer: make(chan bool),
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
func (w *Reducer) Start() {
	go func() {
		for {
			select {
			case <-w.closer:
				return

			case p := <-w.inbound:
				w.count(w.sums, p)
			}
		}
	}()
}

// Close stops the Reducer and prevents it from starting again.
func (w *Reducer) Close() error {
	close(w.closer)
	return nil
}

// count adds the new unit to the sum.
func (w *Reducer) count(sums *Reduction, p *Unit) {
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