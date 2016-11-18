$(document).ready(function () {
    buildCalendar(1);
});

function buildCalendar(roomNumber) {
    var roomReservations = [];
    var userRoomReservations = [];
    var studentId = getCookie("studentId");

    getReservations(roomNumber).success(function(data){
        roomReservations = getReservationsSuccess(data);
    });
    getReservationsUser(roomNumber, studentId).success(function(data){
        renderUserReservationList(data);
        userRoomReservations = getReservationsUserSuccess(data);
    });
    // TODO - Need to distinguish user's reservations and rest of resrvation
    init(roomReservations);
}

function init(reservations) {
    var source =
        {
            dataType: "array",
            dataFields: [
                { name: 'subject', type: 'string' },
                { name: 'calendar', type: 'string' },
                { name: 'start', type: 'date' },
                { name: 'end', type: 'date' },
                { name: 'readOnly', type:'boolean' }
            ],
            id: 'id',
            localData: reservations
        };
    var adapter = new $.jqx.dataAdapter(source);
    $("#scheduler").jqxScheduler({
        date: new $.jqx.date(new Date()),
        width: 1000,
        height: 600,
        source: adapter,
        theme: 'metrodark',
        view: 'weekView',
        showLegend: true,
        contextMenu: false,
        editDialog: false,
        resources:
        {
            colorScheme: "scheme05",
            dataField: "calendar",
            source: new $.jqx.dataAdapter(source)
        },
        appointmentDataFields:
        {
            from: "start",
            to: "end",
            id: "id",
            subject: "subject",
            resourceId: "calendar",
            readOnly: "readOnly",
        },
        views:
        [
            'dayView',
            'weekView',
            'monthView'
        ]
    });
}

function renderUserReservationList(reservations){
    console.log(reservations);
    var reservationListHTML = $(".reservations-table");
    reservations.forEach(function(resv){
        var row = renderReservationRow(resv);
        row.appendTo(reservationListHTML);
    });

    function renderReservationRow(resv){
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
            text: "Save / Delete Buttons",
            class: "cell"
        });
        actionsCell.appendTo(rowHTML);
        return rowHTML;
    }
}

function getReservations(roomNumber) {
    return $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        async: false, 
        url: '/reservationsByRoom',
        data: {dataRoom: roomNumber},
    });
}

function getReservationsSuccess(data){
    return deserializeReservation(data);
}

function getReservationsUserSuccess(data){
    return deserializeReservation(data);
}


function getReservationsUser(roomNumber, userID) {
    return $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        async: false, 
        url: '/reservationsByUser',
        data: {roomID: roomNumber, userID: userID},
    });
}

function createReservation() {
    $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        url: '/createReservation',
        data: {},
    });
}

function deleteReservation() {
    $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        url: '/deleteReservation',
        data: {},
    });
}

function deserializeReservation(reservations){
    if(reservations != undefined && reservations.length > 0){
        var result = [];
        reservations.forEach(function(reservationJSON){
            var reservation = new Reservation(reservationJSON.reservationID, reservationJSON.subject, reservationJSON.roomNumber, reservationJSON.startTime, reservationJSON.endTime);
            result.push(reservation)
        })
    } else{
        console.error("No reservations found.");
        return [];
    }
    return result;
}