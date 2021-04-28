function FindProxyForURL(url, host) {
    if (shExpMatch(url, "*.google.com/*")) {
        return "PROXY 192.168.126.220:1080";
    }

    if (shExpMatch(url, "*.github.com/*")) {
        return "PROXY 192.168.126.220:1080";
    }

    if (shExpMatch(url, "*.githubusercontent.com/*")) {
        return "PROXY 192.168.126.220:1080";
    }

    if (isInNet(host, "10.0.0.0", "255.0.0.0")) {
        return "Direct";
    }
    return "Direct, PROXY 192.168.126.220:1080";
}