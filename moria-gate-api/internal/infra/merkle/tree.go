package merkle

import (
	"crypto/sha256"
	"encoding/hex"
)

// MerkleTree is a simple binary Merkle tree
type MerkleTree struct {
	Leaves []string
	Root   string
}

func NewMerkleTree(leaves []string) MerkleTree {
	hashed := make([]string, len(leaves))
	for i, l := range leaves {
		h := sha256.Sum256([]byte(l))
		hashed[i] = hex.EncodeToString(h[:])
	}
	return MerkleTree{
		Leaves: hashed,
		Root:   buildRoot(hashed),
	}
}

func buildRoot(leaves []string) string {
	if len(leaves) == 0 {
		return ""
	}
	if len(leaves) == 1 {
		return leaves[0]
	}

	var newLevel []string
	for i := 0; i < len(leaves); i += 2 {
		if i+1 < len(leaves) {
			combined := leaves[i] + leaves[i+1]
			sum := sha256.Sum256([]byte(combined))
			newLevel = append(newLevel, hex.EncodeToString(sum[:]))
		} else {
			newLevel = append(newLevel, leaves[i])
		}
	}
	return buildRoot(newLevel)
}
