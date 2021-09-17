function showCookieDisclaimer() {
    return {
        show: false,
        cookieName: 'cookie_disclaimer',
        setCookie(value) {
            let date = new Date();
            date.setTime(date.getTime() + (7 * 24 * 60 * 60 * 1000));
            document.cookie = this.cookieName + '=' + value + ';expires=' + date.toGMTString() + ';path=/;SameSite=Strict';
            this.show = false;
            console.log('setCookie', value);
        },
        getCookie() {
            const cookie = document.cookie.match(new RegExp('(^|;)\\s*' + this.cookieName + '=([^;]*)'));
            this.show = cookie ? false : true;
        }
    }
}