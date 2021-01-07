const express = require('express');
const app = express();
app.get("/codedeployexample", (req, res, next) => {
    const label = "Application currently";

    let currentStateMessage = "is DESTROYED";
    let currentColor = "red";

    if (false == true) {
        currentStateMessage = "exists"
        currentColor = "green";
    }

    res.json({
        "schemaVersion": 1,
        "label": label,
        "message": currentStateMessage,
        "color": currentColor
    });
});

module.exports = app;