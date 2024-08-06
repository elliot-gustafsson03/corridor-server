const days = [
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
];

const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
];

let canvas;
let ctx;

const CLOCK_R = 200;

window.onload = () => {
    canvas = document.querySelector("#canvas");
    ctx = canvas.getContext("2d");

    getNameDay();
    clockTick();
    changeDate();

    setInterval(clockTick, 1000);
};

async function getNameDay() {
    const res = await fetch("/api/get_name_day");
    const json = await res.json();

    document.querySelector("#name_day").style.display =
        json.length === 0 ? "none" : "block";
    if (json.length > 0) {
        const nameDay = `Concratulations to ${json.join(" and ")} on their name day!`;
        document.querySelector("#name_day").innerHTML = nameDay;
    }
}

function clockTick() {
    const dateTime = new Date();

    const h = dateTime.getHours().toString().padStart(2, "0");
    const m = dateTime.getMinutes().toString().padStart(2, "0");
    const s = dateTime.getSeconds().toString().padStart(2, "0");
    const time = `${h}:${m}:${s}`;

    document.querySelector("#time").innerHTML = time;

    if (time === "00:00:00") {
        changeDate(dateTime);
        getNameDay();
    }

    updateClock(
        dateTime.getHours(),
        dateTime.getMinutes(),
        dateTime.getSeconds(),
    );
}

function changeDate() {
    const dateTime = new Date();

    const weekDay = days[dateTime.getDay()];
    const day = dateTime.getDate();
    const month = months[dateTime.getMonth()];
    const year = dateTime.getFullYear();

    let dayEnding = "th";
    if (day % 10 == 1 && day != 11) {
        dayEnding = "st";
    } else if (day % 10 == 2 && day != 12) {
        dayEnding = "nd";
    } else if (day % 10 == 3 && day != 13) {
        dayEnding = "rd";
    }

    const date = `${weekDay}, the ${day}${dayEnding} of ${month}, ${year}`;

    document.querySelector("#date").innerHTML = date;
}

function updateClock(h, m, s) {
    ctx.clearRect(0, 0, 2 * CLOCK_R, 2 * CLOCK_R);
    drawClock();
    drawPointers(h, m, s);
}

function drawClock() {
    ctx.beginPath();
    ctx.arc(CLOCK_R, CLOCK_R, CLOCK_R - 10, 0, 2 * Math.PI, false);
    ctx.fillStyle = "white";
    ctx.fill();
    ctx.strokeStyle = "black";
    ctx.lineWidth = 8;
    ctx.stroke();
}

function drawPointers(h, m, s) {
    ctx.lineWidth = 8;
    drawPointer(h / 12, CLOCK_R * 0.4);
    ctx.lineWidth = 4;
    drawPointer(m / 60, CLOCK_R * 0.7);
    ctx.lineWidth = 2;
    drawPointer(s / 60, CLOCK_R * 0.8);
}

function drawPointer(ratio, length) {
    const angle = 2 * Math.PI * ratio - Math.PI / 2;

    ctx.beginPath();
    ctx.moveTo(CLOCK_R, CLOCK_R);
    ctx.lineTo(
        CLOCK_R + Math.cos(angle) * length,
        CLOCK_R + Math.sin(angle) * length,
    );
    ctx.stroke();
    ctx.closePath();
}
