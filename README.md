# Simple Web Scraper

## Overview The Simple Web Scraper is a command-line application built with Go that fetches HTML content from a specified URL and extracts text from it. This tool can be used to scrape and analyze web pages' textual content.

## Features - Fetches HTML content from a URL. - Extracts and prints all text nodes from the HTML document.

## Prerequisites - Go 1.17 or later installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

## Installation 1. **Clone the Repository**: ```sh git clone https://github.com/yourusername/simple-web-scraper.git cd simple-web-scraper ``` 2. **Fetch Dependencies**: ```sh go mod tidy ```

## Usage 1. **Build the Application**: ```sh go build -o web-scraper ``` 2. **Run the Application**: ```sh ./web-scraper -url https://example.com ``` Replace `https://example.com` with the URL of the webpage you want to scrape. **Command-Line Arguments**: - `-url`: The URL of the webpage to scrape. This argument is required. Example: ```sh ./web-scraper -url https://www.example.com ``` The application will output the extracted text content from the specified URL.

## Code Overview - **`main.go`**: Contains the core logic for fetching HTML content from a URL, parsing it, and extracting text nodes. - **`fetchHTML(url string) (*html.Node, error)`**: Fetches and parses the HTML content from the given URL. - **`extractText(node *html.Node)`**: Recursively extracts and prints text from HTML nodes. - **`main()`**: The entry point of the application. Handles command-line arguments and invokes the scraping functions. - **`main_test.go`**: Contains basic unit tests for the `extractText` function to ensure it works correctly.

## Testing To run the tests: ```sh go test ``` This will execute the unit tests defined in `main_test.go` and verify the functionality of the `extractText` function.

## Contributing Contributions are welcome! If you have suggestions or improvements, feel free to create a pull request or open an issue.

## License This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact For any questions or feedback, please contact [your email](mailto:your.email@example.com).
