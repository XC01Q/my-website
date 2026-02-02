# Production build and run
docker compose up --build
# Development with hot reload
docker compose --profile dev up web-dev
# Image build only
docker build -t my-website .
