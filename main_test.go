package main

import (
	"reflect"
	"testing"
)

//to test the extract function for various test cases
func TestExtract(t *testing.T) {
	tests := []struct {
		name      string
		html      string
		wantWords []string
		wantHrefs []string
	}{
		{
			name:      "simple",
			html:      `<body><p>Some text here</p><a href="http://example.com">Example</a></body>`,
			wantWords: []string{"Some", "text", "here", "Example"},
			wantHrefs: []string{"http://example.com"},
		},
		{
			name:      "cs272",
			html:      `<!DOCTYPE html><html><head><title>CS272 | Welcome</title></head><body><p>Hello World!</p><p>Welcome to <a href="https://cs272-f24.github.io/">CS272</a>!</p></body></html>`,
			wantWords: []string{"CS272", "|", "Welcome", "Hello", "World!", "Welcome", "to", "CS272", "!"},
			wantHrefs: []string{"https://cs272-f24.github.io/"},
		},
		{
			name:      "sample_from_specs",
			html:      `<body><p>Hello World</p><a href="http://world.com">Example 123</a></body>`,
			wantWords: []string{"Hello", "World", "Example", "123"},
			wantHrefs: []string{"http://world.com"},
		},
		{
			name:      "multiple_links",
			html:      `<body><a href="http://link1.com">Link1</a><a href="http://link2.com">Link2</a></body>`,
			wantWords: []string{"Link1", "Link2"},
			wantHrefs: []string{"http://link1.com", "http://link2.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWords, gotHrefs := extract([]byte(tt.html))
            //DeepEqual func usage to check whether Words and refs match or not
			if !reflect.DeepEqual(gotWords, tt.wantWords) {
				t.Errorf("extract() gotWords = %v, want %v", gotWords, tt.wantWords)
			}
			if !reflect.DeepEqual(gotHrefs, tt.wantHrefs) {
				t.Errorf("extract() gotHrefs = %v, want %v", gotHrefs, tt.wantHrefs)
			}
		})
	}
}

//to test the clean function for various cases
func TestCleanHref(t *testing.T) {
	tests := []struct {
		name string
		host string
		href string
		want string
	}{
		{
			name: "relative",
			host: "https://example.com",
			href: "/path",
			want: "https://example.com/path",
		},
		{
			name: "absolute",
			host: "https://example.com",
			href: "https://other.com/path",
			want: "https://other.com/path",
		},
		{
			name: "empty_path",
			host: "https://example.com",
			href: "/",
			want: "https://example.com/",
		},
	}
	//running and checking for diff above test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clean(tt.host, tt.href); got != tt.want {
				t.Errorf("clean() = %v, want %v", got, tt.want)
			}
		})
	}
    //test case example from the spec
	host := "https://cs272-f24.github.io/"
	hrefs := []string{"/", "/help/", "/syllabus/", "https://gobyexample.com/"}
	expected := []string{
		"https://cs272-f24.github.io/",
		"https://cs272-f24.github.io/help/",
		"https://cs272-f24.github.io/syllabus/",
		"https://gobyexample.com/",
	}
	var cleaned []string
	for _, href := range hrefs {
		cleaned = append(cleaned, clean(host, href))
	}
	if !reflect.DeepEqual(cleaned, expected) {
		t.Errorf("Expected %v, got %v", expected, cleaned)
	}
}
