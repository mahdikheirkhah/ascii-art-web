package main

import (
	"errors"
	"html"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// TestGetHandler consolidates multiple GET tests using a table-driven structure.
func TestGetHandler(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		method         string
		body           string
		expectedStatus int
	}{
		{
			name:           "Valid Root Path",
			url:            "/",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Non-existent Path",
			url:            "/non-existent",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "bad Request",
			url:            "/",
			method:         http.MethodPost,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, test.url, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			responseHolder := httptest.NewRecorder()

			handler := http.HandlerFunc(GetHandler)
			handler.ServeHTTP(responseHolder, req)

			if responseHolder.Code != test.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", test.expectedStatus, responseHolder.Code)
			}
		})
	}
}

// TestPostHandler consolidates multiple POST tests using a table-driven structure.
func TestPostHandler(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		method         string
		formData       string
		expectedStatus int
		expectedDivID  string
		expectedDivVal string
	}{
		{
			name:   "Valid Input",
			url:    "/ascii-art",
			method: http.MethodPost,
			formData: `inputField={123}
<Hello> (World)!&banner=standard.txt`,
			expectedStatus: http.StatusOK,
			expectedDivID:  "response",
			expectedDivVal: `   __                     __    
  / /  _   ____    _____  \ \   
 | |  / | |___ \  |___ /   | |  
/ /   | |   __) |   |_ \    \ \ 
\ \   | |  / __/   ___) |   / / 
 | |  |_| |_____| |____/   | |  
  \_\                     /_/   
                                
   __  _    _          _   _          __            __ __          __                 _       _  __    _  
  / / | |  | |        | | | |         \ \          / / \ \        / /                | |     | | \ \  | | 
 / /  | |__| |   ___  | | | |   ___    \ \        | |   \ \  /\  / /    ___    _ __  | |   __| |  | | | | 
< <   |  __  |  / _ \ | | | |  / _ \    > >       | |    \ \/  \/ /    / _ \  | '__| | |  / _` + "`" + ` |  | | | | 
 \ \  | |  | | |  __/ | | | | | (_) |  / /        | |     \  /\  /    | (_) | | |    | | | (_| |  | | |_| 
  \_\ |_|  |_|  \___| |_| |_|  \___/  /_/         | |      \/  \/      \___/  |_|    |_|  \__,_|  | | (_) 
                                                   \_\                                           /_/      
                                                                                                          
`,
		},
		{
			name:           "Valid Input2",
			url:            "/ascii-art",
			method:         http.MethodPost,
			formData:       "inputField=123??&banner=standard.txt",
			expectedStatus: http.StatusOK,
			expectedDivID:  "response",
			expectedDivVal: `                     ___    ___   
 _   ____    _____  |__ \  |__ \  
/ | |___ \  |___ /     ) |    ) | 
| |   __) |   |_ \    / /    / /  
| |  / __/   ___) |  |_|    |_|   
|_| |_____| |____/   (_)    (_)   
                                  
                                  
`,
		},
		{
			name:           "Valid Input3",
			url:            "/ascii-art",
			method:         http.MethodPost,
			formData:       "inputField=%24%25%20%22%3D&banner=shadow.txt", //$% "=
			expectedStatus: http.StatusOK,
			expectedDivID:  "response",
			expectedDivVal: `                        _|  _|            
  _|   _|_|    _|       _|  _|            
_|_|_| _|_|  _|                _|_|_|_|_| 
_|_|       _|                             
  _|_|   _|  _|_|              _|_|_|_|_| 
_|_|_| _|    _|_|                         
  _|                                      
                                          
`,
		},
		{
			name:           "Valid Input4",
			url:            "/ascii-art",
			method:         http.MethodPost,
			formData:       "inputField=123 T/fs#R&banner=thinkertoy.txt",
			expectedStatus: http.StatusOK,
			expectedDivID:  "response",
			expectedDivVal: `                                                       
  0    --  o-o        o-O-o     o  o-o      | |  o--o  
 /|   o  o    |         |      /   |       -O-O- |   | 
o |     /   oo          |     o   -O-  o-o  | |  O-Oo  
  |    /      |         |    /     |    \  -O-O- |  \  
o-o-o o--o o-o          o   o      o   o-o  | |  o   o 
                                                       
                                                       
`,
		},
		{
			name:           "Valid Input 5",
			url:            "/ascii-art",
			method:         http.MethodPost,
			formData:       "inputField=Helloäö&banner=standard.txt",
			expectedStatus: http.StatusOK,
			expectedDivID:  "response",
			expectedDivVal: `Invalid characters are:
ä ö `,
		},
		{
			name:           "Bad request2",
			url:            "/ascii-art",
			method:         http.MethodGet,
			formData:       "inputField=Hello&banner=standard.txt",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Not Found",
			url:            "/ascii-art/jjdjdj",
			method:         http.MethodPost,
			formData:       "inputField=Hello&banner=standard.txt",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Not Found 2",
			url:            "/ascii-art",
			method:         http.MethodPost,
			formData:       "inputField=Hello&banner=stand.txt",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Internal Server Error",
			url:            "/ascii-art",
			method:         http.MethodPost,
			formData:       "inputField=Hello&banner=standard.txt",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var req *http.Request
			var err error

			if test.name == "Valid Input" { //first test file which we have to add \r to \n
				test.formData = strings.ReplaceAll(test.formData, "\n", "\r\n")
			} else if test.name == "Internal Server Error" { //cahnge the name of the file to have internal server Error
				sourceName := "./templates/index.html"
				destinationName := "./templates/home.html"
				err := os.Rename(sourceName, destinationName)
				if err != nil {
					t.Errorf("Error renaming file: %v\n", err)
					return
				}
			}

			if test.method == http.MethodPost {
				req, err = http.NewRequest(test.method, test.url, strings.NewReader(test.formData))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded") //Content-Type: Specifies the media type (MIME type) of the request body.
				//application/x-www-form-urlencoded indicates that the body contains form data encoded as key-value pairs(e.g., key1=value1&key2=value2).
			} else {
				req, err = http.NewRequest(test.method, test.url, nil)
			}

			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			responseHolder := httptest.NewRecorder()

			handler := http.HandlerFunc(PostHandler)
			handler.ServeHTTP(responseHolder, req)

			if responseHolder.Code != test.expectedStatus {
				t.Errorf("Expected status code %d\n, but got\n%d", test.expectedStatus, responseHolder.Code)
				//return
			}

			// Parse the HTML response
			if strings.HasPrefix(test.name, "Valid Input") { //only for the test cases with name <<valid input>> check for the output
				divVal, err := extractDivValueByID(responseHolder.Body.String(), test.expectedDivID)
				decodedDivVal := html.UnescapeString(divVal) //check for unescape chars : < becomes &lt; > becomes &gt; & becomes &amp; " becomes &quot
				if err != nil {
					t.Errorf("Error extracting div: %v", err)
				}
				if decodedDivVal != test.expectedDivVal {
					t.Errorf("Expected div #%s to have value %q, but got %q", test.expectedDivID, test.expectedDivVal, divVal)
				}
			}
			if test.name == "Internal Server Error" { //recahnge the name of the file
				sourceName := "./templates/home.html"
				destinationName := "./templates/index.html"
				err := os.Rename(sourceName, destinationName)
				if err != nil {
					t.Errorf("Error renaming file: %v\n", err)
					return
				}
			}
		})
	}
}

// extractDivValueByID parses the HTML string and extracts the value of a <div> with a specific id.
func extractDivValueByID(html, divID string) (string, error) {
	startTag := `<div id="` + divID + `">`
	endTag := `</div>`

	// Find the start index of the desired <div>
	startIndex := strings.Index(html, startTag)
	if startIndex == -1 {
		return "", errors.New("div with id '" + divID + "' not found")
	}
	startIndex += len(startTag)

	// Find the end index of the </div>
	endIndex := strings.Index(html[startIndex:], endTag)
	if endIndex == -1 {
		return "", errors.New("closing tag for div with id '" + divID + "' not found")
	}

	// Extract and return the content
	return html[startIndex : startIndex+endIndex], nil
}
