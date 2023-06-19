# Example: bash ./scripts/docker-build.sh "RUN" "WECOM_ROBOT_CHECK=b2lsjd46-7146-4nv2-8767-86cb0cncjdbe" "BASE_URL=https://example.com/path/"

echo "Start Build Docker"

# ENV
RUN_FLAG=$1
WECOM_ROBOT_CHECK_ENV_STR=$2
echo ${WECOM_ROBOT_CHECK_ENV_STR}
BASE_URL_ENV_STR=$3
echo ${BASE_URL_ENV_STR}

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
  # docker run -e ${WECOM_ROBOT_CHECK_ENV_STR} -d -p ${visitPort}:${innerPort} ${combinedVersion}
  docker run -e ${WECOM_ROBOT_CHECK_ENV_STR} -e ${BASE_URL_ENV_STR} -d -p ${visitPort}:${innerPort} ${combinedVersion}
  # Notification
  echo "Complete, Visit: http://localhost:${visitPort}/api/ping"
else
  # Build
  # Example: 20230122153113
  # https://www.cyberciti.biz/faq/linux-unix-formatting-dates-for-display/
  DATE_FORMAT=$(date +"%Y%m%d%H%M%S")
  REPOSITORY_TAGNAME="mazeyqian/go-gin-gee:v${DATE_FORMAT}-api"
  echo "DATE_FORMAT: ${DATE_FORMAT}"
  docker tag ${combinedVersion} ${REPOSITORY_TAGNAME}
  docker push ${REPOSITORY_TAGNAME}
  echo "RUN_FLAG: ${RUN_FLAG}"
  echo "REPOSITORY_TAGNAME: ${REPOSITORY_TAGNAME}"
  echo "All done."
fi
