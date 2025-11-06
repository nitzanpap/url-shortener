package urls

import "testing"

func TestIsUrlValid(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "valid url",
			input: "https://leadit-media.com",
			want:  true,
		},
		{
			name:  "valid url with path",
			input: "https://management-platform.zenner-il.com/platform",
			want:  true,
		},
		{
			name:  "invalid scheme",
			input: "ftp://example.com/resource",
			want:  false,
		},
		{
			name:  "missing scheme",
			input: "example.com",
			want:  false,
		},
		{
			name:  "empty string",
			input: "",
			want:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isUrlValid(tc.input); got != tc.want {
				t.Fatalf("isUrlValid(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
