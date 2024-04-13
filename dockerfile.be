FROM golang:1.22.2

WORKDIR /tkbai-dashboard

RUN mkdir -p /tkbai-dashboard/be \
    && mkdir -p /tkbai-dashboard/migration \
    && mkdir -p /tkbai-dashboard/be/log 

COPY migration /tkbai-dashboard/migration
COPY be/ /tkbai-dashboard/be/

RUN cd /tkbai-dashboard/be/ && go mod tidy
RUN cd /tkbai-dashboard/be/ && go build tkbai-be.go

ENV BE_SERVER_PORT=":9070"
ENV BE_WEB_HOST="http://localhost:9071"
ENV BE_API_PREFIX="/api"
ENV BE_BASE_URL="/tkbai"
ENV BE_WEB_STATIC_URL="/static"
ENV BE_WEB_TEMPLATES_PATH="/sertif-validator/src/public/view/*.html"
ENV BE_WEB_STATIC_PATH="/sertif-validator/src/public"
ENV BE_SV_JWT_KEY="LmPZJbddZ9uXW4JE7g6N9R8ZdmDRv5vYihZJRBcOz7U="
ENV BE_KC_DOCKER_URL="http://keycloak:8080"
ENV BE_KC_URL="http://localhost:8080"
ENV BE_KC_SECRET="rK4sSxVGdIjnEKnyBFzdymlN62stQ72m"
ENV BE_KC_INSECURE="TRUE"
ENV BE_KC_ID="tkbai"
ENV BE_KC_REALM="tkbai_dev"
ENV BE_KC_STATE="authExt"
ENV BE_KC_LOGIN_REDIRECT="http://localhost:9070/sv/api/auth/loginCallback"
ENV BE_KC_LOGIN_302REDIRECT='http://localhost:9070/sv/admin/dash'
ENV BE_KC_LOGOUT_REDIRECT="http://localhost:9070/sv/api/auth/logoutCallback"
ENV BE_KC_LOGIN_REDIRECT_PATH="/api/auth/loginCallback"
ENV BE_KC_LOGIN_302REDIRECT_PATH='/admin/dash'
ENV BE_KC_LOGOUT_REDIRECT_PATH="/api/auth/logoutCallback"
ENV BE_TKBAI_DB_URL="root:03IZmt7eRMukIHdoZahl@tcp(mysql:3306)/tkbai"


ENTRYPOINT [ "/tkbai-dashboard/be/tkbai-be" ,"/tkbai-dashboard/fe/tkbai-fe"]