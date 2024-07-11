window.onload = () => {
    document.querySelector("#upload_btn").addEventListener("click", sendForm);
};

async function sendForm() {
    const image = document.querySelector("#image_input").files[0];
    const label = document.querySelector("#label_input").value;

    form = new FormData();
    form.append("image", image);
    form.append("label", label);

    const res = await fetch("/api/upload_image", {
        method: "POST",
        body: form,
    });
}
