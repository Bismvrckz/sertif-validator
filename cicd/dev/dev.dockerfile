FROM debian:bookworm-slim

RUN adduser --system --group --no-create-home tkbai

RUN mkdir -p /tkbai-dashboard/fe \
    && mkdir -p /tkbai-dashboard/be \
    && apt update \
    && mkdir -p /tkbai-dashboard/fe/public \
    && mkdir -p /tkbai-dashboard/migration \
    && mkdir -p /tkbai-dashboard/fe/log \
    && mkdir -p /tkbai-dashboard/be/log 

COPY be/tkbai-be /tkbai-dashboard/be/tkbai-backend
COPY fe/tkbai-fe /tkbai-dashboard/fe/tkbai-frontend
COPY fe/public /tkbai-dashboard/fe/public
COPY migration /tkbai-dashboard/migration
COPY fe/.env /tkbai-dashboard/fe/.env
COPY be/.env /tkbai-dashboard/be/.env

RUN chown -R tkbai:tkbai /tkbai-dashboard && chmod 755 /tkbai-dashboard/fe/tkbai-frontend && chmod 755 /tkbai-dashboard/be/tkbai-backend

USER tkbai

WORKDIR /tkbai-dashboard