package v0

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	abciclient "github.com/tendermint/tendermint/abci/client"
	"github.com/tendermint/tendermint/abci/example/kvstore"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/internal/mempool"
)

var (
	txLimit = flag.Int("num-txn", 1, "Number of transactions")
	ptrLog  = flag.String("logfile", "pointer.log", "Pointer log")
)

func TestMempoolAddRemove(t *testing.T) {
	app := kvstore.NewApplication()
	cc := abciclient.NewLocalCreator(app)
	mp, cleanup, err := newMempoolWithApp(cc)
	if err != nil {
		t.Fatalf("Setup: %v", err)
	}
	defer cleanup()

	for i := 0; i < *txLimit; i++ {
		checkTxs(t, mp, tt.numTxsToCreate, mempool.UnknownPeerID)
		got := mp.ReapMaxBytesMaxGas(tt.maxBytes, tt.maxGas)
		assert.Equal(t, tt.expectedNumTxs, len(got), "Got %d txs, expected %d, tc #%d",
			len(got), tt.expectedNumTxs, tcIndex)
		mp.Flush()
}
