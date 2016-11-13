$(document).ready(function () {
    buildCalendar(123);
});

function buildCalendar(room_number) {
    var roomReservations;
    getReservations(room_number).success(function getReservationsSuccess(data){
        if(data != undefined && data.length > 0){
            var result = [];
            data.forEach(function(reservationJSON){
                var reservation = new Reservation(reservationJSON.subject, reservationJSON.roomNumber, reservationJSON.startTime, reservationJSON.endTime);
                result.push(reservation)
            })
        } else{
            console.error("No reservations found.");
        }
        roomReservations = result;
    });
    //var userReservations = getUserReservations();
    init(roomReservations); // Initialize the calendar with the following data
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
        width: 1200,
        height: 800,
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

function getReservations(room_number) {
    return $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        async: false, 
        url: '/reservationsByRoom',
        data: {dataRoom: room_number},
    });
}


function getUserReservations(room_number) {
    $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        url: '/reservations',
        data: {roomID: room_number},
        error: function (error) {
            return null;
        },
        success: function (data) {
            return data;
        }
    });
}