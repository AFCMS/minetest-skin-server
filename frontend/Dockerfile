# syntax=docker/dockerfile:1

FROM node:20-alpine3.19 as frontend-development

# Update NPM
# https://github.com/npm/cli/issues/6684
RUN npm update -g npm

RUN mkdir /frontend
WORKDIR /frontend

CMD npm install --include=dev && npm run dev
