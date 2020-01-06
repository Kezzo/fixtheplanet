cd backend/src/
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../bin/main .
cd ../
docker build -t fixtheplanet .