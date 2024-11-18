package main

import (
	"errors"
	"html"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGetHandler consolidates multiple GET tests using a table-driven structure.
func TestGetHandler(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		method         string
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
		// {
		// 	name:           "Internal Server Error",
		// 	url:            "/",
		// 	method:         http.MethodGet,
		// 	expectedStatus: http.StatusInternalServerError,
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, test.url, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetHandler)

			handler.ServeHTTP(rr, req)

			if rr.Code != test.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", test.expectedStatus, rr.Code)
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
			name:           "Valid Input",
			url:            "/ascii-art",
			method:         http.MethodPost,
			formData:       "inputField={123}\\n<Hello> (World)!&banner=standard.txt",
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
			formData:       "inputField=%24%25%20%22%3D&banner=shadow.txt",
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
			formData:       "inputField=Hello\\n\\nthere&banner=standard.txt",
			expectedStatus: http.StatusOK,
			expectedDivID:  "response",
			expectedDivVal: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

 _     _                           
| |   | |                          
| |_  | |__     ___   _ __    ___  
| __| |  _ \   / _ \ | '__|  / _ \ 
\ |_  | | | | |  __/ | |    |  __/ 
 \__| |_| |_|  \___| |_|     \___| 
                                   
                                   
`,
		},
		{
			name:           "Bad request",
			url:            "/ascii-art",
			method:         http.MethodPost,
			formData:       "inputField=Helloä&banner=standard.txt",
			expectedStatus: http.StatusBadRequest,
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var req *http.Request
			var err error

			if test.method == http.MethodPost {
				req, err = http.NewRequest(test.method, test.url, strings.NewReader(test.formData))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			} else {
				req, err = http.NewRequest(test.method, test.url, nil)
			}

			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(PostHandler)

			handler.ServeHTTP(rr, req)

			if rr.Code != test.expectedStatus {
				t.Errorf("Expected status code %d\n, but got\n%d", test.expectedStatus, rr.Code)
				//return
			}

			// Parse the HTML response
			divVal, err := extractDivValueByID(rr.Body.String(), test.expectedDivID)
			decodedDivVal := html.UnescapeString(divVal)
			if err != nil && strings.HasPrefix(test.name, "Valid Input") {
				t.Fatalf("Error extracting div: %v", err)
			}

			if decodedDivVal != test.expectedDivVal {
				t.Errorf("Expected div #%s to have value %q, but got %q", test.expectedDivID, test.expectedDivVal, divVal)
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