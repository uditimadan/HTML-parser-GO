**HTML Cleaner**

This Go project provides two primary functionalities: extracting words and links (HREFs) from HTML content and cleaning up URLs, especially handling relative and absolute paths.

**Author**

Uditi Madan

**Features**

1. Extract Functionality:
    * Extracts visible words and HREF links from the provided HTML.
    * Useful for parsing HTML content to retrieve text and hyperlinks.
2. Clean Functionality:
    * Cleans and resolves relative URLs into absolute URLs based on a given host.
    * Useful for normalizing URLs before further processing.

**Project Structure**

* `extract.go`: Contains the extract function, which parses HTML content to retrieve words and links.
* `clean.go`: Contains the clean function, which resolves relative paths in URLs.
* `main_test.go`: Provides unit tests for the extract and clean functions using table-driven tests.
* `go.mod`: Manages the project's module dependencies.
* `go.sum`: Contains checksums of the module dependencies.

**Usage**

Extracting Words and Links
The extract function takes in HTML content as a byte slice and returns two slices:
* A slice of words found in the visible HTML content.
* A slice of HREF links found in the anchor (<a>) tags.

*Example usage:
html := []byte(`<body><p>Hello World</p><a href="http://example.com">Example Link</a></body>`)
words, hrefs := extract(html)
fmt.Println(words)  // Output: ["Hello", "World", "Example", "Link"]
fmt.Println(hrefs)  // Output: ["http://example.com"]*

**Cleaning URLs**

The clean function takes in a host and an HREF link and returns the absolute URL. It can handle both relative and absolute URLs.

*Example usage:
host := "https://example.com"
href := "/path/to/resource"
cleanedURL := clean(host, href)
fmt.Println(cleanedURL)  // Output: https://example.com/path/to/resource*

**Running Tests**

The project includes comprehensive unit tests. To run the tests, simply use:
*go test*
The tests validate different edge cases for both the extract and clean functions in main_test.go file

**`TestExtract` Function Test Cases:**
1. **simple** - Extracts words and one link from basic HTML.
2. **cs272** - Extracts words and a link from HTML with headings and paragraphs.
3. **sample_from_specs** - Extracts words and a link with combined text and numbers.
4. **multiple_links** - Extracts words and multiple links from HTML with two anchor tags.

**`TestCleanHref` Function Test Cases:**
1. **relative** - Resolves a relative URL using the provided host.
2. **absolute** - Verifies an absolute URL remains unchanged.
3. **empty_path** - Handles empty path URLs, returning the base host.
4. **multiple_paths** - Resolves multiple relative paths into complete URLs.

**License**

This project is licensed under the MIT License
