const conn = new WebSocket("ws://" + window.location.host + "/ws/connect");
conn.onmessage = (e) => {
    location.href = "/apps/" + e.data;
};
