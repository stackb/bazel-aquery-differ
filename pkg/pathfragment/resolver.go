package pathfragment

import (
	"fmt"
	"path/filepath"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
)

// Resolver is a helper type that reconstructs full paths based on path IDs.
type Resolver struct {
	fragments map[uint32]*anpb.PathFragment
}

func NewResolver(fragments []*anpb.PathFragment) *Resolver {
	resolver := &Resolver{
		fragments: make(map[uint32]*anpb.PathFragment),
	}
	for _, v := range fragments {
		resolver.fragments[v.Id] = v
	}
	return resolver
}

func (r *Resolver) Resolve(id uint32) (string, error) {
	var tokens []string
	current := id
	for current != 0 {
		if fragment, ok := r.fragments[current]; ok {
			tokens = append(tokens, fragment.Label)
			current = fragment.ParentId
		} else {
			return "", fmt.Errorf("path fragment not found: %d", current)
		}
	}
	reverse(tokens)
	return filepath.Join(tokens...), nil
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
