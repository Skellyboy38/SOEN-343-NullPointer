CREATE TABLE userTable (
	studentId INTEGER UNIQUE PRIMARY KEY,
	password TEXT NOT NULL
);

CREATE TABLE reservation (
	reservationId SERIAL UNIQUE PRIMARY KEY,
	roomId INTEGER,
	studentId INTEGER references userTable,
	startTime TIMESTAMP,
	endTime TIMESTAMP
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
