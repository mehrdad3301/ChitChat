echo "creating table"
cat schema.sql | sqlite3 database.db
echo "building the go binary"
go build -o ChitChat 

echo "starting the binary"
./ChitChat
