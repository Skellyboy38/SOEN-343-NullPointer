$(document).ready(function () {
    init();
    updateCalendar(1);
    //getUserReservationsByRoom();
});

function updateCalendar(room_number) {
    // update the calendar information using room_number as the room ID
    getReservations(room_number);
}

function init() {

    // Test Data - TO BE REMOVED
    var appointments = new Array();
    var appointment1 = {
        id: "id1",
        description: "Test Data",
        location: "",
        subject: "Testing",
        calendar: "Room 1",
        start: new Date(2015, 10, 23, 9, 0, 0),
        end: new Date(2015, 10, 23, 16, 0, 0)
    }
    appointments.push(appointment1);

    var source =
        {
            dataType: "array",
            dataFields: [
                { name: 'id', type: 'string' },
                { name: 'description', type: 'string' },
                { name: 'location', type: 'string' },
                { name: 'subject', type: 'string' },
                { name: 'calendar', type: 'string' },
                { name: 'start', type: 'date' },
                { name: 'end', type: 'date' },
                { name: 'readOnly', type:'boolean' }
            ],
            id: 'id',
            localData: appointments
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
        ready: function () {
            //$("#scheduler").jqxScheduler('ensureAppointmentVisible', 'id1');
        },
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
            description: "description",
            location: "place",
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
    $.ajax({
        type: 'POST',
        contentType: "application/json",
        url: '/reservations',
        data: JSON.stringify({roomID: room_number}),
        dataType: "json",
        error: function (error) { },
        success: function (data) {
            console.log(data);
        }
    });
}