const FolderComponent = (spec) => {
    const folder = spec.folder;
    if (!folder) {
        throw new Error("Cannot construct folder component without folder!");
    }

    const component = document.createElement("div");
    component.classList.add("folder");

    const nameContainer = document.createElement("div");
    nameContainer.innerText = folder.get_name();
    nameContainer.classList.add("folder-name");
    component.appendChild(nameContainer);

    const contentContainer = document.createElement("div");
    contentContainer.classList.add("folder-content");
    component.appendChild(contentContainer);

    const has_subfolders = folder.get_subfolders().length !== 0;
    const has_urls = folder.get_urls().length !== 0;

    // Just urls
    if (has_urls && !has_subfolders) {
        contentContainer.classList.add("urls");
        for (let url of folder.get_urls()) {
            const spec = {};
            contentContainer.appendChild(UrlComponent({ ...spec, url }));
        }
    }

    // Mixed
    if (has_urls && has_subfolders) {
        const rawUrls = folder.get_urls().flatMap((url) => url.get_raw());
        const rawFolder = Folder({ name: "Unsorted", children: rawUrls });
        contentContainer.appendChild(FolderComponent({ folder: rawFolder }));
    }

    // Folders
    for (let subfolder of folder.get_subfolders()) {
        const spec = {};
        contentContainer.appendChild(FolderComponent({ ...spec, folder: subfolder }));
    }

    folder.set_component(component);
    return component;
};

const UrlComponent = (spec) => {
    const url = spec.url;
    if (!url) {
        throw new Error("Cannot construct URL component without URL!");
    }

    const component = document.createElement("div");
    component.classList.add("url");
    component.onclick = () => {
        open(url.get_url(), "_self");
    };

    const favico = document.createElement("img");
    favico.classList.add("favicon");
    favico.src = url.get_favicon();
    favico.addEventListener("load", console.clear); // Clear console on load to prevent error spam
    component.appendChild(favico);

    const textContainer = document.createElement("a");
    textContainer.innerText = url.get_name();
    component.title = url.get_url();
    component.appendChild(textContainer);
    url.set_component(component);
    return component;
};

const ConstructBookmarkTree = (rootNode) => {
    const subfolders = rootNode.get_subfolders() ?? [];
    for (let [i, folder] of subfolders.entries()) {
        const spec = { folder };
        var rootFolder = FolderComponent(spec);
        rootFolder.classList.add("root");
        rootFolder.style.backgroundColor = `var(${ROOT_FOLDER_COLORS[i]})`;
        document.querySelector("#bookmarks").appendChild(rootFolder);
    }
};
