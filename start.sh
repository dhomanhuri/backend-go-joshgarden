sudo docker build -t joshgarden-app:latest .
sudo docker run -d -p 80:8080 --name joshgarden joshgarden-app:latest
