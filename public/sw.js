// Hyperstack PWA Service Worker — Offline-First Caching
const CACHE_NAME = 'hyperstack-bio-pwa-v1';
const PRECACHE_ASSETS = [
  '/bio',
  '/manifest.json',
  '/htmx-bolt.js',
  '/htmx-vibe.js',
  '/appz.js',
  '/vendor/htmx.min.js',
  '/styles.css',
  '/icons/icon-192.svg',
  '/icons/icon-512.svg'
];

self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open(CACHE_NAME).then((cache) => {
      return cache.addAll(PRECACHE_ASSETS).catch((err) => {
        console.warn('[PWA SW] Pre-cache warning:', err);
      });
    }).then(() => self.skipWaiting())
  );
});

self.addEventListener('activate', (event) => {
  event.waitUntil(
    caches.keys().then((keys) => {
      return Promise.all(
        keys.filter((key) => key !== CACHE_NAME).map((key) => caches.delete(key))
      );
    }).then(() => self.clients.claim())
  );
});

self.addEventListener('fetch', (event) => {
  const req = event.request;
  if (req.method !== 'GET') return;

  const url = new URL(req.url);

  // HTML Navigation: Network first, fall back to cached /bio
  if (req.mode === 'navigate' || url.pathname === '/bio' || url.pathname === '/') {
    event.respondWith(
      fetch(req)
        .then((networkRes) => {
          if (networkRes && networkRes.status === 200) {
            const resClone = networkRes.clone();
            caches.open(CACHE_NAME).then((cache) => cache.put(req, resClone));
          }
          return networkRes;
        })
        .catch(() => {
          return caches.match('/bio').then((cached) => cached || caches.match(req));
        })
    );
    return;
  }

  // Static Assets: Cache first, fall back to network
  event.respondWith(
    caches.match(req).then((cachedRes) => {
      if (cachedRes) return cachedRes;
      return fetch(req).then((networkRes) => {
        if (networkRes && networkRes.status === 200 && url.origin === location.origin) {
          const resClone = networkRes.clone();
          caches.open(CACHE_NAME).then((cache) => cache.put(req, resClone));
        }
        return networkRes;
      });
    })
  );
});
