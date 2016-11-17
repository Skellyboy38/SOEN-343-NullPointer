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
        userRoomReservations = getReservationsUserSuccess(data);
    });

    console.log(roomReservations);
    console.log(userRoomReservations);
    init(roomReservations);
}

function init(reservations) {
    console.log(reservations);
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

function renderUserReservationList(){
    // TODO
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

function getCookie(name) {
  var value = "; " + document.cookie;
  var parts = value.split("; " + name + "=");
  if (parts.length == 2) return parts.pop().split(";").shift();
}

function deserializeReservation(reservations){
    if(reservations != undefined && reservations.length > 0){
        var result = [];
        reservations.forEach(function(reservationJSON){
            var reservation = new Reservation(reservationJSON.subject, reservationJSON.roomNumber, reservationJSON.startTime, reservationJSON.endTime);
            result.push(reservation)
        })
    } else{
        console.error("No reservations found.");
        return [];
    }
    return result;
}