(function () {
    "use strict";

    // Make hover effects work on touch devices.
    document.addEventListener("touchstart", function () {}, { passive: true });

    fetch(window.apiEndpoint).then(function (fetchResp) {
        fetchResp.json().then(function (jsonResp) {
            buildListeningList(jsonResp.listening);
            buildReading(jsonResp.reading);
            buildWorking(jsonResp.working);
        });
    });

    function buildListeningList(listening) {
        let albumSet = {};
        listening.forEach(function (track) {
            albumSet[`${track.album} by ${track.artist}`] = "";
        });
        albumSet = Object.keys(albumSet);
        albumSet.length = 4;

        const albumListFragment = document.createDocumentFragment();
        albumSet.forEach(function (album) {
            const item = document.createElement("li");
            item.innerText = album;
            albumListFragment.appendChild(item);
        });

        const list = document.getElementById("list-listening");
        list.appendChild(albumListFragment);
    }

    function buildReading(reading) {
        const readingListFragment = document.createDocumentFragment();
        reading.forEach(function (book) {
            const item = document.createElement("li");
            const authors = book.authors.join(", ");
            item.innerText = `${book.title} by ${authors}`;
            readingListFragment.appendChild(item);
        });
        const list = document.getElementById("list-reading");
        list.appendChild(readingListFragment);
    }

    function buildWorking(working) {
        const workingListFragment = document.createDocumentFragment();
        working.forEach(function (work) {
            const item = document.createElement("li");
            const link = document.createElement("a");
            link.href = work.url;
            link.innerText = work.title;
            item.appendChild(link);
            workingListFragment.appendChild(item);
        });
        const list = document.getElementById("list-working");
        list.appendChild(workingListFragment);
    }
})();
