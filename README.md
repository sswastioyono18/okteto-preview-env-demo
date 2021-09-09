# okteto-test-project
This is a simple project to demonstrate cloud dev in okteto using Golang. 

## How to (Development Flow)
Assuming you have github token and okteto token setup in your repo

1. Build Dockerfile
Once you are done with your change, you have to build docker image and push it
```
docker buildx build  . -t sactio1811/oktetotest:latest --push
```

2. then commit and open PR in github
3. Github action will trigger preview environment creation for feature branch
4. if you have another change, you have to re-build docker image again then git push. Feature branch will be updated
5. When you merge to master, preview environment for feature branch will be closed. Master feature branch will be updated




