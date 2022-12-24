package pathfragment

import (
	"testing"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	"github.com/google/go-cmp/cmp"
)

func TestResolver(t *testing.T) {
	for name, tc := range map[string]struct {
		fragments  []*anpb.PathFragment
		fragmentId uint32
		want       string
	}{
		"degenerate": {},
		"resolves": {
			fragments: []*anpb.PathFragment{
				{
					Id:       3,
					Label:    "java",
					ParentId: 2,
				},
				{
					Id:       2,
					Label:    "main",
					ParentId: 1,
				},
				{
					Id:    1,
					Label: "src",
				},
			},
			fragmentId: 3,
			want:       "src/main/java",
		},
	} {
		t.Run(name, func(t *testing.T) {
			resolver := NewResolver(tc.fragments)
			got, err := resolver.Resolve(tc.fragmentId)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("(-want,+got): %s", diff)
			}
		})
	}
}
