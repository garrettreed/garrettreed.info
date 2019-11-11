(function() {
    "use strict";

    // Make hover effects work on touch devices.
    document.addEventListener("touchstart", function() {}, { passive: true });

    fetch(window.apiEndpoint).then(function(fetchResp) {
        fetchResp.json().then(function(jsonResp) {
            buildListeningList(jsonResp.listening);
            buildReading(jsonResp.reading);
        });
    });

    function buildListeningList(listening) {
        let albumSet = {};
        listening.tracks.forEach(function(track) {
            albumSet[`${track.album} by ${track.artist}`] = "";
        });
        albumSet = Object.keys(albumSet);
        albumSet.length = 4;

        const albumListFragment = document.createDocumentFragment();
        albumSet.forEach(function(album) {
            const item = document.createElement("li");
            item.innerText = album;
            albumListFragment.appendChild(item);
        });

        const list = document.getElementById("list-listening");
        list.appendChild(albumListFragment);
    }

    function buildReading(reading) {
        const readingListFragment = document.createDocumentFragment();
        reading.reviews.forEach(function(review) {
            const item = document.createElement("li");
            const authors = review.book.authors
                .reduce(function(authorList, author) {
                    return authorList.concat([author.name]);
                }, [])
                .join(", ");
            item.innerText = `${review.book.title} by ${authors}`;
            readingListFragment.appendChild(item);
        });
        const list = document.getElementById("list-reading");
        list.appendChild(readingListFragment);
    }
})();
