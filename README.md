# WordCounter

WordCounter is a Go application that fetches content from a list of URLs, counts the top words based on a specified word bank, and handles rate limiting dynamically based on server responses. The application is designed to efficiently manage concurrent requests while respecting server limits.

## Features

- **Concurrent Fetching**: Fetches content from multiple URLs concurrently.
- **Dynamic Rate Limiting**: Adjusts request rates based on server responses (e.g., handling 999 response codes).
- **Word Counting**: Counts valid words based on a predefined word bank.
- **Structured Output**: Displays results in a structured format.
- **Error Handling**: Robust error handling for various HTTP response codes.

## Requirements

- Go version 1.15 or higher.
- Access to a list of URLs to fetch.
- A word bank file containing valid words.

## Project Structure
```
/wordcounter
    ├── main.go                # Main entry point for the application
    ├── handler
    │   ├── fetch_handler.go    # Handles fetching and processing content from URLs
    ├── parser
    │   ├── word_extractor.go    # Extracts words from the content
    ├── utils
    │   ├── word_utils.go        # Utilities for word counting and processing
    └── test
        ├── fetch_handler_test.go # Tests for fetching logic
        ├── word_extractor_test.go # Tests for word extraction logic
        └── word_utils_test.go     # Tests for utility functions
```
## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/nayan702/wordcounter.git
   cd wordcounter

2. Ensure Go is Installed: Verify that you have Go installed on your system by running:
    ```bash
    go version

3. Download Dependencies: Ensure that any required dependencies are downloaded by running:
    ```bash
    go mod tidy

## Configuration
You can configure the application by modifying the parameters in main.go, such as:
The list of URLs to fetch (passed as a command-line argument).

## Usage
Prepare Your URL and Word Bank Files: Create a file named urls.txt containing the URLs to fetch (one per line), and another file named bank.txt containing valid words (one per line).

Run the Application: Execute the following command:

```bash
go run main.go testurls.txt
go run main.go urls.txt
Output: The application will print the top counted words and their frequencies in a structured format.
```

