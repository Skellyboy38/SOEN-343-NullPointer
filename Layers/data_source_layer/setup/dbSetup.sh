psql -U postgres -c 'CREATE DATABASE registry'
psql -U postgres -d registry -c "CREATE USER soen343 PASSWORD 'soen343'"
psql -U postgres -d registry -c 'GRANT ALL ON DATABASE registry TO soen343'