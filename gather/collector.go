package gather

type Reducer struct {
	inbound chan *Unit
	sums    *Reduction
}

func NewReducer(inbound chan *Unit) *Reducer {
	return &Reducer{
		inbound: inbound,
		sums: &Reduction{
			ExtCount: make(map[string]int),
			Units:    make([]*Unit, 0),
		},
	}
}

func (w *Reducer) Start() chan bool {
	sums := w.sums
	finishedReading := make(chan bool)
	wrapItUp := false

	go func() {
		for {
			select {

			case wrapItUp = <-finishedReading:

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

			default:
				if wrapItUp {
					return
				}
			}
		}
	}()

	return finishedReading
}
