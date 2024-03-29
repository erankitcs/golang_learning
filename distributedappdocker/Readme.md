### Distributed App built using Golang.

#### Create a user defined network
```
docker network create teacherportal
```

#### Registry service
```
docker build -f Dockerfile.registryservice -t registryservice .
docker run -p 3000:3000 --name registryservice --network=teacherportal -d registryservice
```

#### Log service
```
docker build -f Dockerfile.logservice -t logservice .
docker run -p 4000:4000 --name logservice --network=teacherportal -d logservice
```

#### Grading service
```
docker build -f Dockerfile.gradingservice -t gradingservice .
docker run -p 6000:6000 --name gradingservice --network=teacherportal -d gradingservice
```

#### Teacher Portal service
```
docker build -f Dockerfile.teacherportalservice -t teacherportalservice .
docker run -p 5000:5000 --name teacherportalservice --network=teacherportal -d teacherportalservice
```

#### Application URL
http://localhost:5000


#### Useful links
1. https://marcofranssen.nl/docker-tips-and-tricks-for-your-go-projects

