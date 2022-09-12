package rootmulti

import (
	"github.com/okex/exchain/libs/iavl"
	dbm "github.com/okex/exchain/libs/tm-db"
	"sync"
)

type (
	MetadataItem struct {
		version      int64
		cInfo        commitInfo
		pruneHeights []int64
		versions     []int64
	}

	Metadata struct {
		db       dbm.DB
		task     chan MetadataItem
		quit     chan struct{}
		mut      sync.Mutex
		isClosed bool
	}
)

func NewMetadata(db dbm.DB) *Metadata {
	mt := &Metadata{
		db:       db,
		task:     make(chan MetadataItem, iavl.CommitGapHeight),
		quit:     make(chan struct{}),
		isClosed: false,
	}
	if iavl.EnableAsyncCommit {
		go func() {
			for {
				select {
				case mc, ok := <-mt.task:
					if ok {
						flushMetadata(mt.db, mc.version, mc.cInfo, mc.pruneHeights, mc.versions)
					} else {
						mt.quit <- struct{}{}
					}
				}
			}
		}()
	}
	return mt
}

func (mt *Metadata) notifyFlushMetadata(version int64, cInfo commitInfo, pruneHeights []int64, versions []int64) {
	mt.mut.Lock()
	defer mt.mut.Unlock()

	if !mt.isClosed {
		mt.task <- MetadataItem{
			version, cInfo, pruneHeights, versions,
		}
	}
}

func (mt *Metadata) GracefulExit() {
	if iavl.EnableAsyncCommit {
		mt.mut.Lock()
		defer mt.mut.Unlock()
		if !mt.isClosed {
			mt.isClosed = true
			close(mt.task)
			<-mt.quit
		}
	}
}
