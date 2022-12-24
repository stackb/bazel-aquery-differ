package depset

import (
	"testing"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	"github.com/google/go-cmp/cmp"
	"github.com/stackb/bazel-aquery-differ/pkg/artifact"
)

func TestPathMap(t *testing.T) {
	for name, tc := range map[string]struct {
		artifacts        []*anpb.Artifact
		fragments        []*anpb.PathFragment
		depSets          []*anpb.DepSetOfFiles
		depSetOfFilesIds []uint32
		want             []string
	}{
		"degenerate": {},
		"resolves": {
			depSets: []*anpb.DepSetOfFiles{
				{
					Id:                  1,
					DirectArtifactIds:   []uint32{1, 2}, // Foo.java,Bar.java
					TransitiveDepSetIds: []uint32{2},
				},
				{
					Id:                2,
					DirectArtifactIds: []uint32{3}, // Helper.java
				},
			},
			artifacts: []*anpb.Artifact{
				{
					Id:             1,
					PathFragmentId: 5,
				},
				{
					Id:             2,
					PathFragmentId: 4,
				},
				{
					Id:             3,
					PathFragmentId: 6,
				},
			},
			fragments: []*anpb.PathFragment{
				{
					Id:       6,
					Label:    "Helper.java",
					ParentId: 3,
				},
				{
					Id:       5,
					Label:    "Bar.java",
					ParentId: 3,
				},
				{
					Id:       4,
					Label:    "Foo.java",
					ParentId: 3,
				},
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
			depSetOfFilesIds: []uint32{1},
			want: []string{
				"src/main/java/Bar.java",
				"src/main/java/Foo.java",
				"src/main/java/Helper.java",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			paths, err := artifact.NewPathMap(tc.artifacts, tc.fragments)
			if err != nil {
				t.Fatal(err)
			}
			depSets := NewMap(tc.depSets)
			resolver := NewResolver(paths, depSets)
			got, err := resolver.ResolveIds(tc.depSetOfFilesIds)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("(-want,+got): %s", diff)
			}
		})
	}
}
