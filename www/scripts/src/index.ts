declare var version: string;

import { getValueFormat, formattedValueToString } from '@grafana/data';

const cache = new Map();

var evtSource = new EventSource("//" + document.location.host + "/events?v=" + version);

evtSource.addEventListener("open", (ev) => {
    console.log("SSE ready !");
});

evtSource.addEventListener("error", (ev) => {
    console.error(ev);
});

evtSource.addEventListener("message", (data) => {
    const msg = JSON.parse(data.data);

    console.log(msg);

    if (msg.type === "update") {
        console.log("New version available, refreshing...");

        document.location.reload();

        return;
    }

    const elm = document.getElementById(msg.target);

    if (elm === null) {
        return;
    }

    const previousValue = cache.get(msg.target);

    switch (msg.type) {
        case "status":
            if (previousValue !== msg.payload.status) {
                (elm.children[1] as any).innerText = msg.payload.status;

                cache.set(msg.target, msg.payload.status);
            }
            break;

        case "stat":
            if (previousValue !== msg.payload.value) {
                //(elm.children[1] as any).innerText = msg.payload.value;

                const val = getValueFormat(msg.payload.unit)( msg.payload.value);

                (elm.children[1] as any).innerText = formattedValueToString(val)

                cache.set(msg.target, msg.payload.value);
            }
            break;
        default:
            break;
    }


});
