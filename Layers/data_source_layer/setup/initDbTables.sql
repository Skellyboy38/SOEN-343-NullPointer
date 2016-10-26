
CREATE TABLE userTable (
	studentId SERIAL UNIQUE PRIMARY KEY
);

CREATE TABLE room (
	roomId	SERIAL UNIQUE PRIMARY KEY,
	roomNumber TEXT UNIQUE
);

CREATE TABLE timeSlot (
		timeSlotId SERIAL UNIQUE PRIMARY KEY,
		startTime TIMESTAMP,
		endTime TIMESTAMP
);

CREATE TABLE reservation (
	reservationId SERIAL UNIQUE PRIMARY KEY,
	roomId SERIAL references room,
	studentId SERIAL references userTable,
	timeSlotId SERIAL references timeSlot
);

CREATE TABLE waitlist (
	waitlistId SERIAL UNIQUE PRIMARY KEY,
	capasity int
);

CREATE TABLE waitlistcontract (
	studentId SERIAL references userTable,
	waitlistId SERIAL references waitlist,
	PRIMARY KEY (studentId,waitlistId)
);
