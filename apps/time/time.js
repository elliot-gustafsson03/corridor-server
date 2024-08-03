const days = [
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
    "Sunday",
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

window.onload = () => {
    clockTick();
    changeDate();

    setInterval(clockTick, 1000);
};

function clockTick() {
    const dateTime = new Date();

    const h = dateTime.getHours().toString().padStart(2, "0");
    const m = dateTime.getMinutes().toString().padStart(2, "0");
    const s = dateTime.getSeconds().toString().padStart(2, "0");
    const time = `${h}:${m}:${s}`;

    document.querySelector("#time").innerHTML = time;

    if (time == "00:00:00") {
        changeDate(dateTime);
    }
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
