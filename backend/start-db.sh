docker build -t fixtheplanet/mysql -f backend/local-db/Dockerfile .
docker run -e MYSQL_ROOT_PASSWORD=local-password -d -p 3306:3306/tcp fixtheplanet/mysql