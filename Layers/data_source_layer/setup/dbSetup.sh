initdb registry
pg_ctl.exe start -D 'registry' -l logfile
sleep 10
createdb registry
createuser -s -l -w -d soen343
psql -U soen343 -d registry -f initUser.sql
psql -U soen343 -d registry -f initDbTables.sql
psql -U soen343 -d registry -f fillUserTable.sql