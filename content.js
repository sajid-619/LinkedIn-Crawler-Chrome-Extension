chrome.runtime.onMessage.addListener(
  function (request, sender, sendResponse) {
    if (request.action === 'scrapeProfile') {
      // Modify the script based on your LinkedIn page structure
      var description = document.querySelector('div.entity-result__primary-subtitle.t-14.t-black.t-normal').innerText;

      sendResponse({ description: description });
    }
  }
);
