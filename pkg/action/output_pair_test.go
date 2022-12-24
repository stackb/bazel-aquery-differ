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
@@ -1,5 +1,5 @@
 Target: 
-ActionKey: abcdef
+ActionKey: 123456
 Mnemonic: 
 PrimaryOutput: 
 OutputFiles: 
`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := tc.pair.UnifiedDiff()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Error("(-want,+got): %s", diff)
			}
		})
	}
}
