package v1

import (
	"crypto/rand"
	"flag"
	"testing"
	"time"

	"github.com/tendermint/tendermint/types"
)

var txLimit = flag.Int("num-txn", 1, "Number of transactions")

func TestMempoolAddRemove(t *testing.T) {
	txmp := setup(t, *txLimit)
	txch := make(chan *WrappedTx, 10)

	go func() {
		defer close(txch)

		b := make([]byte, 10000)
		for i := 0; i < *txLimit; i++ {
			rand.Read(b)
			tx := types.Tx(b)
			wtx := &WrappedTx{
				tx:        tx,
				hash:      tx.Key(),
				timestamp: time.Now().UTC(),
				height:    txmp.height,
			}
			txmp.insertTx(wtx)
			txch <- wtx
		}
	}()
	for tx := range txch {
		txmp.removeTx(tx, true)
	}
}
