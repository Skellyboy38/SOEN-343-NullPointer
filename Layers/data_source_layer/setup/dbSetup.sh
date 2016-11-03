initdb registry
pg_ctl.exe start -D registry -l logfile
echo database server started
sleep 10
createdb registry
echo database registry created
createuser -s -l -w -d soen343
echo user soen343 created
psql -U soen343 -d registry -f initUser.sql
echo  inituser
psql -U soen343 -d registry -f initDbTables.sql
echo inittables
psql -U soen343 -d registry -f fillUserTable.sql