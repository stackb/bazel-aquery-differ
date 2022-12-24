package report

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRenderIndexHtml(t *testing.T) {
	for name, tc := range map[string]struct {
		beforeFile string
		afterFile  string
		want       string
	}{
		"degenerate": {
			want: `<html>

<head>
</head>

<body>
    <h1>aquerydiff report v1</h1>
    <ul>
        <li>before file: <code></code></li>
        <li>after file: <code></code></li>
    </ul>
</body>

</html>`,
		},
		"typical": {
			beforeFile: "a.json",
			afterFile:  "b.json",
			want: `<html>

<head>
</head>

<body>
    <h1>aquerydiff report v1</h1>
    <ul>
        <li>before file: <code>a.json</code></li>
        <li>after file: <code>b.json</code></li>
    </ul>
</body>

</html>`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			var out strings.Builder
			if err := renderIndexHtml(&out, tc.beforeFile, tc.afterFile); err != nil {
				t.Fatal(err)
			}
			got := out.String()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("(-want,+got): %s", diff)
			}
		})
	}
}
