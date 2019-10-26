(function() {
    "use strict";

    // Make hover effects work on touch devices.
    document.addEventListener("touchstart", function() {}, { passive: true });

    // Build "listening" list
    fetch(window.apiEndpoint).then(function(fetchResp) {
        fetchResp.json().then(function(jsonResp) {
            const latest = jsonResp.listening.recenttracks.track[0];
            const list = document.getElementById("list-listening");
            const item = document.createElement("li");
            item.innerText = `${latest.album["#text"]} by ${latest.artist["#text"]}`;
            list.appendChild(item);
        });
    });
})();
