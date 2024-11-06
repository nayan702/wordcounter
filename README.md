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
The constants package includes configurable settings for managing rate limiting, request thresholds, and cooldown durations to optimize performance and stability. Below is a breakdown of each constant:

1. CooldownDurationSeconds: Specifies the cooldown duration (300 seconds) applied when the service encounters responses with status codes 999, 429, or 503, helping to manage retry intervals for temporary errors.

2. MaxConcurrentRequestsPerSec: Defines the maximum concurrent requests allowed per second by the rate limiter. Set to 15, this controls the request load and mitigates spikes.

3. RateLimiterBurstSize: Sets the burst size for rate limiting (5), allowing brief request bursts beyond the steady rate, which is useful for handling short demand spikes without exceeding limits.

4. SuccessThresholdForIncrease: Represents the threshold of consecutive successful requests (100) required to trigger an increase in the rate limit, supporting dynamic adjustments during stable operation.

5. RateLimiterAdjustmentPct: Specifies the rate adjustment percentage (10%) applied when successful request thresholds are met, enabling adaptive scaling based on recent performance.

These constants provide adaptive control over request flows, maintaining service stability with cooldowns and rate adjustments.

### Customizing Application Parameters
Additional parameters, such as the list of URLs to fetch, can be configured in main.go. To specify URLs, pass the list as a command-line argument.

#### Usage
Prepare URL and Word Bank Files:

Create urls.txt containing the URLs to fetch (one URL per line).
Create bank.txt containing valid words (one word per line).
Run the Application: Execute the following commands to process the URLs and display the word count:

For Testing with Fewer URLs: If you have a smaller set of URLs to process, use the following command to run the application with the testurls.txt file:
```bash
go run main.go testurls.txt
```
For Running with the Full List of URLs: To process the entire set of URLs in urls.txt, you can either use the following command or remove the parameter entirely:
```bash
go run main.go urls.txt
```
Output: The application will print the top counted words and their frequencies in a structured format.

