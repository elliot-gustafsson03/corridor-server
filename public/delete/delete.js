window.onload = async () => {
    const res = await fetch("/api/get_all_images");
    const json = await res.json();

    let html = "";

    for (let i = 0; i < json.length; i++) {
        html += `<div class = "item">
            <div class = "grid">
                <div>
                    <img src = "/images/${json[i].image}" />
                </div>
                <div class = "label">
                    <p>${json[i].label}</p>
                </div>
            </div>
            <span onclick = "deleteImage(${json[i].id})">[X]</span>
        </div>`;
    }

    document.querySelector("#delete_items").innerHTML = html;
};

async function deleteImage(id) {
    await fetch("/api/delete_image", {
        method: "POST",
        body: JSON.stringify({ id: id }),
    });

    location.reload();
}
