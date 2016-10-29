CREATE TABLE userTable (
	studentId INTEGER UNIQUE PRIMARY KEY,
	username TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
);

CREATE TABLE session (
	sessionId SERIAL UNIQUE PRIMARY KEY,
	studentId INTEGER references userTable
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
	studentId INTEGER references userTable,
	timeSlotId SERIAL references timeSlot
);

CREATE TABLE waitlist (
	waitlistId SERIAL UNIQUE PRIMARY KEY,
	capasity int
);

CREATE TABLE waitlistcontract (
	studentId INTEGER references userTable,
	waitlistId SERIAL references waitlist,
	PRIMARY KEY (studentId,waitlistId)
);
