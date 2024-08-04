function changeApp(appName) {
    fetch("/api/change_app", {
        method: "POST",
        body: appName,
    });
}
