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

function populateDays(month, year) {
    var select = $("#day");
    $("#day").empty();
    var days = daysInMonth(parseInt(month), parseInt(year));
    for (var i = 1; i <= days ; i += 1) {
        
        select.append($('<option></option>').attr('value', i).text(i));
    }
    $('select').material_select();
}
