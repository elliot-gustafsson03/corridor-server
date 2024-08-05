async function changeApp(appName) {
    document.querySelector("#loading").style.display = "block";

    const res = await fetch("/api/change_app", {
        method: "POST",
        body: appName,
    });
    const text = await res.text();

    document.querySelector("#loading").style.display = "none";

    if (text == "1") {
        alert("The app has been changed :D");
    } else {
        alert("Couldn't change app D:");
    }
}
