app.updateMatch = (url) => {
    if (app.match === url) {
        return;
    }

    app.match = url;
    if (!url) {
        ELEMENTS.body.classList.remove("single-match");
        ELEMENTS.match_box.style = undefined;
        return;
    }

    ELEMENTS.match_name.innerText = url.get_name();
    ELEMENTS.match_favicon.src = url.get_favicon(32);
    ELEMENTS.match_name.innerText = url.get_name();
    ELEMENTS.match_url.innerText = url.get_pretty_url();
    ELEMENTS.match_url.title = url.get_url();

    var coords = url.get_component().getBoundingClientRect();
    const xNudge = (ELEMENTS.match_box.offsetWidth - url.get_component().offsetWidth) / 2;
    const yNudge = (ELEMENTS.match_box.offsetHeight - url.get_component().offsetHeight) / 2;
    ELEMENTS.match_box.style.top = coords.top - yNudge + "px";
    ELEMENTS.match_box.style.left = coords.left - xNudge + "px";

    ELEMENTS.body.classList.add("single-match");

    setTimeout(() => {
        ELEMENTS.match_box.style.left = "calc(50%)";
        ELEMENTS.match_box.style.top = "calc(50%)";
        ELEMENTS.match_box.style.transform = "translate(-50%, -50%)";
    }, 10);
};
