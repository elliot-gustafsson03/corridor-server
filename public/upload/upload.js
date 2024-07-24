window.onload = () => {
    document.querySelector("#upload_btn").addEventListener("click", sendForm);
};

async function sendForm() {
    const image = document.querySelector("#image_input").files[0];
    const label = document.querySelector("#label_input").value;

    if (!image || label == "") {
        alert("Vänligen fyll i alla fält");
        return;
    }

    form = new FormData();
    form.append("image", image);
    form.append("label", label);

    document.querySelector("#loading").style = "";

    const res = await fetch("/api/upload_image", {
        method: "POST",
        body: form,
    });

    const text = await res.text();

    document.querySelector("#loading").style = "display: none";

    if (text === "1") {
        alert("Din bild har laddats upp :D");

        document.querySelector("#image_input").value = null;
        document.querySelector("#label_input").value = "";
    } else {
        alert("Någonting gick fel D:");
    }
}
