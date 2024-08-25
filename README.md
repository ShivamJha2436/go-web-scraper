# Simple Web Scraper

## Overview
The Simple Web Scraper is a command-line application built with Go that fetches HTML content from a specified URL and extracts text from it. This tool can be used to scrape and analyze web pages' textual content.

## Features
- Fetches HTML content from a URL.
- Extracts and prints all text nodes from the HTML document.

## Prerequisites
- Go 1.17 or later installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

## Installation
1. **Clone the Repository**:
   ```sh
   git clone https://github.com/yourusername/simple-web-scraper.git
   cd simple-web-scraper
2. **Fetch Dependencies**:
   ```sh
   go mod tidy

## Usage
1. **Build the Application**:
   ```sh
   go build -o web-scraper
2. **Run the application**:
   ```sh
   ./web-scraper -url https://example.com
  Replace https://example.com with the URL of the webpage you want to scrape.  
