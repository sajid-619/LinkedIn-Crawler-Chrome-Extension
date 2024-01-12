# LinkedIn Scraper Chrome Extension

This project is a simple LinkedIn scraper implemented as a Chrome Extension using Golang, Goquery, and MongoDB. It allows you to scrape LinkedIn profiles and store the information in a MongoDB database.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

Before you begin, ensure you have the following installed:

- [Golang](https://golang.org/doc/install)
- [MongoDB](https://docs.mongodb.com/manual/installation/)
- [Google Chrome Browser](https://www.google.com/chrome/)

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/your-username/linkedin-scraper-chrome-extension.git

   ```

2. Install the necessary Golang packages:

```
go get github.com/go-chi/chi
go get github.com/PuerkitoBio/goquery
go get gopkg.in/mgo.v2

```

3. Set up MongoDB and create a database named `linkedin_scraper`

4. Modify the MongoDB connection string in `main.go` based on your setup:

```

session, err := mgo.Dial("mongodb://localhost:27017/linkedin_scraper")

```

5. Run the Golang server:

```
go run main.go

```

6. Load the Chrome Extension:

    - Open Chrome and navigate to chrome://extensions/.
    - Enable "Developer mode."
    - Click "Load unpacked" and select the folder containing your extension files.

# Usage

1. Open a new tab and navigate to a LinkedIn profile (e.g., https://www.linkedin.com/in/username).
2. Click on the extension icon in the Chrome toolbar.
3. Click the "Scrape Profile" button to scrape the LinkedIn profile.
4. The scraped information will be displayed in the extension popup.

# Project Setup

`main.go`: Golang server implementation using Chi router, Goquery for scraping, and MongoDB for data storage.
`manifest.json`: Chrome Extension manifest file.
`popup.html`: HTML file for the extension popup.
`popup.js`: JavaScript file for handling extension popup behavior.
`background.js`: JavaScript file for background processes.
`content.js`: JavaScript file injected into web pages for content script.

# License

This project is licensed under the MIT License - see the LICENSE file for details.