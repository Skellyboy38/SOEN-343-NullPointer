$(document).ready(function () {
    populateTime();
    populateEndTime();
    populateDays(1, 2016);
    printTodayDate(); // Displays current time and date
    buildCalendar(1); // Default room is 1
});

function buildCalendar(roomNumber, el) {
    var roomReservations = [];
    var userRoomReservations = [];
    var roomReservationsNonUser = [];
    var studentId = getCookie("studentId");

    getReservations(roomNumber).success(function(data){
        roomReservations = deserializeReservation(data);
    });

    getReservationsUser(roomNumber, studentId).success(function(data){
        renderUserReservationList(data);
        userRoomReservations = deserializeReservation(data);
    });

    getReservationsNonUser(roomNumber, studentId).success(function(data){
        roomReservationsNonUser = deserializeReservation(data);
    });

    if(el != null){
        $(".tab").attr("id", "");
        $(el).attr("id", "active");
    }
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
        theme: 'office',
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

function getReservationsNonUser(roomNumber, userID) {
    return $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        async: false, 
        url: '/reservationsOthers',
        data: {dataRoom: roomNumber, userID: userID},
    });
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
    var span = $("#message");
    span.html("");
	var userID = getCookie("studentId");
	var room = $("#room").val();
    var year = $("#year").val();
    var month = $("#month").val();
    var day = $("#day").val();

	var start = $("#start_time").val();
	var end = $("#end_time").val();
    if(room == null || year == null || month == null || day == null || start == null || end == null) {
        span.html("Missing information.");
        return;
    }
    var start_time = "";
    var end_time = "";
    var day_time = "";

    if(parseInt(day)<10) {
        day_time = "0" + String(day);
    }
    else {
        day_time = String(day);
    }

    if(parseInt(start)<10) {
        start_time = "0" + String(start);
    }
    else {
        start_time = String(start);
    }

    if(parseInt(end)<10) {
        end_time = "0" + String(end);
    }
    else {
        end_time = String(end);
    }

    var startDate = String(year) + "-" + String(month) + "-" + day_time + " " + start_time + ":00:00";
    var endDate = String(year) + "-" + String(month) + "-" + day_time + " " + end_time + ":00:00";

    if(!verifyTimeConflicts(room, startDate, endDate)) {
        pushReservation(userID, room, startDate, endDate);
        location.reload();
        span.html("Reservation created.");
    }
    else { // Add the person to a wait list
        $.ajax({
            type: 'POST',
            contentType: "application/x-www-form-urlencoded",
            async: false,
            url: '/addToWaitList',
            data: {userID: userID, dataRoom: room, startTime: startDate, endTime: endDate},
        });
        span.html("Time conflict. Added to wait list.");
    }
}

function pushReservation(userID, room, startDate, endDate) {
    $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        async: false,
        url: '/createReservation',
        data: {userID: userID, dataRoom: room, startTime: startDate, endTime: endDate},
    });
}

function splitTime(time) {
    var segments = [];
    var split = time.split("-")
    segments.push(split[0]);
    segments.push(split[1]);
    var splitDay = split[2].split(" ");
    segments.push(splitDay[0]);
    var splitTime = splitDay[1].split(":");
    segments.push(splitTime[0]);
    segments.push(splitTime[1]);
    segments.push(splitTime[2]);

    return segments;
}

function formatTimeFromJSON(time) {
    var timeString = String(time);
    var split = timeString.split(" ");
    var year = split[3];
    var month = monthToInt(split[1]);
    var day = split[2];
    var timeSplit = split[4].split(":");
    var hour = timeSplit[0];
    return year + "-" + month + "-" + day + " " + hour + ":00:00";
}

function verifyTimeConflicts(roomID, startTime, endTime) {
    var status = false;
    startTimeSplit = splitTime(startTime);
    endTimeSplit = splitTime(endTime);

    start_year = startTimeSplit[0];
    start_month = startTimeSplit[1];
    start_day = startTimeSplit[2];
    start_hour = startTimeSplit[3];

    end_year = endTimeSplit[0];
    end_month = endTimeSplit[1];
    end_day = endTimeSplit[2];
    end_hour = endTimeSplit[3];

    getReservations(roomID).success(function(data){
        roomReservations = deserializeReservation(data);
        roomReservations.forEach(function(reservation) {
        var start = String(reservation.start);
        var end = String(reservation.end);
        var startSplit = start.split(" ");
        var endSplit = end.split(" ");
        var startYear = startSplit[3];
        var endYear = endSplit[3];

        if(start_year != startYear && end_year != endYear) {
            return;
        }

        var startMonth = monthToInt(startSplit[1]);
        var endMonth = monthToInt(startSplit[1]);

        if(start_month != startMonth && end_month != endMonth) {
            return;
        }

        var startDay = startSplit[2];
        var endDay = endSplit[2];

        if(start_day != startDay && end_day != endDay) {
            return;
        }

        var startTimeSplit = startSplit[4].split(":");
        var endTimeSplit = endSplit[4].split(":");
        var startHour = startTimeSplit[0];
        var endHour = endTimeSplit[0];

        if(end_hour <= startHour || start_hour >= endHour) {
            return;
        }
        else {
            status = true;
        }
    });
    });
    return status;
}

function monthToInt(month) {
    switch(month) {
            case "Jan":
                return 1;
            case "Feb":
                return 2;
            case "Mar":
                return 3;
            case "Apr":
                return 4;
            case "May":
                return 5;
            case "Jun":
                return 6;
            case "Jul":
                return 7;
            case "Aug":
                return 8;
            case "Sep":
                return 9;
            case "Oct":
                return 10;
            case "Nov":
                return 11;
            case "Dec":
                return 12;
            default:
                return 0;
        }
}

function modifyReservation() {
    var userID = getCookie("studentId");
    var reservationID = $("#modifyDataID").data("reservationID");
	var updatedRoom = $("#modifyRoom").val();
    var updatedYear = $("#modifyYear").val();
    var updatedMonth = $("#modifyMonth").val();
    var updatedDay = $("#modifyDay").val();

	var updatedStart = $("#modifyStart_time").val();
	var updatedEnd = $("#modifyEnd_time").val();

    if(updatedRoom == null || updatedYear == null || updatedMonth == null || updatedDay == null || updatedStart == null || updatedEnd == null) {
        console.error("Missing reservation field for edit");
        return;
    }

    var day_time = parseTimeValue(updatedDay);
    var start_time = parseTimeValue(updatedStart);
    var end_time = parseTimeValue(updatedEnd);

    var startDate = String(updatedYear) + "-" + String(updatedMonth) + "-" + day_time + " " + start_time + ":00:00";
    var endDate = String(updatedYear) + "-" + String(updatedMonth) + "-" + day_time + " " + end_time + ":00:00";

    if(!verifyTimeConflicts(updatedRoom, startDate, endDate)) {
        $.ajax({
            type: 'POST',
            contentType: "application/x-www-form-urlencoded",
            async: false,
            url: '/updateReservation',
            data: {userID: userID, reservationID: reservationID, dataRoom: updatedRoom, startTime: startDate, endTime: endDate},
            success: function(data){
                console.log(data);
                window.location.reload(true);
            }
        });
    }
    else {
        // TODO waiting list
    }
}

function parseTimeValue(value){
    var parsedTimeValue;
    if(parseInt(value)<10) {
        parsedTimeValue = "0" + String(value);
    }
    else {
        parsedTimeValue = String(value);
    }
    return parsedTimeValue;
}

function deleteReservation(reservationID, roomNumber) {
    var reservationID = reservationID;
    var roomNumber = roomNumber;
    $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        url: '/deleteReservation',
        data: { reservationID: reservationID },
        success: function(resp){
            updateWaitingList(roomNumber);
            window.location.reload(true);
        }
    });

}

function updateWaitingList(room) { // This function updates the waitlist by checking if someone's reservation can be created.
    var idsToRemove = [];
    getAllWaitingListEntriesByRoom(room).success(function(data){ //TODO Darrel the ajax call returns a list of JsonWaitingReservation
        entries = deserializeReservation(data);  //TODO Darrel entries should be all the waiting list elements which have the same room. Not sure if entries is populated proberly. 
        entries.forEach(function(entry) {
            var startTime = formatTimeFromJSON(entry.start);
            var endTime = formatTimeFromJSON(entry.end);
            if(!verifyTimeConflicts(entry.room, startTime, endTime)) {
                pushReservation(entry.userId, entry.room, startTime, endTime);
                idsToRemove.push(entry.id);
            }
            else {
                return;
            }
        });
        removeWaitListEntries(idsToRemove);
    });
}

function removeWaitListEntries(idsToRemove) {
    idsToRemove.forEach(function(id) {
        $.ajax({
            type: 'POST',
            contentType: "application/x-www-form-urlencoded",
            async: false,
            url: '/removeWaitListEntriesById',
            data: {waitListId: id},
        });
    });
}

function getAllWaitingListEntriesByRoom(room) {
    return $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        async: false,
        url: '/getWaitListEntriesByRoom',
        data: {dataRoom: room},
        success: function(data){
        console.log(data);
        }
    });
}

function deserializeReservation(reservations){
    if(reservations != undefined && reservations.length > 0){
        var result = [];
        reservations.forEach(function(reservationJSON){
            var reservation = new Reservation(reservationJSON.reservationID, reservationJSON.subject, reservationJSON.roomNumber, reservationJSON.startTime, reservationJSON.endTime, reservationJSON.studentID);
            result.push(reservation)
        })
    } else{
        console.error("No reservations found."); //TODO darrel we end up here because the reservation is not found.
        return [];
    }
    return result;
}

function changeRoom(roomNumber, el){
    $(".reservations-table").empty();
    buildCalendar(roomNumber, el);
}

function populateTime() {
    var select = $("#start_time");
    var hours
    for (var i = 420; i <= 1320; i += 60){
        hours = Math.floor(i/60);
        select.append($('<option></option>').attr('value',hours).text(hours + ':00'));
    }
}
function populateEndTime() {
    var select = $("#end_time");
    var startTime = $("#start_time").val();
    select.empty();
    for (var i = parseInt(startTime) + 1; i <= 22 ; i++) {
        
        select.append($('<option></option>').attr('value',i).text(i + ':00'));
    }
    $('select').material_select();
}

function daysInMonth(m, y) {
    return /8|3|5|10/.test(--m)?30:m==1?(!(y%4)&&y%100)||!(y%400)?29:28:31; //1337 hax
}

function populateDays(month, year) {
    var select = $("#day");
    $("#day").empty();
    var days = daysInMonth(parseInt(month), parseInt(year));
    for (var i = 1; i <= days ; i += 1) {
        
        select.append($('<option></option>').attr('value', i).text(i));
    }
    $('select').material_select();
}

function updateDays() {
    var selectedMonth = $("#month").val();
    var selectedYear = $("#year").val();
    populateDays(selectedMonth, selectedYear);
}
