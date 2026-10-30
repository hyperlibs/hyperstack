// HMLR DevTools Content Script Bridge
window.addEventListener('message', (event) => {
  if (event.source !== window || !event.data || event.data.source !== 'Hyperstack_DEVTOOLS') return;
  chrome.runtime.sendMessage(event.data);
});

// Inject bridge helper into target page
const script = document.createElement('script');
script.textContent = `
  window.__HMLR_ACTIVE__ = true;
  window.addEventListener('hyperstack:channel', (e) => {
    window.postMessage({ source: 'Hyperstack_DEVTOOLS', type: 'CHANNEL_EVENT', detail: e.detail }, '*');
  });
`;
(document.head || document.documentElement).appendChild(script);
