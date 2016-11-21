function clearModal() {
    $("#modifyStart_time").empty();
    $("#modifyEnd_time").empty();
}

function initializeModifyModal(reservation) {
    console.log(reservation);
    var roomNumber = reservation.roomNumber;
    var startTimeSegments = parseDate(reservation.startTime);
    var endTimeSegments = parseDate(reservation.endTime);
    var year = startTimeSegments[0];
    var month = padNumber(startTimeSegments[1] + 1);
    var day = padNumber(startTimeSegments[2]);
    var startTime = startTimeSegments[3];
    var endTime = padNumber(endTimeSegments[3]);
    $("#modifyDataID").data("reservationID", reservation.reservationID);
    populateModifyStartTime();
    populateModifyEndTime();
    $("#modifyRoom").val(roomNumber);
    $("#modifyYear").val(year);
    $("#modifyMonth").val(month);
    populateModifyDays(month, year);
    $("#modifyDay").val(day);
    $("#modifyStart_time").val(startTime);
    $("#modifyEnd_time").val(endTime);
    $('select').material_select();

    function padNumber(d) {
        return (d < 10) ? '0' + d.toString() : d.toString();
    }
}

function populateModifyDays(month, year) {
    var select = $("#modifyDay");
    $("#modifyDay").empty();
    var days = daysInMonth(parseInt(month), parseInt(year));
    for (var i = 1; i <= days; i += 1) {
        select.append($('<option></option>').attr('value', i).text(i));
    }
    $('select').material_select();
}

function populateModifyStartTime() {
    var select = $("#modifyStart_time");
    var hours;
    for (var i = 420; i <= 1320; i += 60) {
        hours = Math.floor(i / 60);
        select.append($('<option></option>').attr('value', hours).text(hours + ':00'));
    }
    $('select').material_select();
}

function populateModifyEndTime() {
    var select = $("#modifyEnd_time");
    var startTime = $("#modifyStart_time").val();
    select.empty();
    for (var i = parseInt(startTime) + 1; i <= 22; i++) {

        select.append($('<option></option>').attr('value', i).text(i + ':00'));
    }
    $('select').material_select();
}