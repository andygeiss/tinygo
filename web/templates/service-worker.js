var cacheName = '{{ .Name }}';
var dataCacheName = '{{ .Name }}';
var filesToCache = [];

self.addEventListener('install', function(e) {
    e.waitUntil(
        caches.open(cacheName).then(function(cache) {
            return cache.addAll(filesToCache);
        })
    );
});

self.addEventListener('activate', function(e) {
    e.waitUntil(
        caches.keys().then(function(keyList) {
            return Promise.all(keyList.map(function(key) {
                if (key !== cacheName && key !== dataCacheName) {
                    return caches.delete(key);
                }
            }));
        })
    );
    return self.clients.claim();
});

self.addEventListener('fetch', function(e) {
    e.respondWith(
        fetch(e.request).then(function(response) {
            console.log("ServiceWorker fetch from backend: " + e.request.url);
            if (!response.ok) {
                // error response won't cause a reject - do an explicit throw instead.
                throw Error(response.status);
            }
            return response;
        }).catch(function (error) {
            // if fetch failed then try to get request from the cache.
            caches.match(e.request).then(function(response) {
                console.log("ServiceWorker fetch from cache: " + e.request.url);
                return response;
            })
        })
    );
});
