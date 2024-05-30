# City Parser

This project is a Golang parser that processes HTML documents named `[city].html`, parses the streets included in these files, and returns JSON files organized by cities. The program can handle multiple HTML files simultaneously and includes functionality to parse a single HTML file and return a single JSON file named for its city.

## Features

- Parse multiple HTML files and generate JSON files for each city.
- Parse a single HTML file and return a JSON file named for its city.

## Project Structure

- `main.go`: Entry point of the application.
- `parser/`: Contains the parsing logic.
    - `parser.go`: Implements the parsing functionality.
    - `parser_test.go`: Contains tests for the parser.
- `data/`: Directory to place the HTML files to be parsed.

## Usage

1. **Parse Multiple HTML Files:**

   Place your HTML files in the `data/` directory and run:

   ```sh
   go run main.go
