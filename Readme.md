# Ascii-Art-Web

## Description
`Ascii-Art-Web` is a Go-based web application that provides a graphical user interface (GUI) to generate ASCII art from text. It builds on the `ascii-art` project by offering a web-based platform to render ASCII art using different font banners. Users can input text and select a banner, and the application returns the ASCII representation.

---

## Objectives
- Serve a web interface to generate ASCII art.
- Allow users to select different banners:
  - **Standard**
  - **Shadow**
  - **Thinkertoy**
- Implement the following HTTP endpoints:
  - `GET /` - Displays the main HTML page.
  - `POST /ascii-art` - Accepts text and a banner type, then returns the ASCII art.
- Return appropriate HTTP status codes for each scenario:
  - `200 OK` for successful requests.
  - `400 Bad Request` for invalid input or method.
  - `404 Not Found` for missing resources (e.g., templates or banners).
  - `500 Internal Server Error` for unexpected issues.

---
## Project Structure
```
    ascii-art-web/
    ├── ascii-art-web.go        # Main application server
    ├── ascii-art.go            # ASCII art logic
    ├── ascii-art-web_test.go   # Test cases
    ├── go.mod
    ├── banners/
    │   ├── standard.txt
    │   ├── shadow.txt
    │   └── thinkertoy.txt
    ├── templates/
    │   ├── index.html
    │   ├── badRequest.html
    │   ├── internalServer.html
    │   └── notFound.html
    ├── static/
        ├── css/
            └── styles.css
```


## Usage
### How to Run
1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/your-repo/ascii-art-web.git
2. Navigate to the project directory:
    ```bash
    cd ascii-art-web
3. Ensure you have Go installed (version 1.19+ recommended).

4. Run the application:
    ```bash
    go run ascii-art-web.go

5. Open your browser and navigate to http://localhost:8080

---
## Instructions
1. Enter text into the input field on the homepage.
2. Select a banner using the radio buttons.
3. Click the "Generate" button to create ASCII art.
4. View the ASCII art output on the page.
5. For multi-line input, separate lines with \n.

---
## Implementation Details: Algorithm
### ASCII Art Generation (ascii-art.go)
The ASCII art rendering process involves the following steps:
#### 1. Input Validation:
* Check if the input string is valid.
* Validate characters to ensure they fall within the range of printable ASCII value (32–126).
* Split the input string by newline characters (\n) to handle multi-line input.
#### 2. Pattern Retrieval:

* Read the selected font banner file (e.g., standard.txt, shadow.txt, thinkertoy.txt) into memory.
* Convert the banner file's content into an array of strings, where each character's ASCII representation spans 9 lines.
#### 3. Rendering:

* For each line of the input text:
    * For each character in the line:
        * Compute the corresponding index in the banner file.
        * Append the corresponding ASCII art lines for that character.
* Handle special cases such as empty lines or invalid characters.
#### 3. Output Assembly:
* Combine all rendered ASCII art lines into a single output string.
#### 4. Error Handling:

* Return meaningful error messages for scenarios such as:
    * Missing banner files.
    * Invalid input characters.
    * File read errors.
### HTTP Server (ascii-art-web.go)
* The server is built using the Go net/http package.

* Routes:

    * GET /:
        * Renders the homepage using the index.html template.
        * Serves static CSS files for styling.
    * POST /ascii-art:
        * Parses form data (inputField and banner).
        * Calls the AsciiArt function to generate ASCII art.
        * Responds with the generated ASCII art or an appropriate error message.
* Error Handling:
    * Custom HTML templates (badRequest.html, notFound.html, internalServer.html) are served for HTTP error codes.
### Testing (ascii-art-web_test.go)
* Automated tests verify:
    * Correct status codes for different routes and scenarios.
    * Valid response generation for specific inputs and banners.
    * Proper error handling for invalid inputs and requests.

---
## Authors
- Mohammad mahdi Kheirkhah
- Fatemeh Kheirkhah
- Toft Diederichs
---
