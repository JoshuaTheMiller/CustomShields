module.exports = (req, res) => {
    const label = "Application currently";

    let currentStateMessage = "is DESTROYED";
    let currentColor = "red";

    // if () {
    //     currentStateMessage = "exists"
    //     currentColor = "green";
    // }

    res.json({
        "schemaVersion": 1,
        "label": label,
        "message": currentStateMessage,
        "color": currentColor
    });
}