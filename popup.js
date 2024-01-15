// popup.js
document.getElementById("scrapeButton").addEventListener("click", function() {
  chrome.tabs.query({ active: true, currentWindow: true }, function(tabs) {
    var tab = tabs[0];
    chrome.tabs.sendMessage(tab.id, { action: "scrapeLinkedIn" }, function(response) {
      if (response) {
        displayResult(response);
      }
    });
  });
});

function displayResult(response) {
  var resultDiv = document.getElementById("result");
  if (response.description) {
    resultDiv.innerHTML = response.description;
    resultDiv.style.color = response.description.includes("404") ? "red" : "green";
  } else {
    console.error("Unexpected response format:", response);
    resultDiv.innerHTML = "Error handling response";
    resultDiv.style.color = "red";
  }
}
