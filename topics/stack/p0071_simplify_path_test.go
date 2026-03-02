package stack

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/home/", "/home"},
		{"/home//foo/", "/home/foo"},
		{"/home/user/Documents/../Pictures", "/home/user/Pictures"},
		{"/../", "/"},
		{"/.../a/../b/c/../d/./", "/.../b/d"},
		{"/a/./b/../../c/", "/c"},
		{"/", "/"},
	}
	for _, tt := range tests {
		got := SimplifyPath(tt.path)
		if got != tt.want {
			t.Errorf("SimplifyPath(%q) = %q, want %q", tt.path, got, tt.want)
		}
	}
}
