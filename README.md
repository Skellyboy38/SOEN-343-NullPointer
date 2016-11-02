***************************************************** GO ********************************************************

1. Download and install the go binaries. Version 1.7
2. Set $GOROOT to the path where your binaries were installed. (Usualy C:/Go)
3. Create a workspace for your Go projects. Ex (C:Documents/CodeProjects/Go)
4. Set $GOPATH as the path of the directory you just created (Workspace)
5. Set $GOBIN as $GOROOT/bin
6. The Go directory you create should have 3 subdirectories labelled src pkg bin.
7. Once all this is done make sure to open a new terminal.
8. Run go get github.com/Skellyboy38/SOEN-343-NullPointer
	This will pull our repository.
9. cd into C:Documents/CodeProjects/Go/src/github.com/Skellyboy38/SOEN-343-NullPointer/Layers
10. To compile go programs you must run go build <File with main func>
	This will create an executable.
	In our case run go build controller.go
	The resulting executable will be controller.exe
11. You can now run the executable which is our application.
12. If compilation argues that you have missing packages run a go get on those missing packages. 
	Compilation should work once all the depencies have been resolved.
13. Alternately go run controller.go will run the application without creating an executable.
14. go install controller.go will create an executable in the GOBIN directory we specified earlier.

*****************************************************************************************************************


************************************************ DATABASE SETUP *************************************************

1. Download and Install Postgres version 9.7
2. Change directory until data_source_layer/setup is your working directory.
3. Run bash dbSetup.sh
	This will create a soen343 user which we will use in our application.
	It will initialise the database and create its directory called registry
	in the current working directory. (data_source_layer/setup).
	It will run the database server.
	It will create the tables we are using and populate the users table with hardcoded users.
	The database seems to close everytime the application is closed or stops running. 
	Therefore before everytime you run the application you must first start the 
	database server with pg_ctl.exe -D registry start. Once again in the same directory as before.

******************************************************************************************************************


