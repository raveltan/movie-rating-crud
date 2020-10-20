echo "Creating Moorater deployment executable..."
echo ""
echo "Moving files to dist folder..."
rm -rf ./dist || true

mkdir dist
cp config.yaml dist/
cp schema/migration.sql dist/

echo "Building for linux ..."

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/mooreview .

echo "Creating tar package ..."
tar -czf dist.tar.gz dist/
mv dist.tar.gz dist/

echo "Done :)"