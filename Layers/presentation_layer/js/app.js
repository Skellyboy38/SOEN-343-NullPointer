$(document).ready(function () {
    buildCalendar(1);
    //getUserReservationsByRoom();
});

function buildCalendar(room_number) {
    // update the calendar information using room_number as the room ID
    var request = getReservations(room_number);
    var data = request == null ? [
        {
            id: "id1",
            description: "Test Data",
            location: "",
            subject: "Testing",
            calendar: "Room 1",
            start: new Date(2016, 11, 9, 10, 0, 0),
            end: new Date(2016, 11, 9, 19, 0, 0)
        },
        {
            id: "id2",
            description: "Test Data",
            location: "",
            subject: "Testing",
            calendar: "Room 1",
            start: new Date(2016, 11, 10, 10, 0, 0),
            end: new Date(2016, 11, 10, 19, 0, 0)
        },
        {
            id: "id3",
            description: "Test Data",
            location: "",
            subject: "Testing",
            calendar: "Room 1",
            start: new Date(2016, 11, 11, 10, 0, 0),
            end: new Date(2016, 11, 11, 19, 0, 0)
        }
    ] : request; // If the json data is null (which it shouldn't be), this default data will appear (to be removed later)
    init(data); // Initialize the calendar with the following data
}

function init(reservations) {
    var appointments = [];
    var ids = [];
    reservations.forEach(function(entry) { // For each reservation, add an appointment to the calendar
        var appointment = {
            id: entry.id,
            description: entry.description,
            location: entry.location,
            subject: entry.subject,
            calendar: entry.calendar,
            start: entry.start,
            end: entry.end,
            readOnly: true
        }
        appointments.push(appointment);
        ids.push(entry.id);
    });

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
            ids.forEach(function(entry) {
                $("#scheduler").jqxScheduler('ensureAppointmentVisible', entry);
            });
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
        error: function (error) {
            return null;
        },
        success: function (data) {
            console.log(data);
            return data;
        }
    });
}