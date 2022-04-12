module.exports = function BrowserEngine() {
    const Browser = require('zombie');

    return function bypass(proxy, uagent, callback) {
        var cookie = false;
        var browser = new Browser();
        browser.maxDuration = 400e3;
        browser.maxWait = 380e3;
        browser.proxy = proxy;
        browser.userAgent = uagent;
        browser.waitDuration = '1000s';
        browser.visit(l7.target, () => {
            browser.wait(370e3, () => {
                browser.reload();
                browser.wait(50e3, async () => {
                    if (browser.cookies.length > 0) {
                        await browser.cookies.forEach(acookie => {
                            if (!cookie) {
                                cookie = acookie.key + '=' + acookie.value;
                            } else {
                                cookie += ('; ' + acookie.key + '=' + acookie.value);
                            }
                        });
                        await callback(cookie);
                    } else {
                        await callback(false);
                    }
                    await browser.deleteCookies();
                    await browser.window.close();
                    await browser.destroy();
                    browser = undefined;
                    delete browser;
                    return false;
                })
            });
        });
    }
}