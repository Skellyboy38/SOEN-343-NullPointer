function Reservation(id, subject, calendar, start, end, userId) {
	this.id = id;
    startDate = parseDate(start);
    endDate = parseDate(end);
    this.room = calendar;
    this.calendar = calendar; // Room Number
    this.userId = userId;
    this.start = new Date(startDate[0], startDate[1], startDate[2], startDate[3], startDate[4], startDate[5]); // Start Time
    this.end = new Date(endDate[0], endDate[1], endDate[2], endDate[3], endDate[4], endDate[5]); // End Time
    this.readOnly = true
    this.subject = "Subject: " + subject + 
    "\n" + "Start: " + this.start.getHours() + ":" + this.start.getMinutes() + ":" + this.start.getSeconds() +
    "\n" + "End: " + this.end.getHours() + ":" + this.end.getMinutes() + ":" + this.end.getSeconds();
	if(getCookie("studentId") == userId){
		this.calendar = "Current Reservations";
	} else {
		this.calendar = "Others";
	}

}

function parseDate(date) {
	dateToReturn = [];
	dateParse = date.split("-");
	year = dateParse[0];
	month = dateParse[1];
	timeT = dateParse[2].split("T");
	day = timeT[0];
	timeZ = timeT[1].split("Z");
	timeParse = timeZ[0].split(":");
	hour = timeParse[0];
	minute = timeParse[1];
	second = timeParse[2];
	dateToReturn.push(parseInt(year));
	dateToReturn.push(parseInt(month) - 1);
	dateToReturn.push(parseInt(day));
	dateToReturn.push(parseInt(hour));
	dateToReturn.push(parseInt(minute));
	dateToReturn.push(parseInt(second));
	return dateToReturn;
}