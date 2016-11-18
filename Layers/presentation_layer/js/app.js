$(document).ready(function () {
    buildCalendar(1);
});

function buildCalendar(roomNumber, el) {
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
    if(el != null){
        $(".tab").attr("id", "");
        $(el).attr("id", "active");
    }
    // TODO - Need to distinguish user's reservations and rest of reesrvation
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
    var reservationListHTML = $(".reservations-table");
    var reservationHeaderHTML = renderReservationHeader();
    reservationHeaderHTML.appendTo(reservationListHTML);
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

function renderReservationHeader(){
    //var reservationsListHTML = $(".reservations-table");
    var header = $("<div></div>", {
            class: "row"
        });
    header.addClass("row reservations-header");
    var roomNumberCell = $("<div></div>", {
        class:"cell",
        text: "Room Number"
    });
    var descriptionCell = $("<div></div>", {
        class:"cell",
        text: "Description"
    });
    var startTimeCell = $("<div></div>",{
        class:"cell",
        text: "Start Time"
    });
    var endTimeCell = $("<div></div>",{
        class:"cell",
        text: "End Time"
    });
    var actionsCell = $("<div></div>",{
        class:"cell",
        text: "Actions"
    });
    roomNumberCell.appendTo(header);
    descriptionCell.appendTo(header);
    startTimeCell.appendTo(header);
    endTimeCell.appendTo(header);
    actionsCell.appendTo(header);
    return header;
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
	var userID = getCookie("studentId");
	var room = $("#room").val();
	var start = $("#start_time").val();
	var end = $("#end_time").val();
	//console.log(userID + " " + room + " " + date + " " + start + " " + end);
    $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        async: false,
        url: '/createReservation',
        data: {userID: userID, roomID: room, start: start, end: end},
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

function changeRoom(roomNumber, el){
    $(".reservations-table").empty();
    buildCalendar(roomNumber, el);
}