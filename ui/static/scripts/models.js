/**
 * spec:
 * id: string
 * name: string
 * type: string
 * parent?: Node
 */
const Node = (spec) => {
    const id = spec.id ?? crypto.randomUUID();
    const name = spec.name;
    const type = spec.type;
    let component;

    const that = {
        get_id: () => {
            return id;
        },
        get_name: () => {
            return name;
        },
        get_type: () => {
            return type;
        },
        get_component: () => {
            return component;
        },
        set_component: (c) => {
            component = c;
        },
    };

    return that;
};

/**
 * spec:
 * ...Node
 * url?: []
 */
const Url = (spec) => {
    spec.type = NODE_TYPES.url;
    const url = spec.url;
    const that = Node(spec);

    app.nodeMap[that.get_id()] = {
        ...that,
        get_url: () => {
            return url;
        },
        get_raw: () => {
            return {
                type: NODE_TYPES.url,
                name: that.get_name(),
                id: that.get_id(),
                url: url,
            };
        },
        get_pretty_url: () => {
            const scheme = url.indexOf("://") + 3;
            const tld = url.indexOf("/", scheme);
            let pretty_url = url.slice(0, tld ?? url.length);
            if (tld + 1 < url.length) {
                const subdirectory = url.lastIndexOf("/", url);
                const suffix = Math.max(subdirectory, url.length - 15);
                pretty_url += "/..." + url.slice(suffix, url.length);
            }
            return pretty_url;
        },
        get_favicon: (size) => {
            if (!size) {
                size = 16;
            }
            return `https://s2.googleusercontent.com/s2/favicons?sz=${size}&domain_url=${url}/"`;
        },
    };

    return app.nodeMap[that.get_id()];
};

/**
 * spec:
 * ...Node
 * children?: []
 */
const Folder = (spec) => {
    spec.type = NODE_TYPES.folder;
    const that = Node(spec);

    const urls = [];
    const subfolders = [];
    for (let child of spec.children ?? []) {
        if (child.type === NODE_TYPES.folder) {
            subfolders.push(Folder(child));
        } else if (child.type === NODE_TYPES.url) {
            urls.push(Url(child));
        } else {
            console.warn(`Node with name ${spec.name} has no defined type, cannot be constructed`);
        }
    }

    app.nodeMap[that.get_id()] = {
        ...that,
        get_urls: () => {
            return urls;
        },
        get_subfolders: () => {
            return subfolders;
        },
        get_length: () => {
            return urls.length + subfolders.length;
        },
        get_raw: () => {
            return {
                id: that.get_id(),
                name: that.get_name(),
                children: urls.flatMap((url) => url.get_raw()) + subfolders.flatMap((sf) => sf.get_raw()),
            };
        },
    };

    return app.nodeMap[that.get_id()];
};
