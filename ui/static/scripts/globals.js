// Global dependencies
const app = {
    nodeMap: {},
    match: undefined,
};

const NODE_TYPES = {
    url: "url",
    folder: "folder",
};

const ROOT_FOLDER_COLORS = ["--midnight-green", "--dark-cyan", "--gamboge", "--alloy-orange", "--rust", "--rufous", "--auburn"];

const ELEMENTS = {
    body: document.querySelector("body"),
    match_box: document.querySelector("#match-box"),
    match_favicon: document.querySelector("#match-box-favicon"),
    match_name: document.querySelector("#match-box-name"),
    match_url: document.querySelector("#match-box-url"),
};
