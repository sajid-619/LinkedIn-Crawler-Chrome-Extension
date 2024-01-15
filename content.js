// content.js
chrome.runtime.onMessage.addListener(function(request, sender, sendResponse) {
    if (request.action === "scrapeLinkedIn") {
      var userURI = getUserURIFromLinkedInURL(window.location.href);
      fetch(`http://localhost:8080/api/profileInfo/get?path=${userURI}`)
        .then(response => response.json())
        .then(data => sendResponse(data))
        .catch(error => sendResponse({ description: "404 - Not Found" }));
    }
    return true;
  });
  
  function getUserURIFromLinkedInURL(url) {
    // Extract the dynamic user URI from the LinkedIn URL
    // Example: "https://www.linkedin.com/in/johndoe"
    var match = url.match(/\/in\/([^\/?]+)/);
    if (match && match[1]) {
      return match[1];
    }
    return ""; // Return an empty string if the pattern is not matched
  }
  