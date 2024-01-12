document.getElementById('scrapeButton').addEventListener('click', function () {
  chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
    chrome.tabs.sendMessage(tabs[0].id, { action: 'scrapeProfile' }, function (response) {
      document.getElementById('result').innerText = response.description || 'Error scraping profile';
    });
  });
});
