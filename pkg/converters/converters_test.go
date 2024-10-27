package converters

import "testing"

func TestToSnakeCase(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"'Foo bar Baz' should be foo_bar_baz", "Foo bar Baz", "foo_bar_baz"},
		{"'Foo-bar Baz' should be foo_bar_baz", "Foo-bar Baz", "foo_bar_baz"},
		{"'Foo_bar Baz' should be foo_bar_baz", "Foo_bar Baz", "foo_bar_baz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ToSnakeCase(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"'Foo bar Baz' should be fooBarBaz", "Foo bar Baz", "fooBarBaz"},
		{"'Foo-bar Baz' should be fooBarBaz", "Foo-bar Baz", "fooBarBaz"},
		{"'foo_bar-Baz' should be fooBarBaz", "foo_bar-Baz", "fooBarBaz"},
		{"'foobar_baz' should be foobarBaz", "foobar_baz", "foobarBaz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ToCamelCase(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"'Foo bar Baz' should be FooBarBaz", "Foo bar Baz", "FooBarBaz"},
		{"'Foo-bar Baz' should be FooBarBaz", "Foo-bar Baz", "FooBarBaz"},
		{"'foo_bar-Baz' should be FooBarBaz", "foo_bar-Baz", "FooBarBaz"},
		{"'foobar_baz' should be FoobarBaz", "foobar_baz", "FoobarBaz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ToPascalCase(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestToTitleCase(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"'Foo bar Baz' should be Foo Bar Baz", "Foo bar Baz", "Foo Bar Baz"},
		{"'Foo-bar Baz' should be Foo Bar Baz", "Foo-bar Baz", "Foo Bar Baz"},
		{"'foo_bar-Baz' should be Foo Bar Baz", "foo_bar-Baz", "Foo Bar Baz"},
		{"'foobar_baz' should be Foobar Baz", "foobar_baz", "Foobar Baz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ToTitleCase(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
