package action

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	dipb "github.com/stackb/bazel-aquery-differ/build/stack/bazel/aquery/differ"
)

func TestOutputPairUnifiedDiff(t *testing.T) {
	for name, tc := range map[string]struct {
		pair OutputPair
		want string
	}{
		"degenerate": {},
		"example": {
			pair: OutputPair{
				Output: "src/main/java/libhelper.jar",
				Before: &dipb.Action{
					ActionKey: "abcdef",
				},
				After: &dipb.Action{
					ActionKey: "123456",
				},
			},
			want: `--- src/main/java/libhelper.jar
+++ src/main/java/libhelper.jar
@@ -1,2 +1,2 @@
-action_key:  "abcdef"
+action_key:  "123456"
 
`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := tc.pair.UnifiedDiff()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("(-want,+got): %s", diff)
			}
		})
	}
}
