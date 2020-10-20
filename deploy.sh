echo "Creating deployment executable..."

rm -rf ./dist || true

mkdir dist
cp config.yaml dist/

echo "Building for linux ..."

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/server.linux .

echo "Done :)"