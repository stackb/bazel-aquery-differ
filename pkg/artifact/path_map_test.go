package artifact

import (
	"testing"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	"github.com/google/go-cmp/cmp"
)

func TestPathMap(t *testing.T) {
	for name, tc := range map[string]struct {
		artifacts  []*anpb.Artifact
		fragments  []*anpb.PathFragment
		artifactId uint32
		want       string
	}{
		"degenerate": {},
		"resolves": {
			artifacts: []*anpb.Artifact{
				{
					Id:             1,
					PathFragmentId: 3,
					IsTreeArtifact: true,
				},
			},
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
			artifactId: 1,
			want:       "src/main/java",
		},
	} {
		t.Run(name, func(t *testing.T) {
			paths, err := NewPathMap(tc.artifacts, tc.fragments)
			if err != nil {
				t.Fatal(err)
			}
			got := paths[tc.artifactId]
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Error("(-want,+got): %s", diff)
			}
		})
	}
}
