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

CREATE TABLE waitListMaster (
	waitlistID SERIAL UNIQUE PRIMARY KEY,
	roomId INTEGER,
	studentId INTEGER references userTable,
	startTime TIMESTAMP,
	endTime TIMESTAMP
);
