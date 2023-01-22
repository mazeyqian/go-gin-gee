# Example: bash ./scripts/docker-build.sh "RUN" "WECOM_ROBOT_CHECK=b2d57746-7146-44f2-8207-86cb0ca832be"

echo "Start Build Docker"

# ENV
RUN_FLAG=$1
WECOM_ROBOT_CHECK_ENV_STR=$2
echo ${WECOM_ROBOT_CHECK_ENV_STR}

# ProjectName/SubName
preVersion="go-gin-gee/api"
# Port
visitPort="3000"
innerPort="3000"

# Build
# GOOS=linux GOARCH=amd64 go build -o dist/api-linux-amd64 cmd/api/main.go

# Stop
echo "Stop Docker Containers"
docker ps
docker ps|awk '{if (NR!=1) print $1}'| xargs docker stop

# Remove
echo "Remove Docker Image"
docker ps -a -q
docker rm $(docker ps -a -q)

# Generate
randomVersion=${RANDOM}
combinedVersion="${preVersion}:v${randomVersion}"
echo "Generate random version: ${combinedVersion}"

# Build
echo "Build Docker Image: ${combinedVersion}"
docker build -t ${combinedVersion} . -f ./Dockerfile

# Run
# https://stackoverflow.com/questions/20449543/shell-equality-operators-eq
if [ ${RUN_FLAG} = "RUN" ]; then
  echo "Run Docker"
  docker run -e ${WECOM_ROBOT_CHECK_ENV_STR} -d -p ${visitPort}:${innerPort} ${combinedVersion}
  # Notification
  echo "Complete, Visit: http://localhost:${visitPort}/api/ping"
else
  echo "RUN_FLAG: ${RUN_FLAG}"
  echo "All done."
fi
