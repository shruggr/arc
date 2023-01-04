package p2p

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	"github.com/TAAL-GmbH/arc/p2p/bsvutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLittleEndian(t *testing.T) {
	le := binary.LittleEndian.Uint32([]byte{0x50, 0xcc, 0x0b, 0x00})

	require.Equal(t, uint32(773200), le)
}

func TestExtractHeight(t *testing.T) {
	coinbase, _ := hex.DecodeString("01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff570350cc0b041547b5630cfabe6d6d0000000000000000000000000000000000000000000000000000000000000000010000000000000047ed20542096bd0000000000143362663865373833636662643732306431383436000000000140be4025000000001976a914c9b0abe09b7dd8e9d1e8c1e3502d32ab0d7119e488ac00000000")
	tx, err := bsvutil.NewTxFromBytes(coinbase)
	require.NoError(t, err)

	height := extractHeightFromCoinbaseTx(tx.MsgTx())

	assert.Equalf(t, uint64(773200), height, "height should be 773200, got %d", height)
}

func TestExtractHeightForRegtest(t *testing.T) {
	coinbase, _ := hex.DecodeString("02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0502dc070101ffffffff012f500900000000002321032efe256e14fd77eea05d0453374f8920e0a7a4a573bb3937ef3f567f3937129cac00000000")
	tx, err := bsvutil.NewTxFromBytes(coinbase)
	require.NoError(t, err)

	height := extractHeightFromCoinbaseTx(tx.MsgTx())

	assert.Equalf(t, uint64(2012), height, "height should be 2012, got %d", height)
}
