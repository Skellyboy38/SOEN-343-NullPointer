function Reservation(subject, calendar, start, end) {
    subject === undefined ? this.subject = "Reservation" : this.subject = subject; // Description
    this.calendar = calendar; // Room Number
    this.start = new Date(start); // Start Time
    this.end = new Date(end); // End Time
    this.readOnly = true
}