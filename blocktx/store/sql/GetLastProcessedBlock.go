package sql

import (
	"github.com/TAAL-GmbH/arc/blocktx/blocktx_api"

	"context"
)

func (s *SQL) GetLastProcessedBlock(ctx context.Context) (*blocktx_api.Block, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	q := `
		SELECT
		 b.hash
		,b.prevhash
		,b.merkleroot
		,b.height
		,b.orphanedyn
		FROM blocks b
		WHERE b.processedyn = true
		ORDER BY b.height DESC
		LIMIT 1
	`

	block := &blocktx_api.Block{}

	if err := s.db.QueryRowContext(ctx, q).Scan(&block.Hash, &block.Prevhash, &block.Merkleroot, &block.Height, &block.Orphaned); err != nil {
		return nil, err
	}

	return block, nil
}
