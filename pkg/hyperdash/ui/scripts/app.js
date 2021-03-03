var evtSource = new EventSource("//localhost:8080/events");

evtSource.addEventListener("error", (ev) => {
    console.error(ev);
});

evtSource.addEventListener("message", (data) => {
    console.log(data);

    const msg = JSON.parse(data.data);

    console.log(msg);

    const elm = document.getElementById(msg.target);

    elm.innerText = msg.payload;
});

console.log("Hello");
