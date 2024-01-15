# LinkedIn Scraper Chrome Extension

This Chrome Extension allows you to scrape LinkedIn profiles and display a visual alert based on the API response from a Golang backend server.

## Features

- Scrapes LinkedIn profiles on the LinkedIn people listing page.
- Calls a Golang API to get profile information for each person listed.
- Displays a red visual alert for a 404 response and a green one for a 200 response.

## Setup

### Golang Backend

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd linkedin-scraper

2. Install necessary Golang packages:

    ```
    go get github.com/PuerkitoBio/goquery
    go get github.com/gorilla/mux
    go get github.com/rs/cors
    ```

3. Run the Golang server:

    ```go run main.go```

# Chrome Extension

a. Open Chrome and go to ```chrome://extensions/```.

b. Enable "Developer mode" in the top-right corner.

c. Click "Load unpacked" and select the folder containing your extension files.

d. The extension icon should appear in the Chrome toolbar.

# Usage

a. Open a new tab and navigate to LinkedIn.

b. Click on the LinkedIn Scraper extension icon in the toolbar.

c. Click the "Scrape LinkedIn" button.

d. The extension will scrape LinkedIn profiles and display a visual alert based on the API response.

# Troubleshooting

a. If you encounter CORS issues, make sure your Golang server includes the necessary CORS headers. Check the Golang server logs and Chrome DevTools for more information.

b. For any other issues, refer to the error messages in the browser console (Chrome DevTools) and the server logs for debugging.

# License

This project is licensed under the MIT License - see the LICENSE file for details.