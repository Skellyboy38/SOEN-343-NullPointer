function buildHTML(tag, html, attrs) {
    if (typeof (html) != 'string') {
        attrs = html;
        html = null;
    }
    var h = '<' + tag;
    for (attr in attrs) {
        if (attrs[attr] === false)
            continue;
        h += ' ' + attr + '="' + attrs[attr] + '"';
    }
    return h += html ? ">" + html + "</" + tag + ">" : "/>";
}

function printTodayDate(){
    var today = new Date();
    var options = {
    weekday: "long", year: "numeric", month: "short",
    day: "numeric", hour: "2-digit", minute: "2-digit"
    };
    $("#todayDate").html(today.toLocaleTimeString("en-us", options));
}

function getCookie(name) {
    var value = "; " + document.cookie;
    var parts = value.split("; " + name + "=");
    if (parts.length == 2) return parts.pop().split(";").shift();
}

function renderUserReservationList(reservations) {
    var reservationListHTML = $(".reservations-table");
    var reservationHeaderHTML = renderReservationHeader();
    reservationHeaderHTML.appendTo(reservationListHTML);
    if (reservations.length === 0) {
        var row = $("<div></div>", {
            class: "row",
            text: "No reservations available."
        })
        row.appendTo(reservationListHTML);
    } else {
        reservations.forEach(function (resv) {
            var row = renderReservationRow(resv);
            row.appendTo(reservationListHTML);
        });
    }

    function renderReservationRow(resv) {
        var rowHTML = $("<div></div>", {
            class: "row"
        });
        var roomNumberCell = $("<div></div>", {
            text: resv.roomNumber,
            class: "cell"
        });
        roomNumberCell.appendTo(rowHTML);
        var descriptionCell = $("<div></div>", {
            text: "TBD",
            class: "cell"
        });
        descriptionCell.appendTo(rowHTML);
        var startTimeCell = $("<div></div>", {
            text: resv.startTime,
            class: "cell"
        });
        startTimeCell.appendTo(rowHTML);
        var endTimeCell = $("<div></div>", {
            text: resv.endTime,
            class: "cell"
        });
        endTimeCell.appendTo(rowHTML);
        var actionsCell = $("<div></div>", {
            class: "cell btnFlex"
        });

        var modifyBtn = $("<a></a>", {
            class: "Waves-effect waves-light btn modifyBtn",
            text: "Modify",
            "data-reservationid": resv.reservationID
        });

        modifyBtn.data("reservationID", resv.reservationID);
        modifyBtn.click(function () {
            $('#modifyModal').modal({
                dismissible: true,
                opacity: .8,
                in_duration: 300,
                out_duration: 200,
                starting_top: '4%',
                ending_top: '10%',
                ready: function (modal, trigger) {
                    clearModal();
                    initializeModifyModal(resv);
                    // Callback for Modal open. Modal and trigger parameters available.
                },
                complete: function () { }
            });
            $('#modifyModal').modal('open');
        });
        modifyBtn.appendTo(actionsCell)

        var deleteBtn = $("<a></a>", {
            class: "Waves-effect waves-light btn deleteBtn",
            text: "Delete",
            "data-reservationid": resv.reservationID
        });

        deleteBtn.data("reservationID", resv.reservationID);
        deleteBtn.click(function () {
            deleteReservation($(this).attr("data-reservationid"));
        });
        deleteBtn.appendTo(actionsCell)
        actionsCell.appendTo(rowHTML);
        return rowHTML;
    }
}

function renderReservationHeader() {
    var header = $("<div></div>", {
        class: "row"
    });
    header.addClass("row reservations-header");
    var roomNumberCell = $("<div></div>", {
        class: "cell",
        text: "Room Number"
    });
    var descriptionCell = $("<div></div>", {
        class: "cell",
        text: "Description"
    });
    var startTimeCell = $("<div></div>", {
        class: "cell",
        text: "Start Time"
    });
    var endTimeCell = $("<div></div>", {
        class: "cell",
        text: "End Time"
    });
    var actionsCell = $("<div></div>", {
        class: "cell",
        text: "Actions"
    });
    roomNumberCell.appendTo(header);
    descriptionCell.appendTo(header);
    startTimeCell.appendTo(header);
    endTimeCell.appendTo(header);
    actionsCell.appendTo(header);
    return header;
}
