$(document).ready(function () {
    populateTime();
    populateDays(1, 2016);
    printTodayDate();
    buildCalendar(1); // Default room is 1
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

function renderUserReservationList(reservations){
    var reservationListHTML = $(".reservations-table");
    var reservationHeaderHTML = renderReservationHeader();
    reservationHeaderHTML.appendTo(reservationListHTML);
    if(reservations.length ===0){
        var row = $("<div></div>",{
            class: "row",
            text: "No reservations available."
        })
        row.appendTo(reservationListHTML);
    } else {
        reservations.forEach(function(resv){
            var row = renderReservationRow(resv);
            row.appendTo(reservationListHTML);
        });
    }

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
            class: "cell"
        });
        var deleteBtn = $("<a></a>", {
            class: "Waves-effect waves-light btn deleteBtn",
            text: "Delete",
            "data-reservationid": resv.reservationID
        });
        deleteBtn.data("reservationID", resv.reservationID);
        deleteBtn.click(function(){
            deleteReservation($(this).attr("data-reservationid"));
        });
        deleteBtn.appendTo(actionsCell)
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
    var year = $("#year").val();
    var month = $("#month").val();
    var day = $("#day").val();

	var start = $("#start_time").val();
	var end = $("#end_time").val();
    var start_time = "";
    var end_time = "";

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

    var startDate = String(year) + "-" + String(month) + "-" + String(day) + " " + start_time + ":00:00";
    var endDate = String(year) + "-" + String(month) + "-" + String(day) + " " + end_time + ":00:00";

    console.log(startDate);
    console.log(endDate);

    if(!verifyTimeConflicts(room, startDate, endDate)) {
        $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        async: false,
        url: '/createReservation',
        data: {userID: userID, dataRoom: room, startTime: startDate, endTime: endDate},
    });
        location.reload();
    }
    else {
        console.log("A time conflict exists. Abort.");
    }
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
        roomReservations = getReservationsSuccess(data);
    });
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

        if(start_month != startMonth && end_month != endMoth) {
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

// TODO
function modifyReservation() {
    $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        url: '/modifyReservation',
        data: {},
    });
}

function deleteReservation(reservationID) {
    var reservationID = reservationID;
    $.ajax({
        type: 'POST',
        contentType: "application/x-www-form-urlencoded",
        url: '/deleteReservation',
        data: { reservationID: reservationID },
        success: function(){
            buildCalendar(1);
        }
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

function printTodayDate(){
    var today = new Date();
    var options = {
    weekday: "long", year: "numeric", month: "short",
    day: "numeric", hour: "2-digit", minute: "2-digit"
    };
    $("#todayDate").html(today.toLocaleTimeString("en-us", options));
}

function populateTime() {
    var select = $("#start_time");
    //var selectHTML = $("<select></elect")
    var hours
    for (var i = 420; i <= 1320; i += 60){
        hours = Math.floor(i/60);
        /*var optionsHTML = $("<option></option>",{
            value : i
        });
        optionsHTML.html(hours + ':00');
        optionsHTML.appendTo(select);
        */select.append($('<option></option>').attr('value',hours).text(hours + ':00'));
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

//this.$("#start_time").value

function daysInMonth(m, y) {
    return /8|3|5|10/.test(--m)?30:m==1?(!(y%4)&&y%100)||!(y%400)?29:28:31; //1337 hax
}

function populateDays(month, year) {
    var select = $("#day");
    // for (var i = 0 ; i < select.options.length; i++){
    //     select.options[i] = null;
    // }
    $("#day").empty();
    var days = daysInMonth(parseInt(month), parseInt(year));
    for (var i = 1; i <= days ; i += 1) {
        
        select.append($('<option></option>').attr('value', i).text(i));
    }
    $('select').material_select();
}
function updateDays() {
    //var selectedMonth = document.getElementById("month").value;
    //var selectedYear = document.getElementById("year").value;
    var selectedMonth = $("#month").val();
    var selectedYear = $("#year").val();
    populateDays(selectedMonth, selectedYear);
}
