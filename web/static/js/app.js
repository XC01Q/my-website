document.addEventListener('DOMContentLoaded', function() {
    const data = {
        screenWidth: window.screen.width,
        screenHeight: window.screen.height,
        viewportWidth: window.innerWidth,
        viewportHeight: window.innerHeight,
        userAgent: navigator.userAgent,
        language: navigator.language,
        platform: navigator.platform,
        colorDepth: window.screen.colorDepth,
        pixelRatio: window.devicePixelRatio
    };

    document.body.addEventListener('htmx:configRequest', function(evt) {
        evt.detail.headers['X-Client-Data'] = JSON.stringify(data);
    });
});
