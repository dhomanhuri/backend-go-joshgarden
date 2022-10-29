sudo docker build -t joshgarden-app:latest .
sudo docker run -d -p 8080:8080 --name joshgarden joshgarden-app:latest
