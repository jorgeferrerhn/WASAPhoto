# WASAPhoto

Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! You can
upload your photos directly from your PC, and they will be visible to everyone following you.

WASAPhoto is a simulator of a social web app made specially for the Web And Software Architecture subject. It contains an API for storing and receiving the information, a server written in Go that works on the backend, a frontend written in HTML, CSS and Javascript and it implements a Docker container to deploy all the application. 

# How to build container images
# Backend
`docker build -t wasa-photos-backend:latest -f Dockerfile.backend .`
# Frontend`
`docker build -t wasa-photos-frontend:latest -f Dockerfile.frontend .`

# How to run container images`
# Backend
`docker run -it --rm -p 3000:3000 wasa-photos-backend:latest`
#Frontend
`docker run -it --rm -p 8081:80 wasa-photos-frontend:latest`
