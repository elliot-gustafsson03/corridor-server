const INTERVAL_TIME = 5;

async function fetchImage() {
    const res = await fetch("/api/get_next_image")
    const json = await res.json()

    document.querySelector("#image").src = "/images/" + json.image
    document.querySelector("#label").innerHTML = json.label
}

window.onload = () => {
    fetchImage()
    setInterval(fetchImage, INTERVAL_TIME * 1000)
}